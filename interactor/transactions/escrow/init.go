package escrow

import (
	"github.com/baas/tge-sol/interactor/transactions/escrow/setup"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "founder transaction",
		ShortName: "e",
		Name:      "escrow",
		Subcommands: []cli.Command{
			*setup.Command(),
		},
	}
}
