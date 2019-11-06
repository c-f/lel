/*
Package controllers coordinates everything for LEL backend,
as well as managing other components, which are used to generate information

*/

package controllers

import (
	"errors"
	"github.com/c-f/lel/controllers/graph"
	"github.com/c-f/lel/controllers/media"
	"github.com/c-f/lel/controllers/misato"

	"github.com/c-f/lel/utils"

	"net/http"
	"os"

	"path/filepath"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"sort"
	//	"context"
)

// const for Websockets
const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (60 * time.Second * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 5 * time.Second
)

// APIHandler contains all handler funcs for api calls
type APIHandler struct {
	handler        http.Handler
	Notes          string
	ExecutablePath string

	// handles all interactions regarding graphs/meta
	graphHandler *graph.MetaGraph

	// handles logs and milestones
	misatoHandler *misato.Misato

	// handles media (img/video)
	mediaHandler *media.Media

	//TODO team mode
	//info
	auth map[string]string
	mode string
}

// NewAPIHandler creates a new router for the api calls
// also performs basic validation and creates the environment if necessary
func NewAPIHandler(notes, executablePath string, auth map[string]string, mode string) (*APIHandler, error) {
	notePath, err := filepath.Abs(notes)

	// Validating Options
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(notes, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// execuablePath
	if executablePath != "" {
		if ok := utils.FileExist(executablePath); !ok {
			log.Println("[*] Executable does not exist - please specify the complete path")
			return nil, os.ErrNotExist
		}
	}

	api := &APIHandler{
		Notes:          notePath,
		ExecutablePath: executablePath,
		auth:           auth,
		mode:           mode,
	}
	return api, nil
}

// AddGraphHandler adds a handler for graph interaction
func (api *APIHandler) AddGraphHandler(graphPath string) error {
	grapher := graph.New()
	log.Printf("[%s] %s MetaGraph", "api:graph", "initialize")

	for _, fullpath := range utils.FindMd(api.Notes) {
		log.Printf("\t[%s] %s %s", "api:graph", "index", fullpath)
		origin := utils.TrimPart(fullpath, api.Notes)
		grapher.IndexFile(fullpath, origin)

	}

	api.graphHandler = grapher
	return nil
}

// AddMisatoHandler returns a new misatohandler logger
func (api *APIHandler) AddMisatoHandler(misatoPath string) error {
	misato, err := misato.New(misatoPath)
	api.misatoHandler = misato
	return err
}

// AddImageHandler returns an image filehandler
func (api *APIHandler) AddUploadHandler(imagePath, videoPath string) error {
	handler, err := media.New(imagePath, videoPath)
	api.mediaHandler = handler
	return err

}

// Init initialize the APIHandler through routes registration
func (api *APIHandler) Init() {
	api.registerRoutes()
}

// Error codes returned by failures to parse an expression.
var (
	ErrNotCleanPath     = errors.New("Path is not clean")
	ErrHasNoMDExtension = errors.New("Wrong Extension - only MD files are accecpted")
	ErrFolderCreation   = errors.New("Error when creating a new folder !")
	ErrTagParsing       = errors.New("Tags could not be parsed !")
	ErrNoReference      = errors.New("No Reference found ")
	ErrNoContent        = errors.New("No Content found")
)

var (
	// for websockets
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// registerRoutes register all routes for the api server
func (api *APIHandler) registerRoutes() {
	root := mux.NewRouter()

	// --[API]--
	router := root.PathPrefix("/api/").Subrouter()

	// Single Note basic function
	router.HandleFunc(Routes.Get("Ok"), api.requireAuth(api.Ok))
	router.HandleFunc(Routes.Get("Get"), api.requireAuth(api.GetMDFile))
	router.HandleFunc(Routes.Get("Meta"), api.requireAuth(api.GetMDMeta))
	router.HandleFunc(Routes.Get("Open"), api.requireAuth(api.OpenMDFile))
	router.HandleFunc(Routes.Get("Remove"), api.requireAuth(api.RemoveMDFile))

	// Information about the notes and files
	router.HandleFunc(Routes.Get("Navigation"), api.requireAuth(api.GetNavArray))
	router.HandleFunc(Routes.Get("Folder"), api.requireAuth(api.GetFolderArray))
	router.HandleFunc(Routes.Get("Tags"), api.requireAuth(api.GetTagsArray))
	router.HandleFunc(Routes.Get("Metas"), api.requireAuth(api.GetMetasArray))
	router.HandleFunc(Routes.Get("Search"), api.requireAuth(api.SearchMDFiles))

	// keep track of changes on the filesystem. Sends notification to the frontend
	router.HandleFunc(Routes.Get("Watcher"), api.requireAuth(api.Watcher))

	//
	// --[Enable Features]--

	// Add GRAPH
	if api.graphHandler != nil {
		router.HandleFunc(Routes.Get("GraphView"), api.requireAuth(api.BuildComplete))
	}

	// Add Medias
	if api.mediaHandler != nil {
		router.HandleFunc(Routes.Get("ImageUpload"), api.requireAuth(api.ImageUpload))
		router.HandleFunc(Routes.Get("VideoUpload"), api.requireAuth(api.VideoUpload))
		router.HandleFunc(Routes.Get("Videos"), api.requireAuth(api.Videos))
		router.HandleFunc(Routes.Get("Images"), api.requireAuth(api.Images))
		router.HandleFunc(Routes.Get("MarkdownUpload"), api.requireAuth(api.MarkdownUpload))
	}

	// Add Loginformation
	if api.misatoHandler != nil {
		router.HandleFunc(Routes.Get("Milestone"), api.requireAuth(api.Milestone))
		router.HandleFunc(Routes.Get("Misato"), api.requireAuth(api.Misato))
		router.HandleFunc(Routes.Get("MisatoSearch"), api.requireAuth(api.MisatoSearch))
	}

	api.handler = router
}

// Ok returns the ok alive message
// this can be used to test the connection and especially the API token
func (api *APIHandler) Ok(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, "ok", http.StatusOK)
	return
}

// Videos return a list of video names based on requested criteria
// GET structue=fast,start=1572445498,stop=1572445498
func (api *APIHandler) Videos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		structure := r.URL.Query().Get("structure")
		switch structure {
		case "fast":
			infos := api.mediaHandler.ListVideos()
			JSONResponse(w, infos, http.StatusOK)
			return
		default:
			_start := r.URL.Query().Get("start")
			_stop := r.URL.Query().Get("stop")

			start, err := strconv.Atoi(_start)
			if err != nil {
				JSONErrorResponse(w, err)
				return
			}
			stop, err := strconv.Atoi(_stop)
			if err != nil {
				JSONErrorResponse(w, err)
				return
			}

			infos := api.mediaHandler.ListVideosByTime(start, stop)
			JSONResponse(w, infos, http.StatusOK)
			return
		}
	}
}

