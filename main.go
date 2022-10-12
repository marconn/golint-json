package main

import (
	"flag"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/output"
	"os"
)

func main() {
	var path string

	// create default console styler
	out := output.NewConsoleOutput(true, nil)

	flag.StringVar(&path, "path", "", "Directory path containing JSON files")
	flag.Parse()

	if path == "" {
		out.Writeln("<error>Please provide a path</error>")
		os.Exit(1)
	}

	if _, err := isValidDirectory(path); err != nil {
		out.Writeln(fmt.Sprintf("<error>%v</error>", err))
		os.Exit(1)
	}

	files := getJSONFiles(path)

	for _, file := range files {
		out.Writeln(fmt.Sprintf("<fg=yellow;options=underscore>Validating: %s</>", file))

		err := lint(file)
		if err != nil {
			out.Writeln(fmt.Sprintf("<fg=red>\u274C  %v</>", err))
		} else {
			out.Writeln(fmt.Sprintf("<fg=green>\u2713  OK</>"))
		}
	}
}
