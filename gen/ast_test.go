package gen

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

var code = `package test

//grizzly:replaceName Get{{.Name}}
func GetChicken() {}

//grizzly:replaceName GetGet{{.Name}}
func GetPig() {}
`

var expectedCode = `package test

//grizzly:replaceName Get{{.Name}}
func GetTest() {}

//grizzly:replaceName GetGet{{.Name}}
func GetGetTest() {}
`

func HelperGetAst(code []byte) (*ast.File, func() string) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "test.go", code, parser.ParseComments)

	return f, func() string {
		var buf bytes.Buffer

		format.Node(&buf, fset, f)

		return buf.String()
	}
}

func TestGetGrizzlyCommand(t *testing.T) {
	testAst, _ := HelperGetAst([]byte(code))
	fun, _ := testAst.Decls[0].(*ast.FuncDecl)

	commands := GetGrizzlyCommand(fun.Doc)

	expected := map[string]GrizzlyCommand{
		"grizzly:replaceName": GrizzlyCommand{
			Command: "grizzly:replaceName",
			Action:  "Get{{.Name}}",
		},
	}

	assert.Equal(t, expected, commands)
}

func TestApplyCommands(t *testing.T) {
	config := &GrizzlyConfigCollection{Name: "Test"}
	testAst, finish := HelperGetAst([]byte(code))

	ApplyCommands(testAst, config)

	result := finish()

	assert.Equal(t, expectedCode, result)
}

var code1 = `package test

type Model struct{}
type Collection struct{}

func GetModel(arg *Model) *Model {
	return &Model{}
}

func GetCollection(arg *TestCollection) *Collection {
	return &Collection{}
}
`

var expectedCode1 = `package test

type Test struct{}
type TestCollection struct{}

func GetModel(arg *Test) *Test {
	return &Test{}
}

func GetCollection(arg *TestCollection) *TestCollection {
	return &TestCollection{}
}
`

func TestSwapTypes(t *testing.T) {
	testAst, finish := HelperGetAst([]byte(code1))
	config := &GrizzlyConfigCollection{Name: "Test"}

	SwapTypes(testAst, config)

	result := finish()

	assert.Equal(t, expectedCode1, result)
}

var code2 = `package test

type Test struct {
	One string
}
`

var expectedCode2 = `package test

type Test struct {
	Id int
}
`

func TestInjectTypes(t *testing.T) {
	testAst, finish := HelperGetAst([]byte(code2))
	config := &GrizzlyConfigCollection{
		Name: "Test",
		Types: map[string]string{
			"Id": "int",
		},
	}

	InjectTypes(testAst, config)

	result := finish()

	assert.Equal(t, expectedCode2, result)
}
