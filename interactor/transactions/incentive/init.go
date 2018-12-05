package incentive

import (
	"github.com/baas/tge-sol/interactor/transactions/incentive/reward"
	"github.com/baas/tge-sol/interactor/transactions/incentive/setup"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "interacts with roi contract",
		ShortName: "i",
		Name:      "incentive",
		Subcommands: []cli.Command{
			*setup.Command(),
			*reward.Command(),
		},
	}
}