// Images returns a list of images
func (api *APIHandler) Images(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		infos := api.mediaHandler.ListImages()
		JSONResponse(w, infos, http.StatusOK)
	}
}

// GetMilestone returns all the milestones or create a new one
// GET structure=fast
// POST {}
func (api *APIHandler) Milestone(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		milestones := api.misatoHandler.GetMileStones()
		structure := r.URL.Query().Get("structure")

		switch structure {
		case "fast":
			JSONResponse(w, milestones, http.StatusOK)
			return
		default:
			log.Println("misato")
			sort.Stable(misato.MilestoneByTime{milestones})
			JSONResponse(w, milestones, http.StatusOK)
			return
		}

		return
	} else if r.Method == "POST" {
		log.Printf("[%s] %s Milestone", "debug:api:milestone", "create")
		//id := ctx.Get(r, "user_id").(string)
		id := "lel-server"

		decoder := json.NewDecoder(r.Body)
		var milestone misato.Milestone
		err := decoder.Decode(&milestone)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		err = api.misatoHandler.PutMileStone(milestone, id)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		JSONResponse(w, "ok", http.StatusOK)
	}

}

// BuildComplete builds the complete graph
//TODO Optimization  updating the graph is currently not implemented
func (api *APIHandler) BuildComplete(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s Graph", "api:graph", "building")

	// reset Graph
	api.graphHandler.Clear()

	// FindMd
	for _, fullpath := range utils.FindMd(api.Notes) {
		log.Printf("\t[%s] %s %s", "api:graph", "index", fullpath)
		origin := utils.TrimPart(fullpath, api.Notes)
		api.graphHandler.IndexFile(fullpath, origin)
	}

	export := api.graphHandler.Export()
	// TODO handle Errs and doublicate

	JSONResponse(w, export, http.StatusOK)
}

