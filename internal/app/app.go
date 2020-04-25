package app

import (
	"github.com/urfave/cli"
	"time"
)

// NewApp returns pointer to a new cli.App.
func NewApp() *cli.App {
	return &cli.App{
		Name:    "Sieve",
		Version: "0.0.1",
		Authors: []*Author{
			&cli.Author{
				Name:  "Parker Gabel",
				Email: "parker.d.gabel@gmail.com",
			},
		},
		Compiled:    time.Now(),
		Description: "Sieve recursively searchs the current directory or path argument for matchs to a regex pattern.",
		Flags: []Flag{
			&cli.BoolFlag{
				Name:  "ignore-case, i",
				Usage: "Perform a case-insensitive search",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "invert-match, v",
				Usage: "Invert matchings. Prints lines that do not match the patterns",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "count, c",
				Usage: "Shows the number of lines that match the given patterns. Suppresses the normal output",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "paths, P",
				Usage: "Shows the path of all files that contain matchs. Suppresses the normal output",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "line-nums,  l",
				Usage: "Prints the line numbers with the matchs.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "unordered, u",
				Usage: "Sieve is a concurrent program. The matchs may come in any order. This flag prints the  matchs in the order they are recieved.",
				Value: false,
			},
		},
	}
}
