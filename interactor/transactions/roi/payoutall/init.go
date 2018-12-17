package payoutall

import (
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/urfave/cli"
)


func Command() *cli.Command {
	return &cli.Command{
		Usage:     "payout all",
		ShortName: "p",
		Name:      "payout",
		Action: func(c *cli.Context) error {
			tgeContext, err := web3.GetDAppContext()

			if err != nil {
				return err
			}

			return process(tgeContext)
		},
	}
}