// BuildGraph builds the current state of the graph
// currently not supported
func (api *APIHandler) BuildGraph(w http.ResponseWriter, r *http.Request) {

	api.graphHandler.Clear()
	export := api.graphHandler.Export()
	// TODO handle Errs and doublicate

	JSONResponse(w, export, http.StatusOK)
}

// ImageUpload handles image uploaded as json
// PUT {}
func (api *APIHandler) ImageUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {
		absPath, err := api.mediaHandler.Image(r)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		log.Printf("[%s] uploading", "image:upload", absPath)
		newFile := utils.TrimPart(api.mediaHandler.ImageDir, absPath)
		JSONResponse(w, newFile, http.StatusOK)

		return
	}

	JSONResponse(w, r.Method, http.StatusMethodNotAllowed)
}

// VideoUpload handles video uploaded as multi form
func (api *APIHandler) VideoUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // maybe PUT
		absPath, err := api.mediaHandler.Video(r)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		log.Printf("[%s] uploading", "video:upload", absPath)
		newFile := utils.TrimPart(api.mediaHandler.VideoDir, absPath)
		JSONResponse(w, newFile, http.StatusOK)

		return
	}
	JSONResponse(w, r.Method, http.StatusMethodNotAllowed)
}

// MarkdownUpload handels upload
// TODO raceconditions !
func (api *APIHandler) MarkdownUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // maybe PUT
		inputPath := r.URL.Query().Get("path")

		log.Printf("[%s] got %s", "notes:put", inputPath)

		testPath, err := utils.GetAbsPath(api.Notes, inputPath)
		log.Printf("[%s] resolved to:  %s", "notes:get", testPath)
		if err != nil {
			JSONErrorResponse(w, ErrNotCleanPath)
			return
		}
		if ok := strings.HasSuffix(testPath, ".md"); !ok {
			JSONErrorResponse(w, ErrHasNoMDExtension)
			return
		}

		body := r.PostFormValue("content")
		if len(body) <= 0 {
			JSONErrorResponse(w, ErrNoContent)
		}

		err = utils.WriteTo(testPath, []byte(body))
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		JSONResponse(w, []byte("ok"), http.StatusOK)
		return
	}
}

// SearchMatch contains all information based on path, match and context ..
type SearchMatch struct {
	Origin string `json:"path"`
	Match  string `json:"match"`
	Key    string `json:"key"`

	Context *misato.CLILog `json:"ctx,omitempty"`
}

