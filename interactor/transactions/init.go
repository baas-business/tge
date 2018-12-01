package transactions

import (
	"github.com/baas/tge-sol/interactor/transactions/pp"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "execute transactions on contract",
		ShortName: "e",
		Name:      "execute",
		Subcommands: []cli.Command{
			*pp.Command(),
		},
	}
}
