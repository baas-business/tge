package main

import (
	"github.com/baas/tge-sol/interactor/info/contractinfo"
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
		*contractinfo.Command(),
	}

	app.Run(os.Args)
}
