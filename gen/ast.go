package gen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/format"
	"text/template"
	"bytes"
	"strings"
)

// Return grizzly commands by comments of code. Example:
//
// //grizzly:replaceName GetGrizzly{{.Name}}
// func GetGrizzly() {}
func GetGrizzlyCommand(doc *ast.CommentGroup) map[string]GrizzlyCommand {
	commands := make(map[string]GrizzlyCommand)

	for _, comment := range doc.List {
		if strings.Contains(comment.Text, "//grizzly:") {
			var command GrizzlyCommand

			arr := strings.Split(comment.Text, " ")

			command.Command = arr[0][2:]
			command.Action = arr[1]

			commands[command.Command] = command
		}
	}

	return commands
}

func GenCode(config *GrizzlyConfigCollection, code []byte) []byte {
	fset := token.NewFileSet()

	f, _ := parser.ParseFile(fset, "main.go", code, parser.ParseComments)

	ApplyCommands(f, config)
	SwapTypes(f, config)
	InjectTypes(f, config)

	var buf bytes.Buffer

	format.Node(&buf, fset, f)

	return buf.Bytes()
}

func ApplyCommands(node *ast.File, config *GrizzlyConfigCollection) *ast.File {
	ast.Inspect(node, func (n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			commands := make(map[string]GrizzlyCommand)

			if x.Doc != nil {
				commands = GetGrizzlyCommand(x.Doc)
			}

			if command, ok := commands[CommandReplaceName]; ok {
				var buf bytes.Buffer

				tmp, _ := template.New(CommandReplaceName).Parse(command.Action)

				tmp.Execute(&buf, config)

				x.Name.Name = buf.String()
			}
		}

		return true
	})

	return node
}

// Replaces grizzly Model and Collection
func SwapTypes(node *ast.File, config *GrizzlyConfigCollection) {
	ast.Inspect(node, func (n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			if x.Name == GrizzlyCollection {
				x.Name = config.Name + "Collection"
			}

			if x.Name == GrizzlyModel {
				x.Name = config.Name
			}
		}

		return true
	})
}

func InjectTypes(node *ast.File, config *GrizzlyConfigCollection) {
	ast.Inspect(node, func (n ast.Node) bool {
		if x, ok := n.(*ast.GenDecl); ok && x.Tok == token.TYPE {
			if tSpec, ok := x.Specs[0].(*ast.TypeSpec); ok {
				if sType, ok := tSpec.Type.(*ast.StructType); tSpec.Name.Name == config.Name && ok {
					sType.Fields.List = []*ast.Field{}

					for key, customType := range config.Types {
						sType.Fields.List = append(sType.Fields.List, &ast.Field{
							Names: []*ast.Ident{ast.NewIdent(key)},
							Type: ast.NewIdent(customType),
						})
					}
				}
			}
		}

		return true
	})
}

