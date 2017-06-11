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
		Action: createAction,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name: "dev",
				Usage: "set development mode",
			},
		},
	}
}

func createAction(c *cli.Context) (err error) {
	types := map[string]string{}
	modelName := strings.ToLower(c.Args().First());

	for _, blob := range c.Args().Tail() {
		blobTypes := strings.Split(blob, ":")

		types[blobTypes[0]] = blobTypes[1]
	}

	if modelName == "" {
		return cli.NewExitError("model name is empty", 0)
	}

	code, err := gen.GetCollectionCode(c.Bool("dev"), modelName, types)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	err = gen.CreateCollection(modelName, code, false)

	if err != nil {
		return cli.NewExitError(err, 0)
	}

	return err
}
