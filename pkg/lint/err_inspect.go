package lint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// InspectFileErrUsage runs InspectErrUsage on a file on the filesystem.
func InspectFileErrUsage(
	path string) (lints []ErrInsufficientErrHandling, err error) {
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
	lints = InspectErrUsage(f, fSet)
	return
}

// InspectErrUsage returns the positions of all errors which are insufficiently
// used.
func InspectErrUsage(
	f *ast.File, fSet *token.FileSet) (errs []ErrInsufficientErrHandling) {
	x := 0
	ast.Inspect(f, func(n ast.Node) bool {
		if n != nil {
			x++
			fmt.Printf("%v: %T", x, n)
			switch n := n.(type) {
			case *ast.AssignStmt:
				fmt.Printf(" %s", n)
			case *ast.BasicLit:
				fmt.Printf(" %v", n)
			}
			fmt.Println()
		}
		return true
	})
	return
}

// ErrInsufficientErrHandling represents a position in a file which binds an
// error which is not sufficiently handled.
type ErrInsufficientErrHandling struct {
	Pos token.Position
}

func (e ErrInsufficientErrHandling) Error() string {
	return fmt.Sprintf("%s: error is not well handled", e.Pos)
}
