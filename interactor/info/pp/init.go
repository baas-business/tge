package pp

import (
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "setups baas token with 4 pot address loaded from config file",
		ShortName: "pp",
		Name:      "pp",
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
