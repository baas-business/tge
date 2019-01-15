package pp

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/go-ethertypes"
	"github.com/ellsol/solidity-tools/utils"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func process(dAppContext *web3.DAppContext) error {
	if err := General(dAppContext); err != nil {
		return err
	}
	if err := ProvidedToken(dAppContext); err != nil {
		return err
	}
	return nil
}

func General(dAppContext *web3.DAppContext) error {
	fmt.Println(utils.ConsoleInBlue("\nGeneral Private Placement: "))
	fmt.Println(utils.ConsoleInBlue("................................................................................................"))
	contract, err := contracts.NewBaasPP(tge.PPAddress(dAppContext), dAppContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	tokenAddress, err := contract.TokenAddress(&bind.CallOpts{})
	if err != nil {
		return err
	}
	utils.PrintColoredln("Token Address", tokenAddress.String())

	owner, err := contract.Owner(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintColoredln("Owner", owner.String())

	balance, err := contract.Balance(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintColoredln("Balance", ethertypes.NewEtherValue().SetBigInt(balance).String())


	return nil
}

func ProvidedToken(dAppContext *web3.DAppContext) error {
	fmt.Println(utils.ConsoleInBlue("\nProvided Token Private Placement: "))
	fmt.Println(utils.ConsoleInBlue("................................................................................................"))
	contract, err := contracts.NewBaasPP(tge.PPAddress(dAppContext), dAppContext.Client)

	if err != nil {
		return err
	}

	tt, err := contract.TokensIssued(&bind.CallOpts{}, 0)
	if err != nil {
		return err
	}
	utils.PrintColoredln("Total Provided Token", ethertypes.NewEtherValue().SetBigInt(tt).String())

	ttd, err := contract.TokensIssued(&bind.CallOpts{}, 1)
	if err != nil {
		return err
	}
	utils.PrintColoredln("Discounted Provided Token", ethertypes.NewEtherValue().SetBigInt(ttd).String())

	ttnd, err := contract.TokensIssued(&bind.CallOpts{}, 2)
	if err != nil {
		return err
	}
	utils.PrintColoredln("Not Discounted Provided Token", ethertypes.NewEtherValue().SetBigInt(ttnd).String())

	iterator, err := contract.FilterTokensIssued(&bind.FilterOpts{
		Start: 4000000,
	}, []common.Address{dAppContext.Key.Address}, []uint8{1})

	for iterator.Next() {
		fmt.Println(iterator.Event.Account.String())
		fmt.Println(iterator.Event.Amount)
		fmt.Println(iterator.Event.DiscountType)
		fmt.Println(".........................................")
	}
	return nil
}
