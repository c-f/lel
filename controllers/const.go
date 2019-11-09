package controllers

import (
	"fmt"
	"html/template"
)

const (
	// CERTPREFIX is the prefix for the https certificates
	CERTPREFIX = "LeL-server"
)

var (
	// Routes defines all endpoints for the backend, easy accessible via a keyword
	Routes = &Mapping{
		Config: map[string]template.JSStr{
			// single note
			"Get":  template.JSStr("/notes/get"),
			"Meta": template.JSStr("/notes/meta"),
			"Open": template.JSStr("/notes/open"),
			"Remove": template.JSStr("/notes/remove"),

			// Core information
			"Ok":         template.JSStr("/core/ok"),
			"Stats":      template.JSStr("/core/stats"),

			"Navigation": template.JSStr("/core/nav"),
			"Folder":     template.JSStr("/core/folder"),
			"Tags":       template.JSStr("/core/tags"),
			"Metas":      template.JSStr("/core/metas"),
			"Milestone":  template.JSStr("/core/milestones"),
			"Videos":     template.JSStr("/core/videos"),
			"Images":     template.JSStr("/core/images"),
			"Graphs": 	  template.JSStr("/core/graphs"),

			// search
			"Search":  template.JSStr("/core/search"),
			"Watcher": template.JSStr("/core/channel"),

			// features
			"GraphView":      template.JSStr("/feat/graph/get"),

			"ImageUpload":    template.JSStr("/feat/upload/image"),
			"VideoUpload":    template.JSStr("/feat/upload/video"),
			"MarkdownUpload": template.JSStr("/feat/upload/md"),
			"GraphUpload":    template.JSStr("/feat/upload/graph"),

			"Misato":       template.JSStr("/feat/misato"),
			"MisatoSearch": template.JSStr("/feat/misato/search"),

			// Frontend path
			"Frontend": template.JSStr("/static/app-prod.js"),
		},
	}
)

// Mapping defines a simple interface to access and write a config
type Mapping struct {
	Config map[string]template.JSStr
}

// Get retrieves a string based on the key, when existed
func (m *Mapping) Get(key string) string {

	if template, ok := m.Config[key]; ok {
		return fmt.Sprintf("%s", template)
	}
	return ""
}

// Set adds the val to the config with key as the key
func (m *Mapping) Set(key, val string) {
	m.Config[key] = template.JSStr(val)
}
