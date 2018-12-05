package utils

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli"
)

const RopstenInfura = "https://ropsten.infura.io/qyXg12v7VcFCETVWfUaJ/"
const KovanInfura = "https://kovan.infura.io/qyXg12v7VcFCETVWfUaJ/"
const Ganache = "http://localhost:7545"
const LocalDevNode = "http://localhost:8545"

type CliFlagsBuilder struct {
	flags []cli.Flag
}

func NewFlags() *CliFlagsBuilder {
	return &CliFlagsBuilder{
		flags: []cli.Flag{
			cli.StringFlag{
				Name:  "httpPath",
				Value: LocalDevNode,
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
			cli.StringFlag{
				Name:  "version",
				Value: "",
				Usage: "version of smart contracts",
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

type CliBaasContext struct {
	HttpPath        string
	KeystoreUTCPath string
	PasswordFile    string
	Version         string
}

func GetBaasContext(c *cli.Context) (*CliBaasContext, error) {
	var err error
	version := c.String("version")

	if version == "" {
		version, err = ReadVersion()
		if err != nil {
			return nil,err
		}
	}

	return &CliBaasContext{
		c.String("httpPath"), c.String("keystoreUTCPath"), c.String("passwordFile"), c.String("version"),
	}, nil
}


func GetBaasContextIgnoreVersionFromFile(c *cli.Context) (*CliBaasContext, error) {
	version := c.String("version")

	if version == "" {
		return nil, fmt.Errorf("version not supplied")
	}

	return &CliBaasContext{
		c.String("httpPath"), c.String("keystoreUTCPath"), c.String("passwordFile"), c.String("version"),
	}, nil
}

func (con *CliBaasContext) String() string {
	bb := bytes.Buffer{}

	bb.WriteString(ConsoleWriteLabeledValue("HttpPath", con.HttpPath))
	bb.WriteString(ConsoleWriteLabeledValue("Wallet File", con.KeystoreUTCPath))
	bb.WriteString(ConsoleWriteLabeledValue("Password File", con.PasswordFile))
	bb.WriteString(ConsoleWriteLabeledValue("Version", con.Version))

	return bb.String()
}
