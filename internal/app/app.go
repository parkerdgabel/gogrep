package app

import (
	"time"

	"github.com/urfave/cli"
)

// NewApp returns pointer to a new cli.App.
func NewApp() *cli.App {
	return &cli.App{
		Name: "Sieve",
		Version: "0.0.1",
		Authors: []*Author{
			&cli.Author{
				Name:  "Parker Gabel",
				Email: "parker.d.gabel@gmail.com",
			},
		},
		Compiled: time.Now(),
		Usage: "Sieve recursively searchs the current directory or path argument for matchs to a regex pattern.",
		Flags: []Flag,
		
	}
}
