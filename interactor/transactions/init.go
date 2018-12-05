package transactions

import (
	"github.com/baas/tge-sol/interactor/transactions/deployer"
	"github.com/baas/tge-sol/interactor/transactions/incentive"
	"github.com/baas/tge-sol/interactor/transactions/pp"
	"github.com/baas/tge-sol/interactor/transactions/roi"
	"github.com/baas/tge-sol/interactor/transactions/setup"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "execute transactions on contract",
		ShortName: "e",
		Name:      "execute",
		Subcommands: []cli.Command{
			*pp.Command(),
			*deployer.Command(),
			*setup.Command(),
			*roi.Command(),
			*incentive.Command(),
		},
	}
}
