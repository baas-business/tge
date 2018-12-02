package setup

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(tgeContext *utils.TGEContext) error {
	fmt.Println("Setting up ROI Baas token")
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	_, err := contracts.NewBaasROI(*tgeContext.ContractConfig.PPAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	//tx, err := contract.Setup(txOps)
	//fmt.Printf("%+v",tx)
	//
	//if err != nil {
	//	return err
	//}

	return nil//utils.ExecuteTransaction("Setup ROI", tgeContext.Client, tx)
}
