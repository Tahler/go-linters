package lint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// InspectFileMakeCalls runs InspectMakeCalls on a file on the filesystem.
func InspectFileMakeCalls(
	path string) (lints []ErrMakeSliceWithoutCap, err error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return
	}
	if fileInfo.IsDir() {
		err = fmt.Errorf("%s is a directory", path)
		return
	}

	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, path, nil, parser.Mode(0))
	if err != nil {
		return
	}
	lints = InspectMakeCalls(f, fSet)
	return
}

// InspectMakeCalls returns the positions of all incorrect make calls.
func InspectMakeCalls(
	f *ast.File, fSet *token.FileSet) (errs []ErrMakeSliceWithoutCap) {
	ast.Inspect(f, func(n ast.Node) bool {
		fn, isFunction := n.(*ast.CallExpr)
		if isFunction {
			if ident, ok := fn.Fun.(*ast.Ident); ok {
				if ident.Name == "make" {
					if len(fn.Args) == 2 {
						firstArg := fn.Args[0]
						if _, ok := firstArg.(*ast.ArrayType); ok {
							pos := fSet.Position(fn.Pos())
							errs = append(errs, ErrMakeSliceWithoutCap{pos})
						}
					}
				}
			}
		}
		return true
	})
	return
}

// ErrMakeSliceWithoutCap represents a position in a file which calls make to
// construct a slice with a length but not a capacity.
type ErrMakeSliceWithoutCap struct {
	Pos token.Position
}

func (e ErrMakeSliceWithoutCap) Error() string {
	return fmt.Sprintf("%s: make call does not specify both len and cap", e.Pos)
}
