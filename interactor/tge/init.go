package tge

import (
	"fmt"
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/ellsol/solidity-tools/commands/deployer"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BaaSTokenName      = "BaaS Token"
	BaaSPPName         = "BaaS Private Placement"
	BaaSEscrowName     = "BaaS Escrow"
	BaaSIncentivesName = "BaaS Incentives"
	BaaSFounderName    = "BaaS Founder"
	BaaSROIName        = "BaaS ROI"
)

func Callback(client *ethclient.Client) *deployer.DApp {
	var tokenAddress common.Address
	return &deployer.DApp{
		Name: "BaaS TGE",
		SmartContracts: []deployer.SmartContract{
			{
				Name: BaaSTokenName,
				CreateTransaction: func(txOps *bind.TransactOpts) (common.Address, *types.Transaction, error) {
					add, trans, _, err := contracts.DeployBaasToken(txOps, client)
					tokenAddress = add
					return add, trans, err
				},
			},
			{
				Name: BaaSPPName,
				CreateTransaction: func(txOps *bind.TransactOpts) (common.Address, *types.Transaction, error) {
					add, trans, _, err := contracts.DeployBaasPP(txOps, client, tokenAddress)
					return add, trans, err
				},
			},
			{
				Name: BaaSEscrowName,
				CreateTransaction: func(txOps *bind.TransactOpts) (common.Address, *types.Transaction, error) {
					add, trans, _, err := contracts.DeployBaasEscrow(txOps, client, tokenAddress)
					return add, trans, err
				},
			},
			{
				Name: BaaSIncentivesName,
				CreateTransaction: func(txOps *bind.TransactOpts) (common.Address, *types.Transaction, error) {
					add, trans, _, err := contracts.DeployBaasIncentives(txOps, client, tokenAddress)
					return add, trans, err
				},
			},
			{
				Name: BaaSFounderName,
				CreateTransaction: func(txOps *bind.TransactOpts) (common.Address, *types.Transaction, error) {
					add, trans, _, err := contracts.DeployBaasFounder(txOps, client, tokenAddress)
					return add, trans, err
				},
			},
			{
				Name: BaaSROIName,
				CreateTransaction: func(txOps *bind.TransactOpts) (common.Address, *types.Transaction, error) {
					add, trans, _, err := contracts.DeployBaasROI(txOps, client, tokenAddress)
					return add, trans, err
				},
			},
		},
	}
}

func TokenAddress(context *web3.DAppContext) common.Address {
	fmt.Println(context.ContractConfig)
	return context.ContractConfig.GetContractByName(BaaSTokenName).Address
}

func EscrowAddress(context *web3.DAppContext) common.Address {
	return context.ContractConfig.GetContractByName(BaaSEscrowName).Address
}
func FounderAddress(context *web3.DAppContext) common.Address {
	return context.ContractConfig.GetContractByName(BaaSFounderName).Address
}

func IncentivesAddress(context *web3.DAppContext) common.Address {
	return context.ContractConfig.GetContractByName(BaaSIncentivesName).Address
}
func PPAddress(context *web3.DAppContext) common.Address {
	return context.ContractConfig.GetContractByName(BaaSPPName).Address
}

func ROIAddress(context *web3.DAppContext) common.Address {
	return context.ContractConfig.GetContractByName(BaaSROIName).Address
}
