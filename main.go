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
	args        []string
	files       map[string]*strings.Builder
	mutex       *sync.Mutex
)

func init() {
	flag.BoolVar(&invert, "v", false, "Match all lines not containg REGEXP")
	flag.BoolVar(&insensitive, "i", false, "Match REGEXP with case insensitivity")
	flag.StringVar(&expr, "e", "", "The REGEXP to match")
	flag.Parse()
	args = flag.Args()

	files = make(map[string]*strings.Builder)

	mutex = new(sync.Mutex)

	if insensitive {
		expr = "(?i)(" + expr + ")"
	}
}

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

func grepFile(fileName string, expr string, scanner *bufio.Scanner) {

	var text string

	reg := regexp.MustCompile(expr)

	for scanner.Scan() {

		text = scanner.Text()
		match := reg.MatchString(text)

		if match && !invert {
			loc := reg.FindAllStringIndex(text, len(text))

			for _, l := range loc {
				text = text[0:l[0]] + color.RedString(text[l[0]:l[1]]) + text[l[1]:]
			}
			updateMapValue(fileName, text)
		} else if invert {
			updateMapValue(fileName, text)
		}

	}
}

func main() {
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
		fmt.Print(files[n].String())
	}
}
