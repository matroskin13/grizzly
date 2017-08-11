package cmd

import (
	"github.com/matroskin13/grizzly/gen"
	"github.com/urfave/cli"

	"os"
	"path/filepath"
	"io/ioutil"
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

func generateAction(c *cli.Context) (err error) {
	pwd, _ := os.Getwd()
	generateFileName := c.Args().First()
	generateFilePath := filepath.Join(pwd, generateFileName)

	file, err := ioutil.ReadFile(generateFilePath)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	config, err := gen.GetConfigByCode(file)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	for _, collection := range config.Collections {
		if len(collection.Methods) == 0 {
			collection.Methods = gen.GetDefaultMethods()
		}

		code, err := gen.GenCollectionCode(collection, true)

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
