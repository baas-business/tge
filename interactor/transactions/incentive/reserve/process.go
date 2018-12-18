package reserve

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/go-ethertypes"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func process(dAppContext *web3.DAppContext, args *CommandArgs) error {
	fmt.Println("Reward Incentive")
	fmt.Println(args)
	txOps := bind.NewKeyedTransactor(dAppContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasIncentives(tge.IncentivesAddress(dAppContext), dAppContext.Client)

	if err != nil {
		return err
	}

	etherValue, err := ethertypes.NewEtherValue().FromFloat64String(fmt.Sprintf("%v", args.Amount))
	tx, err := contract.Reserve(txOps,
		common.HexToAddress(args.Target),
		etherValue.BigInt(),
		[32]byte{},
	)

	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("Reward Incentive", dAppContext.Client, tx)
}
