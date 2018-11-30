package contractinfo

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/transactions/deployer"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)


func process(httpPath string, version string) error {
	client, err := utils.RPCClient(httpPath)

	if err != nil {
		return err
	}
	return ShowInfo(client, version)
}

func ShowInfo(client *ethclient.Client, version string) error {
	configFilename := fmt.Sprintf("build/%v/contract_config.json", version)
	cc, err := deployer.LoadContractConfig(configFilename)


	err = cc.ExportResult("test")

	if err != nil {
		return err
	}

	if true {
		return nil
	}

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
