package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		fn, isFunction := n.(*ast.CallExpr)
		if isFunction {
			if ident, ok := fn.Fun.(*ast.Ident); ok {
				if ident.Name == "make" {
					if len(fn.Args) == 2 {
						firstArg := fn.Args[0]
						if _, ok := firstArg.(*ast.ArrayType); ok {
							fmt.Printf("should have 2 args\n")
						}
					}
				}
			}
		}
		shouldCheckChildren := !isFunction
		return shouldCheckChildren
	})
}
