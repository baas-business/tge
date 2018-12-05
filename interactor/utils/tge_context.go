package utils

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contractconf"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
)

type TGEContext struct {
	Client         *ethclient.Client
	Key            *keystore.Key
	ContractConfig *config.ContractConfig
	Version        string
}

func GetTGEContext(c *cli.Context) (*TGEContext, error) {
	con := &TGEContext{}

	cliCon, err := GetBaasContext(c)

	if err != nil {
		return nil, err
	}
	fmt.Println(cliCon.String())

	con.Client, err = RPCClient(cliCon.HttpPath)

	if err != nil {
		return nil, err
	}

	con.Key, err = PrivateKeyFromWalletAndPasswordFile(cliCon.KeystoreUTCPath, cliCon.PasswordFile)

	if err != nil {
		return nil, err
	}
	con.ContractConfig, err = config.LoadContractConfig(fmt.Sprintf("build/%v/contract_config.json", cliCon.Version))

	if err != nil {
		return nil, err
	}

	con.Version = c.String("version")

	return con, nil
}

func GetDeployContext(c *cli.Context) (*TGEContext, error) {
	con := &TGEContext{}
	var err error
	cliCon, err := GetBaasContextIgnoreVersionFromFile(c)

	if err != nil {
		return nil, err
	}
	fmt.Println(cliCon.String())


	con.Client, err = RPCClient(cliCon.HttpPath)

	if err != nil {
		return nil, err
	}

	con.Key, err = PrivateKeyFromWalletAndPasswordFile(cliCon.KeystoreUTCPath, cliCon.PasswordFile)

	if err != nil {
		return nil, err
	}

	con.Version = c.String("version")

	return con, nil
}

func (it *TGEContext) TokenAddress() *common.Address {
	return it.ContractConfig.TokenAddress
}

func (it *TGEContext) EscrowAddress() *common.Address {
	return it.ContractConfig.EscrowAddress
}

func (it *TGEContext) PPAddress() *common.Address {
	return it.ContractConfig.PPAddress
}

func (it *TGEContext) FounderAddress() *common.Address {
	return it.ContractConfig.FounderAddress
}

func (it *TGEContext) IncentivesAddress() *common.Address {
	return it.ContractConfig.IncentivesAddress
}

func (it *TGEContext) ROIAddress() *common.Address {
	return it.ContractConfig.ROIAddress
}