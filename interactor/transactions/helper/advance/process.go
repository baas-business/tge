package advance

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(dappContext *web3.DAppContext) error {
	fmt.Println("Setup Incentive")
	txOps := bind.NewKeyedTransactor(dappContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasToken(tge.TokenAddress(dappContext), dappContext.Client)

	if err != nil {
		return err
	}

	isPaused, err := contract.Paused(&bind.CallOpts{})
	if err != nil {
		return err
	}

	tx, err := contract.SetPause(txOps, !isPaused)


	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("Setup Incentive", dappContext.Client, tx)
}
