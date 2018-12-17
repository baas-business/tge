package claim

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
	}
}

func (it *CommandArgs) String() string {
	bb := bytes.Buffer{}

	bb.WriteString(utils.ConsoleWriteLabeledValue("Target", it.Target))

	return bb.String()
}

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "claim incentive smart contract",
		ShortName: "c",
		Name:      "claim",
		Flags: utils.NewFlags().Add(cli.StringFlag{
			Name:  "target",
			Usage: "target address of investor",
		}).Get(),
		Action: func(c *cli.Context) error {
			context, err := web3.GetDAppContext()

			if err != nil {
				return err
			}

			return process(context, GetArgs(c))
		},
	}
}
