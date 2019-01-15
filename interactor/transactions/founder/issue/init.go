package issue

import (
	"bytes"
	"github.com/ellsol/solidity-tools/utils"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/urfave/cli"
)

type CommandArgs struct {
	FounderId int64
}

func GetArgs(c *cli.Context) (*CommandArgs) {
	return &CommandArgs{
		FounderId: c.Int64("founderId"),
	}
}

func (it *CommandArgs) String() string {
	bb := bytes.Buffer{}

	bb.WriteString(utils.ConsoleWriteLabeledValueI("FounderId", it.FounderId))

	return bb.String()
}

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "issue founder token",
		ShortName: "i",
		Name:      "issue",
		Flags: utils.NewFlags().Add(cli.Int64Flag{
			Name:  "founderId",
			Usage: "founderId - either 0 or 1",
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
