package info

import (
	"fmt"
	"github.com/baas/tge-sol/deployer/contracts"
	"github.com/baas/tge-sol/deployer/deployer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ShowInfo(client *ethclient.Client, filename string) error {
	cc, err := deployer.LoadContractConfig(filename)

	cc.DAppConfig().Write("dapp_v3.json")
	fmt.Println(cc)

	if err != nil {
		return err
	}

	TokenInfo(client, cc)
	EscrowInfo(client, cc)
	PPInfo(client, cc)
	FounderInfo(client, cc)
	IncentiveInfo(client, cc)

	ROIInfo(client, cc)

	return nil
}

func TokenInfo(client *ethclient.Client, cc *deployer.ContractConfig) {
	fmt.Println("\nToken: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasToken(*cc.TokenAddress, client)

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

func EscrowInfo(client *ethclient.Client, cc *deployer.ContractConfig) {
	fmt.Println("\nEscrow: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasEscrow(*cc.EscrowAddress, client)

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

func FounderInfo(client *ethclient.Client, cc *deployer.ContractConfig) {
	fmt.Println("\nFounder: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasFounder(*cc.FounderAddress, client)

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

func IncentiveInfo(client *ethclient.Client, cc *deployer.ContractConfig) {
	fmt.Println("\nIncentives: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasIncentive(*cc.IncentivesAddress, client)

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

func PPInfo(client *ethclient.Client, cc *deployer.ContractConfig) {
	fmt.Println("\nPrivate Placement: ", cc.PPAddress.String())
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasPP(*cc.PPAddress, client)

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

func ROIInfo(client *ethclient.Client, cc *deployer.ContractConfig) {
	fmt.Println("\nROI: ")
	fmt.Println("................................................................................................")
	contract, err := contracts.NewBaasROI(*cc.ROIAddress, client)

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
