package controllers

import (
	"fmt"
)

const (
	CERTPREFIX = "LeL-server"
)

var (
	Routes = &Mapping{
		Config: map[string]string{
			// single note
			"Get":  "/notes/get",
			"Meta": "/notes/meta",
			"Open": "/notes/open",
			"Remove": "/notes/remove",

			// Core information
			"Ok":         "/core/ok",
			"Stats":      "/core/stats",
			"Navigation": "/core/nav",
			"Folder":     "/core/folder",
			"Tags":       "/core/tags",
			"Metas":      "/core/metas",
			"Milestone":  "/core/milestones",
			"Videos":     "/core/videos",
			"Images":     "/core/images",
			"Graphs": 	  "/core/graphs",

			// search
			"Search":  "/core/search",
			"Watcher": "/core/channel",

			// features
			"GraphView":      "/feat/graph/get",

			"ImageUpload":    "/feat/upload/image",
			"VideoUpload":    "/feat/upload/video",
			"MarkdownUpload": "/feat/upload/md",
			"GraphUpload":    "/feat/upload/graph",

			"Misato":       "/feat/misato",
			"MisatoSearch": "/feat/misato/search",
		},
	}
)

type Mapping struct {
	Config map[string]string
}

func (m *Mapping) Get(key string) string {

	if template, ok := m.Config[key]; ok {
		return fmt.Sprintf("%s", template)
	}
	return ""
}
func (m *Mapping) Set(key, val string) {
	m.Config[key] = val
}
