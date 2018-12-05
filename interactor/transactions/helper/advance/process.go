package advance

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(tgeContext *utils.TGEContext) error {
	fmt.Println("Setup Incentive")
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasToken(*tgeContext.ContractConfig.TokenAddress, tgeContext.Client)

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

	return utils.ExecuteTransaction("Setup Incentive", tgeContext.Client, tx)
}
