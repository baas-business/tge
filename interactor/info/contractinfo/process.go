package contractinfo

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)


func process(tgeContext *utils.TGEContext) error {
	TokenInfo(tgeContext)
	EscrowInfo(tgeContext)
	PPInfo(tgeContext)
	FounderInfo(tgeContext)
	IncentiveInfo(tgeContext)
	ROIInfo(tgeContext)

	return nil
}

func TokenInfo(tgeContext *utils.TGEContext) {
	fmt.Println("\nToken: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasToken(*tgeContext.TokenAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", name)

	totalSupply, err := contract.TotalSupply(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total Supply: ", totalSupply)

	owner, err := contract.Owner(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner: ", owner.String())
}

func EscrowInfo(tgeContext *utils.TGEContext) {
	fmt.Println("\nEscrow: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasEscrow(*tgeContext.EscrowAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", name)

	tokenAddress, err := contract.TokenAddress(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Address: ", tokenAddress.String())

	owner, err := contract.Owner(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner: ", owner.String())

	balance, err := contract.Balance(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance.String())
}

func FounderInfo(tgeContext *utils.TGEContext) {
	fmt.Println("\nFounder: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasFounder(*tgeContext.FounderAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", name)

	tokenAddress, err := contract.TokenAddress(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Address: ", tokenAddress.String())

	owner, err := contract.Owner(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner: ", owner.String())

	balance, err := contract.Balance(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance.String())
}

func IncentiveInfo(tgeContext *utils.TGEContext) {
	fmt.Println("\nIncentives: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasIncentive(*tgeContext.IncentivesAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", name)

	tokenAddress, err := contract.TokenAddress(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Address: ", tokenAddress.String())

	owner, err := contract.Owner(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner: ", owner.String())

	balance, err := contract.Balance(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance.String())
}

func PPInfo(tgeContext *utils.TGEContext) {
	fmt.Println("\nPrivate Placement: ", tgeContext.PPAddress().String())
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasPP(*tgeContext.PPAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", name)

	tokenAddress, err := contract.TokenAddress(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Address: ", tokenAddress.String())

	owner, err := contract.Owner(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner: ", owner.String())

	balance, err := contract.Balance(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance.String())
}

func ROIInfo(tgeContext *utils.TGEContext) {
	fmt.Println("\nROI: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasROI(*tgeContext.ROIAddress(), tgeContext.Client)

	if err != nil {
		log.Fatal(err)
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name: ", name)

	tokenAddress, err := contract.TokenAddress(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Address: ", tokenAddress.String())

	owner, err := contract.Owner(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Owner: ", owner.String())

	balance, err := contract.Balance(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance: ", balance.String())
}
