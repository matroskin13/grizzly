package cmd

import (
	"strings"
	"github.com/urfave/cli"

	"github.com/matroskin13/grizzly/gen"
)

func UpdateCommand() cli.Command {
	return cli.Command {
		Name: "update",
		Aliases: []string{"u"},
		Usage: "update collections by config",
		Action: updateAction,
	}
}

func updateAction(c *cli.Context) (err error) {
	config, err := gen.GetConfig()

	if err != nil || config == nil {
		return cli.NewExitError("config is not readable", 0)
	}

	for _, collection := range config.Collections {
		if collection.Name == "" {
			return cli.NewExitError("collection name is empty", 0)
		}

		if len(collection.Methods) == 0 {
			collection.Methods = gen.GetDefaultMethods()
		}

		code, err := gen.GenCollectionCode(collection)

		if err != nil {
			return cli.NewExitError(err, 0)
		}

		err = gen.CreateCollection(strings.ToLower(collection.Name), code, true)

		if err != nil {
			return cli.NewExitError(err, 0)
		}
	}

	return err
}
