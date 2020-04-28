package main

import (
	"github.com/urfave/cli"
	"time"
)

// NewApp returns pointer to a new cli.App.
func NewApp() *cli.App {
	return &cli.App{
		Name:    "Sieve",
		Version: "0.0.1",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Parker Gabel",
				Email: "parker.d.gabel@gmail.com",
			},
		},
		Compiled:    time.Now(),
		Description: "Sieve recursively searchs the current directory or path argument for matchs to a regex pattern.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "ignore-case",
				Aliases: []string{"i"},
				Usage: "Perform a case-insensitive search",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "smart-case",
				Aliases: []string{"s"},
				Usage: "Performs a smart-case search of `PATTERN`.",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "invert-match",
				Aliases: []string{"x"},
				Usage: "Invert matchings. Prints lines that do not match the patterns",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "count",
				Aliases: []string{"c"},
				Usage: "Shows the number of lines that match the given patterns. Suppresses the normal output",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "line-nums",
				Aliases: []string{"n"},
				Usage: "Prints the line numbers with the matchs.",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "unordered",
				Aliases: []string{"u"},
				Usage: "Sieve is a concurrent program. The matchs may come in any order. This flag prints the matchs in the order they are recieved.",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "gitignore",
				Aliases: []string{"g"},
				Usage: "Seach all files matched in a gitignore file. By default, sieve will not search files matched by the gitignore file.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "no-color",
				Usage: "Disables colorized output.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "no-recurse",
				Usage: "Disables recursive searching for the pattern. Sieve will just search the files in the given directory or the working directory if none is given.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "binary",
				Usage: "Searches files that appear to be  binaries.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "list-files-with-match",
				Aliases: []string{"l"},
				Usage: "Lists all the files that contain a match.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "list-files-without-match",
				Aliases: []string{"L"},
				Usage: "Lists all the files that do not contain a match",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "no-file-headers",
				Usage: "Supress file headers in the output.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "quiet",
				Aliases: []string{"q"},
				Usage: "Suppresses normal output. Returns with 0 exit code if a match was found.",
				Value: false,
			},
			&cli.BoolFlag{
				Name: "stats",
				Usage: "Prints the statistics for the search at the end of the output.",
				Value: false,
			},
			&cli.StringFlag{
				Name:  "regex",
				Aliases: []string{"e"},
				Usage: "The pattern the search will match.",
			},
			&cli.StringFlag{
				Name: "output",
				Aliases: []string{"o"},
				Usage: "Specifies the output destination, either a FILE or a network connection.",
				Value: "STDOUT",
			},
			&cli.StringSliceFlag{
				Name: "files",
				Aliases: []string{"f"},
				Usage: "Search all files matching `GLOB`(comma-delimited list).",
			},
			&cli.StringSliceFlag{
				Name: "path",
				Aliases: []string{"P"},
				Usage: "Search all paths matching `GLOB`(comma-delimited list).",
			},
			&cli.StringSliceFlag{
				Name: "exclude-files",
				Usage: "Excludes all files matching `GLOB` from the search(comma-delimited list).",
			},
			&cli.StringSliceFlag{
				Name: "exclude-path",
				Usage: "Excludes all paths matching `GLOB` from the search.(comma-delimited list)",
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
