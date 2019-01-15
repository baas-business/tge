package incentive

import (
	"github.com/baas-business/tge-sol/interactor/transactions/incentive/issue"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "interacts with roi contract",
		ShortName: "i",
		Name:      "incentive",
		Subcommands: []cli.Command{
			*issue.Command(),
		},
	}
}
