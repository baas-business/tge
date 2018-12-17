package payoutall

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(dAppContext *web3.DAppContext) error {
	fmt.Println("Payout all")
	txOps := bind.NewKeyedTransactor(dAppContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasROI(tge.ROIAddress(dAppContext), dAppContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.PayoutAll(txOps, new(big.Int).SetInt64(100))


	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("PayoutAll ROI", dAppContext.Client, tx)
}