// MisatoSearch searches all misatoFiles for a specific searchword
func (api *APIHandler) MisatoSearch(w http.ResponseWriter, r *http.Request) {
	searchword := r.URL.Query().Get("search")
	structure := r.URL.Query().Get("structure")

	log.Printf("[%s] %s", "misato:search", searchword)

	results, err := api.misatoHandler.SearchCommands(searchword)
	if err != nil {
		JSONErrorResponse(w, err)
		return
	}
	matches := []SearchMatch{}
	for name, data := range results {
		for i, match := range data {
			sm := SearchMatch{
				Origin: name,
				Match:  match,
				Key:    fmt.Sprintf("%s-%d", name, i),
			}
			switch structure {
			case "fast":
			default:
				cli, err := misato.CLILogFromJson(match)
				if err != nil {
					log.Printf("[err][%s] %s", "misato:search", match)
					log.Println(err)
					continue
				}
				sm.Context = cli
			}
			matches = append(matches, sm)
		}

	}
	// results

	JSONResponse(w, matches, http.StatusOK)
	return
}

// Misato searches all misatoFiles for a specific timeframe
func (api *APIHandler) Misato(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		_start := r.URL.Query().Get("start")
		_stop := r.URL.Query().Get("stop")

		start, err := strconv.Atoi(_start)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}
		stop, err := strconv.Atoi(_stop)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		logs, err := api.misatoHandler.GetByTime(start, stop)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		JSONResponse(w, logs, http.StatusOK)
		return

	} else if r.Method == "POST" {
		log.Printf("[%s] %s Misato", "debug:api:misato", "create")
		//id := ctx.Get(r, "user_id").(string)

		id := "lel-server"

		decoder := json.NewDecoder(r.Body)
		var clilog misato.CLILog
		err := decoder.Decode(&clilog)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		err = api.misatoHandler.PutCommand(clilog, id)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}

		JSONResponse(w, "ok", http.StatusOK)
		return

	}
}

// SearchMDFiles searches all mdfiles based on the search String. Returns all match names
func (api *APIHandler) SearchMDFiles(w http.ResponseWriter, r *http.Request) {
	fullPaths := utils.FindMd(api.Notes)
	searchword := r.URL.Query().Get("search")
	log.Printf("[%s] %s", "notes:search", searchword)

	matches := []SearchMatch{}
	for _, filename := range fullPaths {
		data, err := utils.SearchWithPreviewInFile(filename, searchword)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}
		origin := utils.TrimPart(filename, api.Notes)
		for i, match := range data {
			matches = append(matches, SearchMatch{
				Origin: origin,
				Match:  match,
				Key:    fmt.Sprintf("%s-%d", origin, i),
			})
		}
	}
	JSONResponse(w, matches, http.StatusOK)
}

// GetTagsArray returns all tags as an array :)
func (api *APIHandler) GetTagsArray(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s]", "get:tags")
	fullPaths := utils.FindMd(api.Notes)
	treeStructure := NewTreeStructure()
	for _, path := range fullPaths {
		tags, err := utils.FindTagsInFile(path)
		log.Println("Tags", path, tags)
		if err != nil {
			JSONErrorResponse(w, ErrTagParsing)
			return
		}
		if len(tags) == 0 {
			continue
		}

		url := utils.TrimParts([]string{path}, api.Notes)
		treeStructure.AddItems(url[0], tags...)
	}
	JSONResponse(w, treeStructure.BuildAntGraph(), http.StatusOK)
}

// GetNavArray returns an json object containing the file structure of notes
func (api *APIHandler) GetNavArray(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s]", "get:nav")

	fullPaths := utils.FindMd(api.Notes)
	paths := utils.TrimParts(fullPaths, api.Notes)

	structure := r.URL.Query().Get("structure")
	switch structure {
	case "plain":
		JSONResponse(w, paths, http.StatusOK)
		return
	case "tree-only":
		treeStructure := NewTreeStructure()
		for _, path := range paths {
			treeStructure.AddItems(path, path)
		}
		JSONResponse(w, treeStructure.BuildAntGraph(), http.StatusOK)
	default:
		treeStructure := NewTreeStructure()
		for _, path := range paths {
			options := map[string]interface{}{}
			if meta := api.graphHandler.Get(path); meta != nil {
				options["meta"] = meta
			}
			treeStructure.AddItem(path, path, options)
		}
		JSONResponse(w, treeStructure.BuildAntGraph(), http.StatusOK)
		return
	}
}

