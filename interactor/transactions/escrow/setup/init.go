package setup

import (
	"bytes"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/urfave/cli"
)


type CommandArgs struct {
	Vesting int64
}

func GetArgs(c *cli.Context) (*CommandArgs) {
	return &CommandArgs{
		Vesting: c.Int64("vesting"),
	}
}

func (it *CommandArgs) String() string {
	bb := bytes.Buffer{}

	bb.WriteString(utils.ConsoleWriteLabeledValueI("Vesting", it.Vesting))

	return bb.String()
}

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "setup escrow smart contract",
		ShortName: "s",
		Name:      "setup",
		Flags: utils.NewFlags().Add(cli.Int64Flag{
			Name:  "vesting",
			Usage: "vesting period for founder",
		}).Get(),
		Action: func (c *cli.Context) error{
		tgeContext, err := utils.GetTGEContext(c)

		if err != nil{
		return err
	}

		return process(tgeContext, GetArgs(c))
	},
	}
}
