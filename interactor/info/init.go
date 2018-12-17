package info

import (
	"github.com/baas-business/tge-sol/interactor/info/contractinfo"
	"github.com/baas-business/tge-sol/interactor/info/pp"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "show read values of contract",
		ShortName: "i",
		Name:      "info",
		Subcommands: []cli.Command{
			*contractinfo.Command(),
			*pp.Command(),
		},
	}
}
