package transactions

import (
	"github.com/baas-business/tge-sol/interactor/transactions/escrow"
	"github.com/baas-business/tge-sol/interactor/transactions/founder"
	"github.com/baas-business/tge-sol/interactor/transactions/helper"
	"github.com/baas-business/tge-sol/interactor/transactions/incentive"
	"github.com/baas-business/tge-sol/interactor/transactions/pp"
	"github.com/baas-business/tge-sol/interactor/transactions/roi"
	"github.com/baas-business/tge-sol/interactor/transactions/token"
	"github.com/urfave/cli"
)

func Command() *cli.Command {
	return &cli.Command{
		Usage:     "execute transactions on contract",
		ShortName: "e",
		Name:      "execute",
		Subcommands: []cli.Command{
			*pp.Command(),
			*token.Command(),
			*roi.Command(),
			*incentive.Command(),
			*helper.Command(),
			*founder.Command(),
			*escrow.Command(),
		},
	}
}
