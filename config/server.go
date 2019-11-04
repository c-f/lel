package config

import (
	"github.com/c-f/lel/utils"
	"os"

	"errors"
	"log"
)

// Errors regarding configs
var (
	ErrNoAuth = errors.New("No Auth specified")
)

// Server contains all information for the LEL server and API configuration
type ServerConfig struct {
	ListenURL string
	StaticDir string
	Editor    string

	Debug        bool
	EnableUpload bool

	TLS      bool
	CertPath string
	Hostname string

	// TODO not implemented
	Auth map[string]string `yaml`

	// Mode of the ServerConfig
	// single, team
	Mode string
}

// Valid checks if the config is valid
func (s *ServerConfig) Valid() error {
	if len(s.Auth) < 1 && s.Mode != "single" {
		return ErrNoAuth
	}
	return nil
}

// Prepare the ServerConfig
func (s *ServerConfig) Prepare() (err error) {
	if s.CertPath != "" {
		if s.CertPath, err = utils.GetAbsPath(s.CertPath, ""); err != nil {
			return
		}
		if err = os.MkdirAll(s.CertPath, 0755); err != nil {
			return
		}
	}
	if s.StaticDir != "" {
		if s.StaticDir, err = utils.GetAbsPath(s.StaticDir, ""); err != nil {
			return
		}
		empty, err := utils.DirIsEmpty(s.StaticDir)
		if err != nil {
			log.Printf("[%s] dir empty:  %s", err)
			return err
		}
		if empty {
			log.Printf("[%s] no files found in %s:  %s ", "!", "staticDir", s.StaticDir)
		}
	}
	return
}
