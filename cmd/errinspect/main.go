package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Tahler/go-linters/pkg/lint"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		lints, err := lint.InspectFileErrUsage(path)
		if err != nil {
			panic(err)
		}
		for _, lint := range lints {
			fmt.Fprintln(os.Stderr, lint)
		}
	}
}
