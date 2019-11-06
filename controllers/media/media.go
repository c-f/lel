package media

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/c-f/lel/utils"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	ErrWrongFormat       = errors.New("Wrong Format ")
	ErrWrongExtension = errors.New("Wrong format")
	ErrFileAlreadyExists = errors.New("File already exist")
)

type VideoInfo struct {
	Name  string `json:"name"`
	Start int    `json:"start"`
	Stop  int    `json:"stop"`
	// duration in seconds
	Duration int `json:"duration"`
}

// 1568723655748-1568723660188.webm
func ConvertToVideoInfo(raw string) (*VideoInfo, error) {
	parts := strings.Split(strings.TrimSuffix(raw, ".webm"), "-")
	if len(parts) != 2 {
		return nil, ErrWrongFormat
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	stop, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	duration := (stop - start)

	return &VideoInfo{
		Name:     raw,
		Start:    start,
		Stop:     stop,
		Duration: duration,
	}, nil
}

// Image represents the objects from where files are uploaded and gathered
type Media struct {
	ImageDir string
	VideoDir string
	GraphDir string
}

// New creates a new Media for uploading and serving images from a specific folder
func New(imagePath, videoPath, graphDir string) (im *Media, err error) {
	if err = os.MkdirAll(imagePath, 0755); err != nil {
		return
	}
	if err = os.MkdirAll(videoPath, 0755); err != nil {
		return
	}
	if err = os.MkdirAll(graphDir, 0755); err != nil {
		return
	}

	im = &Media{
		ImageDir: imagePath,
		VideoDir: videoPath,
		GraphDir: graphDir,
	}
	return im, err
}

// TODO implement
func (im *Media) Get(filename string) string {
	return "notImplementedYet"
}

// ListVideosByTime
func (im *Media) ListVideosByTime(start, stop int) []VideoInfo {
	result := []VideoInfo{}
	infos := im.ListVideos()
	for _, info := range infos {
		if start < info.Start && stop > info.Stop {
			result = append(result, info)
		}
	}

	return result
}

// ListVideos returns a list of videos
func (im *Media) ListVideos() []VideoInfo {
	files := utils.FindExtension(im.VideoDir, ".webm")
	files = utils.TrimParts(files, im.VideoDir)

	infos := []VideoInfo{}
	for _, file := range files {
		info, err := ConvertToVideoInfo(file)
		if err != nil {
			continue
		}
		infos = append(infos, *info)
	}

	return infos
}

// ListImages returns a list of uploaded images
func (im *Media) ListImages() []string {
	files := utils.FindExtension(im.ImageDir, ".png")
	files = utils.TrimParts(files, im.ImageDir)

	return files
}

// ListGraphs returns a list of all saved graphs
func (im *Media) ListGraphs() []string{
	files := utils.FindExtension(im.GraphDir, ".graph.json")
	files = utils.TrimParts(files, im.GraphDir)

	return files
}

// Upload Graph
func (im *Media)Graph(r *http.Request)(fileName string, err error){
	name := r.URL.Query().Get("name")

	testPath, err := utils.GetAbsPath(im.GraphDir, name)
	if err != nil {
		return "", err
	}
	if ok := strings.HasSuffix(testPath, ".graph.json"); !ok {
		return "", ErrWrongExtension
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return "", err
	}
	if len(body) <= 0 {
		return "", errors.New("empty")
	}
	err = SaveBts(body, testPath)
	if err != nil {
		return "", err 
	}

	return testPath, nil
}

// UploadVideo
func (im *Media) Video(r *http.Request) (fileName string, err error) {

	// get TimeStamp
	_timeStart := r.PostFormValue("timestamp-start")
	_timeStop := r.PostFormValue("timestamp-end")

	start, err := strconv.Atoi(_timeStart)
	if err != nil {
		fmt.Println("not a integer!")
		err = ErrWrongFormat
		return
	}
	stop, err := strconv.Atoi(_timeStop)
	if err != nil {
		fmt.Println("not a integer!")
		err = ErrWrongFormat
		return
	}

	file, info, err := r.FormFile("uploadFile")
	if err != nil {
		return
	}
	defer file.Close()

	// io.Copy
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	// todo
	// ErrWrongFormat
	// Content-Type
	fmt.Println(info.Header)
	// fileBytes
	// video/mp4

	// get AbsPath
	fileName, err = utils.GetAbsPath(im.VideoDir, fmt.Sprintf("%d-%d.webm", start/1000, stop/1000))
	if err != nil {
		return
	}
	//
	if utils.FileExist(fileName) {
		err = ErrFileAlreadyExists
		return
	}

	err = SaveBts(fileBytes, fileName)

	return fileName, err
}

// Image handles the upload of the request
func (im *Media) Image(r *http.Request) (fileName string, err error) {
	var data ImageData

	// decode to json
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	// decode to img
	img, err := data.GetPicture()
	if err != nil {
		log.Println("Picture")
		return
	}
	// get AbsPath
	fileName, err = utils.GetAbsPath(im.ImageDir, fmt.Sprintf("%s.png", data.Title))
	if err != nil {
		return
	}
	if utils.FileExist(fileName) {
		err = ErrFileAlreadyExists
		return
	}

	// save
	err = SavePng(img, fileName)

	return fileName, err
}
