package token

import (
	"github.com/baas-business/tge-sol/interactor/transactions/token/setup"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "interacts with token contract",
		ShortName: "t",
		Name:      "token",
		Subcommands: []cli.Command{
			*setup.Command(),
		},
	}
}
