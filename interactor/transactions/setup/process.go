package setup

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func process(tgeContext *utils.TGEContext) error {
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasToken(*tgeContext.ContractConfig.TokenAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	ii, err := contract.IsInitialized(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("initialized: ", ii)

	//Setup(opts *bind.TransactOpts, escrowAddress common.Address, ppAddress common.Address, founderAddress common.Address, incentivesAddress common.Address) (*types.Transaction, error) {
	tx, err := contract.Setup(txOps, *tgeContext.EscrowAddress(), *tgeContext.PPAddress(), *tgeContext.FounderAddress(), *tgeContext.IncentivesAddress())
	return utils.ExecuteTransaction("Setup Token", tgeContext.Client, tx)
}
