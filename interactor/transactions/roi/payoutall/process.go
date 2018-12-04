package payoutall

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(tgeContext *utils.TGEContext) error {
	fmt.Println("Payout all")
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)
	//txOps.GasPrice = new(big.Int).SetInt64(500000000 *10)
	//txOps.GasLimit = 5000000

	contract, err := contracts.NewBaasROI(*tgeContext.ContractConfig.ROIAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.PayoutAll(txOps, new(big.Int).SetInt64(100))


	if err != nil {
		return err
	}

	return utils.ExecuteTransaction("PayoutAll ROI", tgeContext.Client, tx)
}
