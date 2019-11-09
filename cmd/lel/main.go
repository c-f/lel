package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/c-f/lel"
	"github.com/c-f/lel/config"
	"github.com/c-f/lel/utils"
	_ "time"

	yml "gopkg.in/yaml.v2"
	// homedir "github.com/mitchellh/go-homedir"
)

var (
	// types
	modes = []string{
		"single",
		"team",

		"generate",
		"install",
	}
)

// User flags
var (
	lelDir = flag.String("docu", "", "Directory of the documentation")
	appDir = flag.String("appDir", utils.GetLeLExePath(), "Directory of LEL where the certificates and static files are stored")

	addr         = flag.String("addr", "127.0.0.1:8888", "Listening URL")
	editor       = flag.String("editor", "", "Choose favorite code editor path")
	enableUpload = flag.Bool("enable-upload", true, "Enable Upload ")

	configFile  = flag.String("c", "", "Config stuff")
	enableDebug = flag.Bool("debug", false, "Enable debug modes for development")
	mode        = flag.String("mode", "single", strings.Join(modes, ", "))
)

// MAIN
func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, config.BANNER)
		flag.PrintDefaults()
	}
	flag.Parse()

	var c config.Config

	// projectConfig
	projectConf := config.ProjectConfig{
		NotesDir: filepath.Join(*lelDir, "notes"),

		ImageDir:  filepath.Join(*lelDir, "images"),
		VideoDir:  filepath.Join(*lelDir, "videos"),
		MisatoDir: filepath.Join(*lelDir, "misato"),
		GraphDir:  filepath.Join(*lelDir, "graphs"),
	}

	// get config if provided
	if *configFile != "" {
		f, err := os.Open(*configFile)
		if err != nil {

		}
		defer f.Close()

		decoder := yml.NewDecoder(f)
		err = decoder.Decode(&c)
		if err != nil {
			fmt.Printf("[!] %s\n", err)
			os.Exit(1)
		}

		// Change current settings if different from default
		if *enableUpload != true {
			c.Server.EnableUpload = *enableUpload
		}
		if *enableDebug != false {
			c.Server.Debug = *enableDebug
		}
		if *mode != "single" {
			c.Server.Mode = *mode
		}
		if *addr != "127.0.0.1:8888" {
			c.Server.ListenURL = *addr
		}
		if *editor != "/usr/bin/subl" {
			c.Server.Editor = *editor
		}
		if *lelDir != "" {
			c.Project = projectConf
		}

	} else {
		if *lelDir == "" {
			fmt.Printf("[!] %s\n", "Please specify the the -docu option")
			os.Exit(1)
		}

		// set the config
		c = config.Config{
			Server: config.ServerConfig{
				ListenURL: *addr,

				Editor:    *editor,
				StaticDir: filepath.Join(*appDir, "static/dist"),

				Debug:        *enableDebug,
				EnableUpload: *enableUpload,

				TLS:      true,
				CertPath: filepath.Join(*appDir, "crts"),

				Mode: *mode,
				Auth: map[string]string{},
			},
			Project: projectConf,
		}

		// generate the config and prints it to stdout
		if *mode == "generate" {
			c.Server.Mode = "single"
			bts, err := yml.Marshal(&c)
			if err != nil {
				fmt.Println("[!] %s\n", err)
			}
			fmt.Println(string(bts))
			os.Exit(0)
		}
		if os.Getenv("LEL_FRONTEND_URL") != "" {
			
			c.Server.Develop = os.Getenv("LEL_FRONTEND_URL")
		}

	}

	// --[new Server ]--
	server := lel.New(&c)

	err := server.Start()
	if err != nil {
		fmt.Printf("[!] %s\n", err)
	}
}
