package setup

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/transactions/deployer"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func process(httpPath string, keystoreUTCPath string, passwordFile string, version string) error {
	client, err := utils.RPCClient(httpPath)

	if err != nil {
		return err
	}

	privKey, err := utils.PrivateKeyFromWalletAndPasswordFile(keystoreUTCPath, passwordFile)

	if err != nil {
		return err
	}

	txOps := bind.NewKeyedTransactor(privKey.PrivateKey)
	txOps.Value = big.NewInt(0)

	cc, err := deployer.LoadContractConfig(fmt.Sprintf("build/%v/contract_config.json", version))

	if err != nil {
		return err
	}

	contract, err := contracts.NewBaasToken(*cc.TokenAddress, client)

	if err != nil {
		return err
	}

	ii, err := contract.IsInitialized(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("initialized: ", ii)

	//Setup(opts *bind.TransactOpts, escrowAddress common.Address, ppAddress common.Address, founderAddress common.Address, incentivesAddress common.Address) (*types.Transaction, error) {
	tx, err := contract.Setup(txOps, *cc.EscrowAddress, *cc.PPAddress, *cc.FounderAddress, *cc.IncentivesAddress)
	return utils.ExecuteTransaction("Setup Token", client, tx)
}
