package reward

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ellsol/go-ethertypes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func process(tgeContext *utils.TGEContext, args *CommandArgs) error {
	fmt.Println("Reward Incentive")
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasIncentives(*tgeContext.ContractConfig.IncentivesAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	etherValue, err := ethertypes.NewEtherValue().FromFloat64String(fmt.Sprintf("%v", args.Amount))
	tx, err := contract.Reward(txOps,
		common.HexToAddress(args.Target),
		etherValue.BigInt(),
		new(big.Int).SetInt64(args.Stages),
		new(big.Int).SetInt64(args.Blocks),
	)

	if err != nil {
		return err
	}

	return utils.ExecuteTransaction("Reward Incentive", tgeContext.Client, tx)
}
