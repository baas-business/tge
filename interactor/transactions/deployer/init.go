package deployer

import (
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/urfave/cli"
)


func Command() *cli.Command {
	return &cli.Command{
		Usage:     "deploys 6 tge contracts",
		ShortName: "d",
		Name:      "deploy",
		Flags: utils.NewFlags().Add(
			cli.StringFlag{
				Name:  "version",
				Value: "v0",
				Usage: "version of smart contracts",
			}).Get(),
		Action: func(c *cli.Context) error {
			return process(c.String("httpPath"), c.String("passwordFile"), c.String("keystoreUTCPath"), c.String("version"))
		},
	}
}
