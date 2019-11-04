package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// NewStaticHandler returns a new Static Fileserver
func NewStaticHandler(staticPath string) http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath)))
}

// NewImageHandler creates a new http.Handler for serving pictures
func NewSuffixHandler(path, suffix string) func(string) http.Handler {
	if path == "" {
		return func(p string) http.Handler {
			return http.HandlerFunc(EmptyHandler)
		}
	}
	fs := SuffixFS{
		fs:              http.Dir(path),
		whiteListSuffix: suffix,
	}
	return func(p string) http.Handler {
		return http.StripPrefix(p, http.FileServer(fs))
	}
}

var (
	ErrNoDirectoryListing = errors.New("No Directory Listing ")
	ErrOnlyValidExtension = errors.New("Only request allowed extensions")
)

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sorry not specified")
	return
}

type SuffixFS struct {
	fs              http.FileSystem
	whiteListSuffix string
}

func (nfs SuffixFS) Open(path string) (http.File, error) {

	if ok := strings.HasSuffix(path, nfs.whiteListSuffix); !ok {
		return nil, ErrOnlyValidExtension
	}
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		return nil, ErrNoDirectoryListing
	}

	return f, nil
}
