package advance

import (
	"github.com/ellsol/solidity-tools/utils"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/urfave/cli"
)


func Command() *cli.Command {
	return &cli.Command{
		Usage:     "setup incentive smart contract",
		ShortName: "a",
		Name:      "advance",
		Flags: utils.NewFlags().Get(),
		Action: func(c *cli.Context) error {
			context, err := web3.GetDAppContext()

			if err != nil {
				return err
			}

			return process(context)
		},
	}
}
