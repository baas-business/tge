package setup

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(dappContext *web3.DAppContext, args *CommandArgs) error {
	fmt.Println("Setup Founder")
	fmt.Println(args)
	txOps := bind.NewKeyedTransactor(dappContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	_, err := contracts.NewBaasFounder(tge.FounderAddress(dappContext), dappContext.Client)

	if err != nil {
		return err
	}

	//tx, err := contract.Setup(txOps, new(big.Int).SetInt64(args.Vesting))
	//
	//if err != nil {
	//	return err
	//}

	return nil //web3.ExecuteTransaction("Setup Founder", dappContext.Client, tx)
}
