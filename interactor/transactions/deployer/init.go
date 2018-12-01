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
		Flags: utils.NewFlags().Get(),
		Action: func(c *cli.Context) error {

			tgeContext, err := utils.GetTGEContext(c)

			if err != nil {
				return err
			}
			return process(tgeContext)
		},
	}
}