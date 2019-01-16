package contractinfo

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
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

	log.Println("Baas Token Address", tge.TokenAddress(dAppContext).String())
	_escrowAddress, _ppAddress, _founderAddress, _incentivesAddress, err := contract.Pots(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pots: ")
	fmt.Println("................................................................................................")
	fmt.Println("Escrow Address: ", _escrowAddress.String())
	fmt.Println("PP Address: ", _ppAddress.String())
	fmt.Println("Founder Address: ", _founderAddress.String())
	fmt.Println("Incentives Address: ", _incentivesAddress.String())
	fmt.Println("................................................................................................")


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


	founder1Received, err := contract.HasFounderReceivedTokens(&bind.CallOpts{}, new(big.Int).SetInt64(0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Founder 1 recieved: ", founder1Received)

	founder2Received, err := contract.HasFounderReceivedTokens(&bind.CallOpts{}, new(big.Int).SetInt64(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Founder 2 recieved: ", founder2Received)
}

func IncentiveInfo(dAppContext *web3.DAppContext) {
	fmt.Println("\nIncentives: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasIncentives(tge.IncentivesAddress(dAppContext), dAppContext.Client)

	if err != nil {
		log.Fatal(err)
	}


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
