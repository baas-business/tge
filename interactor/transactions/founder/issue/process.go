package issue

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func process(dappContext *web3.DAppContext, args *CommandArgs) error {
	fmt.Println("Setup Founder")
	fmt.Println(args)
	txOps := bind.NewKeyedTransactor(dappContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasFounder(tge.FounderAddress(dappContext), dappContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Issue(txOps, common.HexToAddress("0xb11700D501A9b5D3960749E0B7314A21743Dab5d"), new(big.Int).SetInt64(args.FounderId))

	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("Setup Founder", dappContext.Client, tx)
}
