package claim

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(tgeContext *utils.TGEContext, args *CommandArgs) error {
	fmt.Println("Setup Incentive")
	fmt.Println(args.String())
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasIncentives(*tgeContext.ContractConfig.IncentivesAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Claim(txOps)


	if err != nil {
		return err
	}

	return utils.ExecuteTransaction("Setup Incentive", tgeContext.Client, tx)
}
