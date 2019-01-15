package issue

import (
	"bytes"
	"github.com/ellsol/solidity-tools/utils"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/urfave/cli"
)

type CommandArgs struct {
	Target string
	Amount float64
	Stages int64
	Blocks int64
}

func GetArgs(c *cli.Context) (*CommandArgs) {
	return &CommandArgs{
		Target: c.String("target"),
		Amount: c.Float64("amount"),
	}
}

func (it *CommandArgs) String() string {
	bb := bytes.Buffer{}

	bb.WriteString(utils.ConsoleWriteLabeledValue("Target", it.Target))
	bb.WriteString(utils.ConsoleWriteLabeledValueI("Amount", it.Amount))

	return bb.String()
}

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "setup incentive smart contract",
		ShortName: "r",
		Name:      "reward",
		Flags: utils.NewFlags().
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
			context, err := web3.GetDAppContext()

			if err != nil {
				return err
			}


			args := GetArgs(c)

			if err != nil {
				return err
			}

			return process(context, args)
		},
	}
}
