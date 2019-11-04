package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/c-f/lel/config"
	"github.com/c-f/lel/utils"
	"github.com/gorilla/mux"

	"github.com/kabukky/httpscerts"
)

// Server is the webserver in order to display all entries
type Server struct {
	server *http.Server
	c      *config.ServerConfig
}

// NewServer creates a new Server based on the configs
func NewServer(sConf *config.ServerConfig, pConf *config.ProjectConfig) (s *Server, err error) {
	if err = sConf.Valid(); err != nil {
		return
	}
	if err = sConf.Prepare(); err != nil {
		return
	}
	if err = pConf.Valid(); err != nil {
		return
	}
	if err = pConf.Prepare(); err != nil {
		return
	}

	s = &Server{
		server: &http.Server{
			ReadTimeout: 10 * time.Second,
			Addr:        sConf.ListenURL,
		},
		c: sConf,
	}
	Routes.Set("Hostname", s.c.Hostname)

	// --[handler]--
	api, err := NewAPIHandler(pConf.NotesDir, s.c.Editor, s.c.Auth, s.c.Mode)
	if err != nil {
		return
	}
	// register subroutes
	router := mux.NewRouter()
	router.HandleFunc("/", s.Base)

	// Debug Output
	if s.c.Debug {
		log.Printf("[*] Development mode on \n")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		router.Use(devMiddleware)
	}
	if s.c.StaticDir != "" {
		static := NewStaticHandler(s.c.StaticDir)
		router.PathPrefix("/static/").Handler(static)
	}

	// --[api handler]--
	if pConf.GraphDir != "" {
		log.Printf("[+] Created Graph API \n")
		api.AddGraphHandler(pConf.GraphDir)
	}
	if pConf.MisatoDir != "" {
		log.Printf("[+] Created Misato API \n")
		api.AddMisatoHandler(pConf.MisatoDir)
	}
	if pConf.ImageDir != "" && pConf.VideoDir != "" {

		staticImage := NewSuffixHandler(pConf.ImageDir, ".png")
		staticVideo := NewSuffixHandler(pConf.VideoDir, ".webm")

		router.PathPrefix("/images/").Handler(staticImage("/images/"))
		router.PathPrefix("/videos/").Handler(staticVideo("/videos/"))

		// Enable Upload
		if s.c.EnableUpload {
			log.Printf("[+] Created Upload API \n")
			api.AddUploadHandler(pConf.ImageDir, pConf.VideoDir)
		}
	}

	// initialize api routes
	api.Init()
	router.PathPrefix("/api/").Handler(api)

	// set Handler
	s.server.Handler = router

	return s, nil
}

// Base returns the app template for vue js
func (s *Server) Base(w http.ResponseWriter, r *http.Request) {
	getBase(w).ExecuteTemplate(w, "base", Routes.Config)
}

// Start starts the webserver
func (s *Server) Start() error {
	if s.c.TLS {

		keyFile, err := utils.GetAbsPath(s.c.CertPath, CERTPREFIX+".key.pem")
		if err != nil {
			return err
		}
		crtFile, err := utils.GetAbsPath(s.c.CertPath, CERTPREFIX+".pem")
		if err != nil {
			return err
		}

		if err := httpscerts.Check(crtFile, keyFile); err != nil {
			log.Printf("[*] Generate CA and HTTPS certs\n")
			serverlist := fmt.Sprintf("127.0.0.1:8888,srv.l3l.lol:8888,srv.l3l.lol:4443,%s", s.c.Hostname)

			err = httpscerts.Generate(crtFile, keyFile, serverlist)
			if err != nil {
				log.Fatal("Error: Couldn't create https certs.")
			}
		}

		log.Printf("[*] Using HTTPS")
		log.Printf("Starting admin server at https://%s", s.c.ListenURL)
		return s.server.ListenAndServeTLS(crtFile, keyFile)
	} else {

		log.Printf("[!] Insecure HTTP - switch to HTTPS !")
		log.Printf("Starting admin server at http://%s", s.c.ListenURL)
		return s.server.ListenAndServe()
	}
}

// Shutdown attempts to gracefully shutdown the server.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	return s.server.Shutdown(ctx)
}
