package contractinfo

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)


func process(dAppContext *web3.DAppContext) error {
	TokenInfo(dAppContext)
	EscrowInfo(dAppContext)
	PPInfo(dAppContext)
	FounderInfo(dAppContext)
	IncentiveInfo(dAppContext)
	ROIInfo(dAppContext)

	return nil
}

func TokenInfo(dAppContext *web3.DAppContext) {
	fmt.Println("\nToken: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasToken(tge.TokenAddress(dAppContext), dAppContext.Client)

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

func EscrowInfo(dAppContext *web3.DAppContext) {
	fmt.Println("\nEscrow: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasEscrow(tge.EscrowAddress(dAppContext), dAppContext.Client)

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

func FounderInfo(dAppContext *web3.DAppContext) {
	fmt.Println("\nFounder: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasFounder(tge.FounderAddress(dAppContext), dAppContext.Client)

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

func IncentiveInfo(dAppContext *web3.DAppContext) {
	fmt.Println("\nIncentives: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasIncentives(tge.IncentivesAddress(dAppContext), dAppContext.Client)

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

func PPInfo(dAppContext *web3.DAppContext) { 
	contract, err := contracts.NewBaasPP(tge.PPAddress(dAppContext), dAppContext.Client)

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

func ROIInfo(dAppContext *web3.DAppContext) {
	fmt.Println("\nROI: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasROI(tge.ROIAddress(dAppContext), dAppContext.Client)

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
