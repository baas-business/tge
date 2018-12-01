package pp

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ellsol/go-ethertypes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func process(tgeContext *utils.TGEContext) error {
	if err := General(tgeContext); err != nil {
		return err
	}
	if err := ProvidedToken(tgeContext); err != nil {
		return err
	}
	return nil
}

func General(tgeContext *utils.TGEContext) error {
	fmt.Println(utils.ConsoleInBlue("\nGeneral Private Placement: "))
	fmt.Println(utils.ConsoleInBlue("................................................................................................"))
	contract, err := contracts.NewBaasPP(*tgeContext.PPAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		return err
	}
	utils.PrintColoredln("Name", name)

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

func ProvidedToken(tgeContext *utils.TGEContext) error {
	fmt.Println(utils.ConsoleInBlue("\nProvided Token Private Placement: "))
	fmt.Println(utils.ConsoleInBlue("................................................................................................"))
	contract, err := contracts.NewBaasPP(*tgeContext.PPAddress(), tgeContext.Client)

	if err != nil {
		return err
	}

	tt, err := contract.TotalTokenProvided(&bind.CallOpts{}, 0)
	if err != nil {
		return err
	}
	utils.PrintColoredln("Total Provided Token", ethertypes.NewEtherValue().SetBigInt(tt).String())

	ttd, err := contract.TotalTokenProvided(&bind.CallOpts{}, 1)
	if err != nil {
		return err
	}
	utils.PrintColoredln("Discounted Provided Token", ethertypes.NewEtherValue().SetBigInt(ttd).String())

	ttnd, err := contract.TotalTokenProvided(&bind.CallOpts{}, 2)
	if err != nil {
		return err
	}
	utils.PrintColoredln("Not Discounted Provided Token", ethertypes.NewEtherValue().SetBigInt(ttnd).String())

	iterator, err := contract.FilterTokenDelivered(&bind.FilterOpts{
		Start: 4000000,
	}, []common.Address{tgeContext.Key.Address})

	for iterator.Next() {
		fmt.Println(iterator.Event.To.String())
		fmt.Println(iterator.Event.Amount)
		fmt.Println(iterator.Event.DiscountType)
		fmt.Println(".........................................")
	}
	return nil
}
