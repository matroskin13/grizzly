package cmd

import (
	"strings"

	"github.com/urfave/cli"

	"grizzly/gen"
)

func CreateCommand() cli.Command {
	return cli.Command {
		Name: "create",
		Aliases: []string{"c"},
		Usage: "create model and collection by name",
		Action: createCommand,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "dev",
				Usage: "set development mode",
			},
		},
	}
}

func createCommand(c *cli.Context) (err error) {
	modelName := strings.ToLower(c.Args().First());

	if modelName == "" {
		return cli.NewExitError("model name is empty", 0)
	}

	code, err := gen.GetCollectionCode(c.Bool("dev"), modelName)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	err = gen.CreateCollection(modelName, code)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	return err
}
