package cmd

import (
	"github.com/matroskin13/grizzly/gen"
	"github.com/urfave/cli"

	"os"
	"path/filepath"
	"go/parser"
	"go/token"
	"io/ioutil"
	"go/ast"
	"strings"
)

func GenerateCommand() cli.Command {
	return cli.Command {
		Name: "generate",
		Aliases: []string{"g"},
		Usage: "generate collection from file",
		Action: generateAction,
	}
}

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

func generateAction(c *cli.Context) (err error) {
	pwd, _ := os.Getwd()
	generateFileName := c.Args().First()
	generateFilePath := filepath.Join(pwd, generateFileName)

	file, err := ioutil.ReadFile(generateFilePath)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, generateFileName, string(file), parser.ParseComments)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	var config gen.GrizzlyConfig

	ast.Inspect(f, func (n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if x.Tok == token.TYPE && grizzlyComment(x.Doc) {
				itemConfig := gen.GrizzlyConfigCollection{Types: map[string]string{}}

				ast.Inspect(x, func (n ast.Node) bool {
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

	for _, collection := range config.Collections {
		if len(collection.Methods) == 0 {
			collection.Methods = gen.GetDefaultMethods()
		}

		code, err := gen.GenCollectionCode(collection)

		if err != nil {
			return cli.NewExitError(err, 0)
		}

		err = gen.CreateCollection(strings.ToLower(collection.Name) + "_collection", code, true, pwd)

		if err != nil {
			return cli.NewExitError(err, 0)
		}
	}

	return err
}
