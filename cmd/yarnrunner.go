//go:build example
// +build example

// Copyright 2021 Josh Deprez
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The yarnrunner binary runs a yarnc+string table combo as a "text game" on
// the terminal.
//
// Quick usage from the root of the repo:
//
//	go run -tags example cmd/yarnrunner.go \
//	    --program=testdata/Example.yarn.yarnc \
//	    --strings=testdata/Example.yarn.csv
//
// The "example" build tag is used to prevent this being installed to ~/go/bin
// if you use the go get command. If for some reason you want to install it to
// your ~/go/bin, use `go install -tags example cmd/yarnrunner.go` or similar.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/DrJosh9000/yarn"
)

func main() {
	yarncFilename := flag.String("program", "", "File name of program (e.g. Example.yarn.yarnc)")
	csvFilename := flag.String("strings", "", "File name of string table (e.g. Example.yarn.csv)")
	startNode := flag.String("start", "Start", "Name of the node to run")
	langCode := flag.String("lang", "en-AU", "Language tag (BCP 47)")
	flag.Parse()

	program, stringTable, err := yarn.LoadFiles(*yarncFilename, *csvFilename, *langCode)
	if err != nil {
		log.Fatalf("Loading files: %v", err)
	}

	vm := &yarn.VirtualMachine{
		Program: program,
		Handler: &dialogueHandler{
			stringTable: stringTable,
		},
		Vars: make(yarn.MapVariableStorage),
	}
	if err := vm.Run(*startNode); err != nil {
		log.Printf("Yarn VM error: %v", err)
	}
}

// dialogueHandler implements yarn.DialogueHandler by playing the lines and
// options on the terminal.
type dialogueHandler struct {
	stringTable *yarn.StringTable

	yarn.FakeDialogueHandler // implements remaining methods
}

func (h *dialogueHandler) Line(line yarn.Line) error {
	text, err := h.stringTable.Render(line)
	if err != nil {
		return err
	}
	fancyPrintln(text)
	fmt.Print("(Press ENTER to continue)")
	fmt.Scanln()
	// This next string is VT100 for "move to the first column, go up a line,
	// and erase it" (erasing the Press ENTER message).
	fmt.Print("\r\033[A\033[2K")
	return nil
}

func (h *dialogueHandler) Options(opts []yarn.Option) (int, error) {
	fmt.Println("Choose:")
	for _, opt := range opts {
		text, err := h.stringTable.Render(opt.Line)
		if err != nil {
			return 0, err
		}
		fmt.Printf("%d: ", opt.ID)
		fancyPrintln(text)
	}
	var choice int
	for {
		fmt.Print("Enter the number corresponding to your choice: ")
		if _, err := fmt.Scanln(&choice); err != nil {
			continue
		}
		break
	}
	return choice, nil
}

// fancyPrintln prints an attributed string with ANSI escape sequences that
// apply formatting, corresponding to the BBCode-style tags from the original
// yarn file.
//
// This is needed because fmt.Println(text) would only print the string, not the
// escape sequences needed to apply bold, underline, colour, etc.
func fancyPrintln(text *yarn.AttributedString) {
	str := text.String()
	cursor := 0
	open := make(map[*yarn.Attribute]struct{})
	text.ScanAttribEvents(func(pos int, atts []*yarn.Attribute) {
		if pos > cursor {
			// Print text to this position
			fmt.Print(str[cursor:pos])
			cursor = pos
		}
		// atts either start here, end here, or both.
		var newopen []string
		closed := false
		for _, a := range atts {
			if a.Start == a.End {
				// has no net effect; skip it
				continue
			}
			if a.Start == pos {
				open[a] = struct{}{}
				newopen = append(newopen, a.Name)
			}
			if a.End == pos {
				closed = true
				delete(open, a)
			}
		}
		if closed {
			// Clear all character attributes
			fmt.Print("\033[m")
			for a := range open {
				// Print escape codes for all the remaining open attributes
				fmt.Print(escapes[a.Name])
			}
		} else {
			// Only print escape codes for newly opened attributes
			for _, n := range newopen {
				fmt.Print(escapes[n])
			}
		}

	})
	// Print remainder
	fmt.Println(str[cursor:])
}

// Maps style tags to ANSI escape sequences. Printing these changes the style or
// colour. See https://en.wikipedia.org/wiki/ANSI_escape_code.
var escapes = map[string]string{
	"b":         "\033[1m",
	"bold":      "\033[1m",
	"lowint":    "\033[2m",
	"u":         "\033[4m",
	"underline": "\033[4m",
	"blink":     "\033[5m",
	"reverse":   "\033[7m",
	"invisible": "\033[8m",
	"strike":    "\033[9m",
	"red":       "\033[31m",
	"green":     "\033[32m",
	"yellow":    "\033[33m",
	"blue":      "\033[34m",
	"purple":    "\033[35m",
	"cyan":      "\033[36m",
	"gray":      "\033[37m",
	"grey":      "\033[37m",
	"white":     "\033[97m",
}
