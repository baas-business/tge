package claim

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(dAppContext *web3.DAppContext, args *CommandArgs) error {
	fmt.Println("Setup Incentive")
	fmt.Println(args.String())
	txOps := bind.NewKeyedTransactor(dAppContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasIncentives(tge.IncentivesAddress(dAppContext), dAppContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Claim(txOps)


	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("Setup Incentives", dAppContext.Client, tx)
}
