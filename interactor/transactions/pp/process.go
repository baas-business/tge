package pp

import (
	"fmt"
	"github.com/baas-business/baas-cli/utils"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/go-ethertypes"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func process(dAppContext *web3.DAppContext, isDiscounted bool, target string, amount float64) error {
	txOps := bind.NewKeyedTransactor(dAppContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasPP(tge.PPAddress(dAppContext), dAppContext.Client)

	if err != nil {
		return err
	}

	tokenProvided, err := contract.TokensIssued(&bind.CallOpts{}, 0)
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
	tx, err := contract.Issue(txOps, t, val.BigInt(), discountType)

	if err != nil {
		return err
	}
	return web3.ExecuteTransaction("Provide Token", dAppContext.Client, tx)
}
