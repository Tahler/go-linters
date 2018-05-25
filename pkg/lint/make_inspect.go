package lint

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
)

// InspectMakeCalls prints the positions of all incorrect make calls to stderr.
func InspectMakeCalls(f *ast.File, fSet *token.FileSet) {
	ast.Inspect(f, func(n ast.Node) bool {
		fn, isFunction := n.(*ast.CallExpr)
		if isFunction {
			if ident, ok := fn.Fun.(*ast.Ident); ok {
				if ident.Name == "make" {
					if len(fn.Args) == 2 {
						firstArg := fn.Args[0]
						if _, ok := firstArg.(*ast.ArrayType); ok {
							fmt.Fprintf(
								os.Stderr,
								"%s: make call does not specify both len and cap\n",
								fSet.Position(fn.Pos()))
						}
					}
				}
			}
		}
		shouldCheckChildren := !isFunction
		return shouldCheckChildren
	})
}
