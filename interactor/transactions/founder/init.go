package founder

import (
	"github.com/baas-business/tge-sol/interactor/transactions/founder/issue"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "founder transaction",
		ShortName: "f",
		Name:      "founder",
		Subcommands: []cli.Command{
			*issue.Command(),
		},
	}
}
