package gen

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func grizzlyComment(doc *ast.CommentGroup) bool {
	if doc == nil {
		return false
	}

	for _, comment := range doc.List {
		if comment.Text == "//grizzly:generate" {
			return true
		}
	}

	return false
}

type GrizzlyConfigCollection struct {
	Name    string            `json:"name"`
	Types   map[string]string `json:"types"`
	Methods []string
	Package string
}

type GrizzlyConfig struct {
	Collections []GrizzlyConfigCollection `json:"collections"`
}

// Create config from GO code
func GetConfigByCode(code []byte) (*GrizzlyConfig, error) {
	var config GrizzlyConfig

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "main.go", string(code), parser.ParseComments)

	if err != nil {
		return nil, err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		x, ok := n.(*ast.GenDecl)

		if ok && x.Tok == token.TYPE && x.Doc != nil {
			_, isExist := GetGrizzlyCommand(x.Doc)[CommandGenerate]

			if isExist {
				itemConfig := GrizzlyConfigCollection{
					Types:   map[string]string{},
					Package: f.Name.Name,
					Methods: GetDefaultMethods(),
				}

				ast.Inspect(x, func(n ast.Node) bool {
					switch x := n.(type) {
					case *ast.Field:
						switch y := x.Type.(type) {
						case *ast.Ident:
							itemConfig.Types[x.Names[0].Name] = y.Name
						}
					case *ast.Ident:
						if itemConfig.Name == "" {
							itemConfig.Name = x.Name
						}
					}

					return true
				})

				config.Collections = append(config.Collections, itemConfig)
			}
		}

		return true
	})

	return &config, nil
}
