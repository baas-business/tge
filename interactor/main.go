package main

import (
	"github.com/baas-business/tge-sol/interactor/info"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/baas-business/tge-sol/interactor/transactions"
	"github.com/ellsol/solidity-tools/commands"
	"github.com/ellsol/solidity-tools/commands/deployer"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "baas tge"
	app.Usage = "app to deploy tge contracts and execute transactions"
	app.Version = "0.0.5"

	deployer.SetDAppCallback(tge.Callback)

	app.Commands = append(commands.Commands(),
		*info.Command(),
		*transactions.Command(),
	)

	app.Run(os.Args)
}
