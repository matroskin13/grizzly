package cmd

import (
	"os"

	"github.com/urfave/cli"
)

func Init() {
	app := cli.NewApp()

	app.Name = "grizzly"
	app.Usage = "codegen for golang"
	app.UsageText = "grizzly command"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		CreateCommand(),
		UpdateCommand(),
		GenerateCommand(),
	}

	app.Run(os.Args)
}
