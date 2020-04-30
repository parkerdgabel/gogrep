package main

import "github.com/urfave/cli"

type caseSearch int

const (
	Sensitive caseSearch = iota
	Insensitive
	Smart
)

func newCase(c *cli.Context) caseSearch {
	var caseS caseSearch

	if c.IsSet("smart-case") {
		caseS = Smart
	} else if c.IsSet("ignore-case") {
		caseS = Insensitive
	} else {
		caseS = Sensitive
	}

	return caseS
}

type output struct {
	Color             bool
	Count             bool
	Unordered         bool
	Stats             bool
	Headers           bool
	LineNums          bool
	Quiet             bool
	Destination       string
	FilesWithMatch    bool
	FilesWithoutMatch bool
}

func newOutput(c *cli.Context) *output {
	o := output{}

	if c.IsSet("quiet") {
		o.Quiet = true
	} else if c.IsSet("count") {
		o.Count = true
	} else if c.IsSet("list-files-with-match") {
		o.FilesWithMatch = true
	} else if c.IsSet("list-files-without-match") {
		o.FilesWithoutMatch = true
	} else {
		// Booleans are set to false by default
		if !c.IsSet("no-file-headers") {
			o.Headers = true
		}

		if !c.IsSet("no-color") {
			o.Color = true
		}

		if c.IsSet("line-nums") {
			o.LineNums = true
		}

		if c.IsSet("unordered") {
			o.Unordered = true
		}

		if c.IsSet("stats") {
			o.Stats = true
		}
	}

	o.Destination = c.String("output")

	return &o
}

type files struct {
	IncludedFiles []string
	IncludedPaths []string
	ExcludedFiles []string
	ExcludedPaths []string
}

func newFiles(c *cli.Context) *files {
	f := &files{}

	if c.IsSet("files") {
		f.IncludedFiles = c.StringSlice("files")
	}

	if c.IsSet("path") {
		f.IncludedPaths = c.StringSlice("path")
	}

	if c.IsSet("excluded-files") {
		f.ExcludedFiles = c.StringSlice("excluded-files")
	}

	if c.IsSet("excluded-path") {
		f.ExcludedPaths = c.StringSlice("excluded-path")
	}

	return f
}

func newPattern(c *cli.Context) string {
	var p string

	if c.IsSet("regex") {
		p = c.String("regex")
	} else {
		p = c.Args().First()
	}

	return p
}

// Option implements the configuration for a search in Sieve.
type Option struct {
	caseS    caseSearch
	output   *output
	fileOpts *files
	pattern  string
	invert   bool
	recurse  bool
	// If gitignore or binarary are true then that means to search them.
	gitignore bool
	binary    bool
}

// IsSearchCaseSensitive returns true if the search is supposed to be case sensitive.
func (o Option) IsSearchCaseSensitive() bool {
	return o.caseS == Sensitive
}

// IsSearchCaseInsensitive returns true it the seach is supposed to be case insensitive.
func (o Option) IsSearchCaseInsensitive() bool {
	return o.caseS == Insensitive
}

// IsSearchSmartCase returns true if the search is supposed to be a smart case search.
func (o Option) IsSearchSmartCase() bool {
	return o.caseS == Smart
}


// IsSearchInverted returns true if the search should produce inverted matchings.
func (o Option) IsSearchInverted() bool {
	return o.invert
}

// IsSearchRecursive returns true if the search should recursively search directories.
func (o Option) IsSearchRecursive() bool {
	return o.recurse
}

// ShouldSearchgitignore returns true if the search should search files matched in a gitignore file.
func (o Option) ShouldSearchGitignore() bool {
	return o.gitignore
}

// ShouldSearchBinaries returns true if the search should seach binary files.
func (o Option) ShouldSearchBinaries() bool {
	return o.binary
}

// HasIncludedFiles returns true if there are any included files to search.
func (o Option) HasIncludedFiles() bool {
	return len(o.fileOpts.IncludedFiles) != 0
}

// HasIncludedFiles returns true if there are any included paths to seach.
func (o Option) HasIncludedPaths() bool {
	return len(o.fileOpts.IncludedPaths) != 0
}

// HasExcludedFiles returns true if there are any excluded files to search.
func (o Option) HasExcludedFiles() bool {
	return len(o.fileOpts.ExcludedFiles) != 0
}

//  HasExcludedPaths returns true  if there are any paths to exclude.
func (o Option) HasExcludedPaths() bool {
	return len(o.fileOpts.ExcludedPaths) != 0
}

// IsOutputColored returns true if the output is suppoesed to be supressed.
func (o Option) IsOutputQuiet() bool {
	return o.output.Quiet
}

// IsOutputColored returns true if the output should be colored.
func (o Option) IsOutputColored() bool {
	return o.output.Color
}

// IsOutputSTDOUT returns true if the output destination is STDOUT.
func (o Option) IsOutputSTDOUT() bool {
	return o.output.Destination == "STDOUT"
}

// IsOutputUnordered returns true if the output should be unordered.
func (o Option) IsOutputUnordered() bool {
	return o.output.Unordered
}

// ShouldOutputContainStats returns true if the output should contain a stats output at the end.
func (o Option) ShouldOutputContainStats() bool {
	return o.output.Stats
}

// ShouldOutputContainHeaders returns true if the output should contain file headers.
func (o Option) ShouldOutputContainHeaders() bool {
	return o.output.Headers
}

// ShouldOutputContainLineNumbers return true if the the output should have line numbers.
func (o Option) ShouldOutputContainLineNumbers() bool {
	return o.output.LineNums
}

// IsOutputFilesWithMatch returns true if only files with matchs should be printed
func (o Option) IsOutputFilesWithMatch() bool {
	return o.output.FilesWithMatch
}

// IsOutputFilesWithoutMatch returns true if only files without matchs should be printed
func (o Option) IsOutputFilesWithoutMatch() bool {
	return o.output.FilesWithoutMatch
}

type optionsError struct {
	s string
}

func (e *optionsError) Error() string {
	return "Options error: " + e.s
}

func newOptionsError(err string) *optionsError {
	return &optionsError{s: err}
}

func checkListFilesContradiction(c *cli.Context) error {
	if c.IsSet("list-files-with-match") && c.IsSet("list-files-without-match") {
		return newOptionsError("List files contradiction.")
	}
	return nil
}

func checkContradictions(c *cli.Context) error {
	var err error
	err = checkListFilesContradiction(c)
	if err != nil {
		return err
	}

	return nil
}

// NewOption returns a pointer to an Option struct configured by the command
// line context. Returns an error if any errors in configuration.
func NewOption(c *cli.Context) (*Option, error) {
	err := checkContradictions(c)
	if err != nil {
		return nil, err
	}

	var o Option

	o.caseS = newCase(c)
	o.output = newOutput(c)
	o.fileOpts = newFiles(c)
	o.pattern = newPattern(c)

	if c.IsSet("invert-match") {
		o.invert = true
	}

	if !c.IsSet("no-recurse") {
		o.recurse = true
	}

	if c.IsSet("gitignore") {
		o.gitignore = true
	}

	if c.IsSet("binary") {
		o.binary = true
	}

	return &o, nil
}
