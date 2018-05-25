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
	fmt.Println(make([]string, 5))

	y := struct {
		l []string
	}{l: make([]string, 2)}
	fmt.Println(y)
}
`

	// Create the AST by parsing src.
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, "src.go", src, parser.Mode(0))
	if err != nil {
		panic(err)
	}

	expected := []string{
		"src.go:5:9",
		"src.go:10:7",
		"src.go:12:14",
		"src.go:16:7",
	}

	actual := make([]string, 0, len(expected))
	errs := InspectMakeCalls(f, fSet)
	for _, err := range errs {
		actual = append(actual, err.Pos.String())
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v; actual %v", expected, actual)
	}
}
