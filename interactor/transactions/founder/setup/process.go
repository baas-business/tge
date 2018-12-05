package setup

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(tgeContext *utils.TGEContext, args *CommandArgs) error {
	fmt.Println("Setup Founder")
	fmt.Println(args)
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasFounder(*tgeContext.ContractConfig.FounderAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Setup(txOps, new(big.Int).SetInt64(args.Vesting))

	if err != nil {
		return err
	}

	return utils.ExecuteTransaction("Setup Founder", tgeContext.Client, tx)
}
