package main

import (
	"github.com/baas/tge-sol/interactor/info"
	"github.com/baas/tge-sol/interactor/transactions"
	"github.com/baas/tge-sol/interactor/transactions/deployer"
	"github.com/baas/tge-sol/interactor/transactions/setup"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "baas tge"
	app.Usage = "app to deploy tge contracts and execute transactions"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		*deployer.Command(),
		*setup.Command(),
		*info.Command(),
		*transactions.Command(),
	}

	app.Run(os.Args)
}
