package setup

import (
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "setups baas token with 4 pot address loaded from config file",
		ShortName: "s",
		Name:      "setup",
		Action: func(c *cli.Context) error {
			context, err := web3.GetDAppContext()

			if err != nil {
				return err
			}

			return process(context)
		},
	}
}
