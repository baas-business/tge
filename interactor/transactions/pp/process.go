package pp

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ellsol/go-ethertypes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func process(tgeContext *utils.TGEContext, isDiscounted bool, target string, amount float64) error {
	txOps := bind.NewKeyedTransactor(tgeContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)


	contract, err := contracts.NewBaasPP(*tgeContext.ContractConfig.PPAddress, tgeContext.Client)

	if err != nil {
		return err
	}

	tokenProvided, err := contract.TotalTokenProvided(&bind.CallOpts{}, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Token already provided: ", tokenProvided)

	t := common.HexToAddress(target)

	val, err := ethertypes.NewEtherValue().FromFloat64String(fmt.Sprintf("%f", amount))

	fmt.Println("Amount: ", val)
	if err != nil {
		return err
	}

	var discountType uint8 = 2

	if isDiscounted {
		discountType = 1
	}

	log.Println("DiscountType:", discountType)


	utils.PrintColoredln("GasPrice", txOps.GasPrice)
	utils.PrintColoredln("GasLimit", txOps.GasLimit)
	tx, err := contract.ProvideToken(txOps, t, val.BigInt(), discountType)

	if err != nil {
		return err
	}
	return utils.ExecuteTransaction("Provide Token", tgeContext.Client, tx)
}
