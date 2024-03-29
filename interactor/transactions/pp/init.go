package pp

import (
	"fmt"
	"github.com/ellsol/solidity-tools/utils"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "interacts with private placement contract",
		ShortName: "pp",
		Name:      "pp",
		Subcommands: []cli.Command{
			*provideTokenCommand(),
		},
	}
}

func provideTokenCommand() *cli.Command {
	return &cli.Command{
		Usage:     "provide tokens",
		ShortName: "p",
		Name:      "provide",
		Flags: utils.NewFlags().
			Add(
				cli.BoolFlag{
					Name:  "discounted",
					Usage: "are tokens discounted?",
				}).
			Add(
				cli.StringFlag{
					Name:  "target",
					Usage: "target address of investor",
				}).
			Add(
				cli.Float64Flag{
					Name:  "amount",
					Usage: "token amount to be provided",
				}).
			Get(),
		Action: func(c *cli.Context) error {
			isDiscounted := c.Bool("discounted")
			target := c.String("target")
			amount := c.Float64("amount")

			if target == "" {
				return fmt.Errorf("target cannot be empty")
			}

			context, err := web3.GetDAppContext()

			if err != nil {
				return err
			}

			fmt.Println("providing token discounted:", isDiscounted, "target: ", target, "amount: ", amount)

			return process(context, isDiscounted, target, amount)
		},
	}
}