// GetFolderArray returns an json array containing a plain structures
func (api *APIHandler) GetFolderArray(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s]", "get:folder")
	fullPaths := utils.FindMd(api.Notes)
	paths := utils.TrimParts(fullPaths, api.Notes)
	folders := []string{}
	seen := map[string]bool{
		// ignore local
		"./": true,
	}

	// add new paths
	for _, p := range paths {
		newPath := fmt.Sprintf("%s/", filepath.Dir(p))

		if _, ok := seen[newPath]; !ok {
			seen[newPath] = true
			folders = append(folders, newPath)
		}
	}

	JSONResponse(w, folders, http.StatusOK)
	return
}

/*
func(api *APIHandler) GetTodos(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		structure := r.URL.Query().Get("structure")
		switch structure{
		case "fast":
			// TODO
		default:
			log.Printf("[%s]", "get:metas")
			fullPaths := utils.FindMd(api.Notes)
			metas := []*graph.MetaInfo{}
			for _, fullPath := range fullPaths{
				id := utils.TrimPart(fullPath,api.Notes)
				info, err := graph.MetaInfoFromFile(fullPath, id)
				if err != nil{
					log.Printf("[!] error %s %s", "get:metas", err)
					continue
				}
				info.Todos
				//
				metas = append(metas, info)
			}
		}
	}
}
*/

// GetMetasArray returns all meta information from all the files
func (api *APIHandler) GetMetasArray(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Printf("[%s]", "get:metas")
		fullPaths := utils.FindMd(api.Notes)
		metas := []*graph.MetaInfo{}
		for _, fullPath := range fullPaths {
			id := utils.TrimPart(fullPath, api.Notes)
			info, err := graph.MetaInfoFromFile(fullPath, id)
			if err != nil {
				log.Printf("[!] error %s %s", "get:metas", err)
				continue
			}
			metas = append(metas, info)
		}

		JSONResponse(w, metas, http.StatusOK)
		return
	}
}

// GetMDMeta returns the Metadata from the requested path
func (api *APIHandler) GetMDMeta(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		inputPath := r.URL.Query().Get("path")

		log.Printf("[%s] got %s", "notes:get", inputPath)

		testPath, err := utils.GetAbsPath(api.Notes, inputPath)
		log.Printf("[%s] resolved to:  %s", "notes:get", testPath)
		if err != nil {
			JSONErrorResponse(w, ErrNotCleanPath)
			return
		}
		if ok := strings.HasSuffix(testPath, ".md"); !ok {
			JSONErrorResponse(w, ErrHasNoMDExtension)
			return
		}
		id := utils.TrimPart(testPath, api.Notes)

		meta, err := graph.MetaInfoFromFile(testPath, id)
		if err != nil {
			JSONErrorResponse(w, err)
			return
		}
		JSONResponse(w, meta, http.StatusOK)
		return
	}
}

// GetFile returns the md file in the notedir
func (api *APIHandler) GetMDFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	inputPath := r.URL.Query().Get("path")

	log.Printf("[%s] got %s", "notes:get", inputPath)
	testPath, err := utils.GetAbsPath(api.Notes, inputPath)
	log.Printf("[%s] resolved to:  %s", "notes:get", testPath)
	if err != nil {
		JSONErrorResponse(w, ErrNotCleanPath)
		return
	}
	if ok := strings.HasSuffix(testPath, ".md"); !ok {
		JSONErrorResponse(w, ErrHasNoMDExtension)
		return
	}

	FileResponse(w, r, testPath)
}

