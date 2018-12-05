package helper

import (
	"github.com/baas/tge-sol/interactor/transactions/helper/advance"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "helper",
		ShortName: "h",
		Name:      "helper",
		Subcommands: []cli.Command{
			*advance.Command(),
		},
	}
}
