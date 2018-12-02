package roi

import (
	"github.com/baas/tge-sol/interactor/transactions/roi/setup"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "interacts with roi contract",
		ShortName: "r",
		Name:      "roi",
		Subcommands: []cli.Command{
			*setup.Command(),
		},
	}
}