// RemoveMDFile removes the file for the LeL editor, either because through rm or mv the file
func (api *APIHandler)RemoveMDFile(w http.ResponseWriter, r *http.Request){
	inputPath := r.URL.Query().Get("path")
	shouldDeleted := r.URL.Query().Get("rm")

	log.Printf("[%s] got %s", "notes:remove", inputPath)
	testPath, err := utils.GetAbsPath(api.Notes, inputPath)
	if err != nil {
		JSONErrorResponse(w, ErrNotCleanPath)
		return
	}
	if ok := strings.HasSuffix(testPath, ".md"); !ok {
		JSONErrorResponse(w, ErrHasNoMDExtension)
		return
	}
	if ok := utils.FileExist(testPath); !ok{
		JSONErrorResponse(w, os.ErrNotExist)
		return 
	}
	if shouldDeleted != "" {
		log.Printf("[%s] got %s", "notes:delete", inputPath)
		if err = os.Remove(testPath); err != nil{
			JSONErrorResponse(w, err)
			return 
		}
		JSONResponse(w, Success("deleted"), http.StatusOK)
		return 
	}else{
		log.Printf("[%s] got %s", "notes:mv", inputPath)
		if err = os.Rename(testPath,testPath + ".d"); err != nil{
			JSONErrorResponse(w, err)
			return 
		}
		JSONResponse(w, Success("moved"), http.StatusOK)
		return 
	}

}

// EditFile opens the File in an editor and create
func (api *APIHandler) OpenMDFile(w http.ResponseWriter, r *http.Request) {
	inputPath := r.URL.Query().Get("path")
	editor := r.URL.Query().Get("editor")

	log.Printf("[%s] got %s", "notes:open", inputPath)
	testPath, err := utils.GetAbsPath(api.Notes, inputPath)
	log.Printf("[%s] resolved to:  %s", "notes:open", testPath)
	dir := filepath.Dir(testPath)

	if err != nil {
		JSONErrorResponse(w, ErrNotCleanPath)
		return
	}
	if ok := strings.HasSuffix(testPath, ".md"); !ok {
		JSONErrorResponse(w, ErrHasNoMDExtension)
		return
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		JSONErrorResponse(w, ErrFolderCreation)
		return
	}

	if editor != "false" {
		if api.ExecutablePath == "" {
			return
		}
		err = utils.OpenInEditor(api.ExecutablePath, testPath)
		if err != nil {
			JSONErrorResponse(w, err)
		}

	}

	JSONResponse(w, Success("opened"), http.StatusOK)
}

// ServeHTTP starts serving the API
func (api *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.handler.ServeHTTP(w, r)
}

// Websocket to start changing handler
func (api *APIHandler) Watcher(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %s", err)
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	api.writer(ws, api.Notes)
}

// Thanks bro https://github.com/gorilla/websocket/blob/master/examples/filewatch/main.go
// writer indicates that a file was changed
func (api *APIHandler) writer(ws *websocket.Conn, watchPath string) {

	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()
	now := time.Now()
	log.Printf("[%s] Starting Websocket", "socket:channel")

	for {
		select {
		case <-fileTicker.C:

			changes, hasChanged := utils.HasMDFilesChanged(now, watchPath)

			if hasChanged {
				log.Println("Websocket: Sending something. File has changed")
				now = time.Now()
				ws.SetWriteDeadline(time.Now().Add(writeWait))

				for _, fullpath := range changes {
					origin := utils.TrimPart(fullpath, api.Notes)
					// api.graphHandler.Index(fullpath, origin)

					msg := &UpdateMessage{
						MessageType: "fileupdate",
						Message:     origin,
					}
					bts, err := msg.ToJson()
					if err != nil {
						log.Println("Marshal was not possible")
					}
					if err := ws.WriteMessage(websocket.TextMessage, bts); err != nil {
						log.Println("Message could not be sent")
					}
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// UpdateMessage contains all data, which is sent to the client
type UpdateMessage struct {
	MessageType string `json:"type"`
	Message     string `json:"message"`
}

// ToJson returns the bytes representation of the UpdateMessage
func (um *UpdateMessage) ToJson() ([]byte, error) {
	bts, err := json.Marshal(um)
	return bts, err
}
