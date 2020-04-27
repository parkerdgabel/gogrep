// Package option implements an Option struct to conifigure the parameters for
// a search in Sieve.
package option

import "github.com/urfave/cli"

type caseSearch int

const (
	Sensitive caseSearch = iota
	Insensitive
	Smart
)

type output struct {
	Color       bool
	Count       bool
	Unordered   bool
	Stats       bool
	Headers     bool
	LineNums    bool
	Quiet       bool
	Destination string
}

type files struct {
	IncludedFiles []string
	IncludedPaths []string
	ExcludedFiles []string
	ExcludedPaths []string
}

// Option implements the configuration for a search in Sieve.
type Option struct {
	Case      caseSearch
	Output    output
	Files     files
	Pattern   string
	Invert    bool
	Recurse   bool
	Gitignore bool
	Binary    bool
}

// NewOption returns a pointer to an Option struct configured by the command
// line context. Returns an error if any errors in configuration.
func NewOption(c *cli.Context) (*Option, error){
	
}
