package lint

import (
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestInspectMakeCalls(t *testing.T) {
	src := `package foo

import "fmt"

var _ = make([]string, 2)
var _ = make([]string, 0, 3)
var _ = make(map[string]string, 2)

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

	expected := []ErrMakeSliceWithoutCap{
		ErrMakeSliceWithoutCap{fSet.Position(36)},
		ErrMakeSliceWithoutCap{fSet.Position(138)},
	}

	actual := InspectMakeCalls(f, fSet)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v; actual %v", expected, actual)
	}
}
