package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/Tahler/go-linters/pkg/lint"
)

func main() {
	src := `package foo

import (
	"fmt"
)

var y = make([]string, 2)
var t = make([]string, 0, 3)
var z = make(map[string]string, 2)

func bar() {
	x := make([]string, 5)
	fmt.Println(x)
}

`

	// Create the AST by parsing src.
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "src.go", src, parser.Mode(0))
	if err != nil {
		panic(err)
	}

	errs := lint.InspectMakeCalls(f, fSet)
	for _, err := range errs {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
