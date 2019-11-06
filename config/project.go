package config

import (
	"github.com/c-f/lel/utils"

	"errors"
)

var (
	ErrNoDirSpecified = errors.New("No Dir specified")
)

// Project contains all information about the project paths
// this includes paths, where project relevant information are stored
type ProjectConfig struct {
	NotesDir string

	ImageDir  string
	VideoDir  string
	MisatoDir string
	GraphDir  string
}

// Valid checks if the config is valid
func (p *ProjectConfig) Valid() (err error) {
	if p.NotesDir == "" {
		return ErrNoDirSpecified
	}
	return nil
}

// Prepare the path
func (p *ProjectConfig) Prepare() (err error) {
	if p.NotesDir != "" {
		if p.NotesDir, err = utils.GetAbsPath(p.NotesDir, ""); err != nil {
			return
		}
	}
	if p.ImageDir != "" {
		if p.ImageDir, err = utils.GetAbsPath(p.ImageDir, ""); err != nil {
			return
		}
	}
	if p.VideoDir != "" {
		if p.VideoDir, err = utils.GetAbsPath(p.VideoDir, ""); err != nil {
			return
		}
	}
	if p.MisatoDir != "" {
		if p.MisatoDir, err = utils.GetAbsPath(p.MisatoDir, ""); err != nil {
			return
		}
	}
	if p.GraphDir != "" {
		if p.GraphDir, err = utils.GetAbsPath(p.GraphDir, ""); err != nil {
			return
		}
	}
	return
}

// NewProjectConfig creates a new ProjectConfig, with a minimum requirements of specifying the notesDir
func NewProjectConfig(notesDir string) *ProjectConfig {
	notesDir, _ = utils.GetAbsPath(notesDir, "")
	return &ProjectConfig{
		NotesDir: notesDir,
	}
}
