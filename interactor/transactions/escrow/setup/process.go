package setup

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(tgeContext *utils.TGEContext, args *CommandArgs) error {
	fmt.Println("Setup Escrow")
	fmt.Println(args)
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasEscrow(*tgeContext.ContractConfig.EscrowAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Setup(txOps, new(big.Int).SetInt64(args.Vesting))

	if err != nil {
		return err
	}

	return utils.ExecuteTransaction("Setup Escrow", tgeContext.Client, tx)
}
