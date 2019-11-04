package controllers

// Genric error

// md output

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type Msg struct {
	Msg     string
	Success bool
}

func Success(msg string) *Msg {
	return &Msg{
		Msg:     msg,
		Success: true,
	}
}

// copied from https://github.com/gophish/gophish/blob/master/controllers/api/response.go
// JSONResponse attempts to set the status code, c, and marshal the given interface, d, into a response that
// is written to the given ResponseWriter.
func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		//log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

// JSONErrorResponse returns a static response Object
func JSONErrorResponse(w http.ResponseWriter, err error) {
	d := &Msg{
		Msg:     err.Error(),
		Success: false,
	}

	JSONResponse(w, d, http.StatusInternalServerError)
}

// MDResponse handles the ServeFile
// most likely be used for md files
func FileResponse(w http.ResponseWriter, r *http.Request, filePath string) {
	http.ServeFile(w, r, filePath)
}
