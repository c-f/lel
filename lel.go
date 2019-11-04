package lel

import (
	"github.com/c-f/lel/config"
	"github.com/c-f/lel/controllers"

	"log"
)

// NewLel returns the server
func New(c *config.Config) *controllers.Server {

	s, err := controllers.NewServer(&c.Server, &c.Project)
	if err != nil {
		log.Fatal(err)
	}
	return s
}
