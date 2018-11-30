package contractinfo

import (
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "setups baas token with 4 pot address loaded from config file",
		ShortName: "ci",
		Name:      "contractinfo",
		Flags: utils.NewFlags().Add(
			cli.StringFlag{
				Name:  "version",
				Value: "v0",
				Usage: "version of smart contracts",
			}).Get(),
		Action: func(c *cli.Context) error {
			return process(c.String("httpPath"), c.String("version"))
		},
	}
}
