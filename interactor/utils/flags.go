package utils

import "github.com/urfave/cli"


const Infura = "https://ropsten.infura.io/3l5dxBOP3wPspnRDdG1u"

type CliFlagsBuilder struct {
	flags []cli.Flag
}

func NewFlags() *CliFlagsBuilder {
	return &CliFlagsBuilder{
		flags: []cli.Flag{
			cli.StringFlag{
				Name:  "httpPath",
				Value: "https://ropsten.infura.io/qyXg12v7VcFCETVWfUaJ/",
				Usage: "HTTP-RPC server listening interface",
			},
			cli.StringFlag{
				Name:  "passwordFile",
				Value: "./wallets/password.txt",
				Usage: "Password file for unlock account",
			},
			cli.StringFlag{
				Name:  "keystoreUTCPath",
				Value: "./wallets/wallet_android.json",
				Usage: "Location of keystore",
			},
		},
	}
}

func (it *CliFlagsBuilder) Add(flag cli.Flag) *CliFlagsBuilder {
	it.flags = append(it.flags, flag)
	return it
}

func (it *CliFlagsBuilder) Get() []cli.Flag {
	return it.flags
}