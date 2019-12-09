package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

var (
	invert      bool
	insensitive bool
	expr        string
	help        bool
	args        []string
	name        bool
	files       map[string]*strings.Builder
	mutex       *sync.Mutex
)

// Parse all the flag arguments and make the mutex lock
func init() {
	flag.BoolVar(&invert, "v", false, "Match all lines not containg string")
	flag.BoolVar(&insensitive, "i", false, "Match string with case insensitivity")
	flag.BoolVar(&help, "h", false, "Print this help message")
	flag.BoolVar(&name, "f", false, "Prints the file name for each of the files")
	flag.Parse()

	if len(flag.Args()) != 0 {
		expr = flag.Args()[0]
		args = flag.Args()[1:]

	}

	files = make(map[string]*strings.Builder)

	mutex = new(sync.Mutex)

	// This makes a case insensitive regexp
	if insensitive {
		expr = "(?i)(" + expr + ")"
	}

}

// Appends the text to the map value. This is thread-safe
func updateMapValue(fileName string, text string) {

	mutex.Lock()

	matchs, ok := files[fileName]

	if !ok {
		matchs = new(strings.Builder)
		files[fileName] = matchs
	}

	matchs.WriteString(text + "\n")

	mutex.Unlock()
}

// Scans the file for lines that match the search expression.
func grepFile(fileName string, expr string, scanner *bufio.Scanner) {

	var text string

	reg := regexp.MustCompile(expr)
	lineNum := 1
	for scanner.Scan() {

		text = scanner.Text()
		match := reg.MatchString(text)

		if match && !invert {
			loc := reg.FindAllStringIndex(text, len(text))

			for _, l := range loc {
				text = text[0:l[0]] + color.RedString(text[l[0]:l[1]]) + text[l[1]:]
			}

			if name {
				text = color.YellowString("%v: ", lineNum) + text
			}
			updateMapValue(fileName, text)
		} else if invert {

			if name {
				text = color.YellowString("%v: ", lineNum) + text
			}
			updateMapValue(fileName, text)
		}

		lineNum++
	}
}

// The main routine for the function
func main() {

	if help {
		fmt.Println("Usage: gogrep [flags] string [args...]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var wg sync.WaitGroup
	var names []string

	if len(args) == 0 {
		names = append(names, "STDIN")

		wg.Add(1)

		go func() {
			defer wg.Done()
			mutex.Lock()
			files["STDIN"] = new(strings.Builder)
			mutex.Unlock()
			grepFile("STDIN", expr, bufio.NewScanner(os.Stdin))
		}()

	} else {
		for _, fname := range args {
			b := path.Base(fname)
			names = append(names, b)

			file, err := os.Open(fname)

			if err != nil {
				log.Fatal(err.Error())
			}
			wg.Add(1)

			go func() {
				defer wg.Done()
				mutex.Lock()
				files[b] = new(strings.Builder)
				mutex.Unlock()
				grepFile(b, expr, bufio.NewScanner(file))
			}()
		}
	}

	wg.Wait()

	for _, n := range names {
		if name {
			fmt.Println(color.GreenString(n) + ": ")
		}
		fmt.Print(files[n].String())
	}
}
