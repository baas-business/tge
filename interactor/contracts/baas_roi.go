// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// BaasROIABI is the input ABI used to generate the binding from.
const BaasROIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"omittedReceiver\",\"type\":\"address\"}],\"name\":\"PayoutOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokensReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensPossessed\",\"type\":\"uint256\"}],\"name\":\"PaidOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokensProvided\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensPossessed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenHolders\",\"type\":\"uint256\"}],\"name\":\"PaidOutAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"payoutAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eligibleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"hasEnoughTokensForPayout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\"},{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"roi\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\"},{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"roiOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"currentPayoutObligation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"maxTokensToBeRewarded\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"optimalPayoutDistribution\",\"outputs\":[{\"name\":\"maxTokenToBeRewarded\",\"type\":\"uint256\"},{\"name\":\"minPayoutBalance\",\"type\":\"uint256\"},{\"name\":\"error\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasROIBin is the compiled bytecode used for deploying new contracts.
const BaasROIBin = `0x608060405234801561001057600080fd5b506040516020806117468339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506115f88061014e6000396000f3fe6080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301b32f0a146100f65780631473ada31461014f5780631da2649b146101ac57806351cff8d9146102225780635f86d7901461028b578063715018a6146102da5780637c3a00fd146102f15780638d10ca261461031c5780638da5cb5b1461036f5780638f32d59b146103c65780639358928b146103f55780639d76ea5814610420578063b69ef8a814610477578063bb731467146104a2578063be4d0aa2146104cd578063daf384191461051c578063f2fde38b1461056f575b600080fd5b34801561010257600080fd5b506101396004803603604081101561011957600080fd5b8101908080359060200190929190803590602001909291905050506105c0565b6040518082815260200191505060405180910390f35b34801561015b57600080fd5b506101886004803603602081101561017257600080fd5b8101908080359060200190929190505050610602565b60405180848152602001838152602001828152602001935050505060405180910390f35b3480156101b857600080fd5b50610205600480360360408110156101cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106ba565b604051808381526020018281526020019250505060405180910390f35b34801561022e57600080fd5b506102716004803603602081101561024557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107ce565b604051808215151515815260200191505060405180910390f35b34801561029757600080fd5b506102c4600480360360208110156102ae57600080fd5b8101908080359060200190929190505050610913565b6040518082815260200191505060405180910390f35b3480156102e657600080fd5b506102ef610960565b005b3480156102fd57600080fd5b50610306610a32565b6040518082815260200191505060405180910390f35b34801561032857600080fd5b506103556004803603602081101561033f57600080fd5b8101908080359060200190929190505050610a3b565b604051808215151515815260200191505060405180910390f35b34801561037b57600080fd5b50610384610a6b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103d257600080fd5b506103db610a94565b604051808215151515815260200191505060405180910390f35b34801561040157600080fd5b5061040a610aeb565b6040518082815260200191505060405180910390f35b34801561042c57600080fd5b50610435610bb1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561048357600080fd5b5061048c610bdb565b6040518082815260200191505060405180910390f35b3480156104ae57600080fd5b506104b7610cd8565b6040518082815260200191505060405180910390f35b3480156104d957600080fd5b50610506600480360360208110156104f057600080fd5b8101908080359060200190929190505050610db7565b6040518082815260200191505060405180910390f35b34801561052857600080fd5b506105556004803603602081101561053f57600080fd5b8101908080359060200190929190505050610dd1565b604051808215151515815260200191505060405180910390f35b34801561057b57600080fd5b506105be6004803603602081101561059257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611268565b005b6000808214156105d357600090506105fc565b6105f9826105eb60098661128790919063ffffffff16565b6112c590919063ffffffff16565b90505b92915050565b60008060008061061c8560096112ef90919063ffffffff16565b905061064b8161063d600961062f610aeb565b61128790919063ffffffff16565b6112c590919063ffffffff16565b92506106798161066b8761065d610aeb565b61128790919063ffffffff16565b6112c590919063ffffffff16565b93506106a7846106998561068b610aeb565b61131090919063ffffffff16565b61131090919063ffffffff16565b9150838383935093509350509193909250565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561077a57600080fd5b505afa15801561078e573d6000803e3d6000fd5b505050506040513d60208110156107a457600080fd5b810190808051906020019092919050505090506107c181856105c0565b8192509250509250929050565b60006107d8610a94565b15156107e357600080fd5b60006107ed610bdb565b90506000811115156107fe57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156108c357600080fd5b505af11580156108d7573d6000803e3d6000fd5b505050506040513d60208110156108ed57600080fd5b8101908080519060200190929190505050151561090957600080fd5b6001915050919050565b60008061092a6009846112ef90919063ffffffff16565b90506109588161094a8561093c610aeb565b61128790919063ffffffff16565b6112c590919063ffffffff16565b915050919050565b610968610a94565b151561097357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006009905090565b600080600080610a4a85610602565b809350819450829550505050610a5e610bdb565b8211159350505050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639358928b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b158015610b7157600080fd5b505afa158015610b85573d6000803e3d6000fd5b505050506040513d6020811015610b9b57600080fd5b8101908080519060200190929190505050905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610c9857600080fd5b505afa158015610cac573d6000803e3d6000fd5b505050506040513d6020811015610cc257600080fd5b8101908080519060200190929190505050905090565b6000610db2610ce5610bdb565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639358928b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b158015610d6957600080fd5b505afa158015610d7d573d6000803e3d6000fd5b505050506040513d6020811015610d9357600080fd5b810190808051906020019092919050505061131090919063ffffffff16565b905090565b6000610dca610dc4610cd8565b836105c0565b9050919050565b6000610ddb610a94565b1515610de657600080fd5b610def82610a3b565b1515610dfa57600080fd5b6000610e04610bdb565b9050600080600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a654cfab6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160806040518083038186803b158015610e9057600080fd5b505afa158015610ea4573d6000803e3d6000fd5b505050506040513d6080811015610eba57600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050809450819550829650839750505050506060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637af70c1f6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015610f8157600080fd5b505af1158015610f95573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610fbf57600080fd5b810190808051640100000000811115610fd757600080fd5b82810190506020810184811115610fed57600080fd5b815185602082028301116401000000008211171561100a57600080fd5b505092919050505090506000815190506000809050600080905060008090505b83811015611212576000858281518110151561104257fe5b9060200190602002015190508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806110b357508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b806110e957508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b8061111f57508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b8061115557503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156111c2577f1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb7281604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1611204565b60006111e587848151811015156111d557fe5b906020019060200201518f611332565b90506111fa81866112ef90919063ffffffff16565b9450600184019350505b50808060010191505061102a565b507fda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9828a8360405180848152602001838152602001828152602001935050505060405180910390a1505050505050505050919050565b611270610a94565b151561127b57600080fd5b611284816114d2565b50565b60008083141561129a57600090506112bf565b600082840290508284828115156112ad57fe5b041415156112ba57600080fd5b809150505b92915050565b600080821115156112d557600080fd5b600082848115156112e257fe5b0490508091505092915050565b600080828401905083811015151561130657600080fd5b8091505092915050565b600082821115151561132157600080fd5b600082840390508091505092915050565b600080600061134185856106ba565b8092508193505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561140e57600080fd5b505af1158015611422573d6000803e3d6000fd5b505050506040513d602081101561143857600080fd5b8101908080519060200190929190505050151561145457600080fd5b7f7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb858383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a1819250505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561150e57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fea165627a7a72305820ee369cdd517bdf99ec8b5d7d61d95eac61df8b75488aecbef74ca3dc9afd447c0029`

// DeployBaasROI deploys a new Ethereum contract, binding an instance of BaasROI to it.
func DeployBaasROI(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasROI, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasROIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasROIBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasROI{BaasROICaller: BaasROICaller{contract: contract}, BaasROITransactor: BaasROITransactor{contract: contract}, BaasROIFilterer: BaasROIFilterer{contract: contract}}, nil
}

// BaasROI is an auto generated Go binding around an Ethereum contract.
type BaasROI struct {
	BaasROICaller     // Read-only binding to the contract
	BaasROITransactor // Write-only binding to the contract
	BaasROIFilterer   // Log filterer for contract events
}

// BaasROICaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasROICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasROITransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasROITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasROIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasROIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasROISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasROISession struct {
	Contract     *BaasROI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasROICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasROICallerSession struct {
	Contract *BaasROICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BaasROITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasROITransactorSession struct {
	Contract     *BaasROITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BaasROIRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasROIRaw struct {
	Contract *BaasROI // Generic contract binding to access the raw methods on
}

// BaasROICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasROICallerRaw struct {
	Contract *BaasROICaller // Generic read-only contract binding to access the raw methods on
}

// BaasROITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasROITransactorRaw struct {
	Contract *BaasROITransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasROI creates a new instance of BaasROI, bound to a specific deployed contract.
func NewBaasROI(address common.Address, backend bind.ContractBackend) (*BaasROI, error) {
	contract, err := bindBaasROI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasROI{BaasROICaller: BaasROICaller{contract: contract}, BaasROITransactor: BaasROITransactor{contract: contract}, BaasROIFilterer: BaasROIFilterer{contract: contract}}, nil
}

// NewBaasROICaller creates a new read-only instance of BaasROI, bound to a specific deployed contract.
func NewBaasROICaller(address common.Address, caller bind.ContractCaller) (*BaasROICaller, error) {
	contract, err := bindBaasROI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasROICaller{contract: contract}, nil
}

// NewBaasROITransactor creates a new write-only instance of BaasROI, bound to a specific deployed contract.
func NewBaasROITransactor(address common.Address, transactor bind.ContractTransactor) (*BaasROITransactor, error) {
	contract, err := bindBaasROI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasROITransactor{contract: contract}, nil
}

// NewBaasROIFilterer creates a new log filterer instance of BaasROI, bound to a specific deployed contract.
func NewBaasROIFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasROIFilterer, error) {
	contract, err := bindBaasROI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasROIFilterer{contract: contract}, nil
}

// bindBaasROI binds a generic wrapper to an already deployed contract.
func bindBaasROI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasROIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasROI *BaasROIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasROI.Contract.BaasROICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasROI *BaasROIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasROI.Contract.BaasROITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasROI *BaasROIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasROI.Contract.BaasROITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasROI *BaasROICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasROI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasROI *BaasROITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasROI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasROI *BaasROITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasROI.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasROI *BaasROICaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasROI *BaasROISession) Balance() (*big.Int, error) {
	return _BaasROI.Contract.Balance(&_BaasROI.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) Balance() (*big.Int, error) {
	return _BaasROI.Contract.Balance(&_BaasROI.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasROI *BaasROICaller) CirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "circulatingSupply")
	return *ret0, err
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasROI *BaasROISession) CirculatingSupply() (*big.Int, error) {
	return _BaasROI.Contract.CirculatingSupply(&_BaasROI.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) CirculatingSupply() (*big.Int, error) {
	return _BaasROI.Contract.CirculatingSupply(&_BaasROI.CallOpts)
}

// CurrentPayoutObligation is a free data retrieval call binding the contract method 0xbe4d0aa2.
//
// Solidity: function currentPayoutObligation(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICaller) CurrentPayoutObligation(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "currentPayoutObligation", tokenEuroConversionRate)
	return *ret0, err
}

// CurrentPayoutObligation is a free data retrieval call binding the contract method 0xbe4d0aa2.
//
// Solidity: function currentPayoutObligation(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROISession) CurrentPayoutObligation(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.CurrentPayoutObligation(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// CurrentPayoutObligation is a free data retrieval call binding the contract method 0xbe4d0aa2.
//
// Solidity: function currentPayoutObligation(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICallerSession) CurrentPayoutObligation(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.CurrentPayoutObligation(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// EligibleToken is a free data retrieval call binding the contract method 0xbb731467.
//
// Solidity: function eligibleToken() constant returns(uint256)
func (_BaasROI *BaasROICaller) EligibleToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "eligibleToken")
	return *ret0, err
}

// EligibleToken is a free data retrieval call binding the contract method 0xbb731467.
//
// Solidity: function eligibleToken() constant returns(uint256)
func (_BaasROI *BaasROISession) EligibleToken() (*big.Int, error) {
	return _BaasROI.Contract.EligibleToken(&_BaasROI.CallOpts)
}

// EligibleToken is a free data retrieval call binding the contract method 0xbb731467.
//
// Solidity: function eligibleToken() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) EligibleToken() (*big.Int, error) {
	return _BaasROI.Contract.EligibleToken(&_BaasROI.CallOpts)
}

// HasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x8d10ca26.
//
// Solidity: function hasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROICaller) HasEnoughTokensForPayout(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "hasEnoughTokensForPayout", tokenEuroConversionRate)
	return *ret0, err
}

// HasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x8d10ca26.
//
// Solidity: function hasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROISession) HasEnoughTokensForPayout(tokenEuroConversionRate *big.Int) (bool, error) {
	return _BaasROI.Contract.HasEnoughTokensForPayout(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// HasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x8d10ca26.
//
// Solidity: function hasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROICallerSession) HasEnoughTokensForPayout(tokenEuroConversionRate *big.Int) (bool, error) {
	return _BaasROI.Contract.HasEnoughTokensForPayout(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() constant returns(uint256)
func (_BaasROI *BaasROICaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "interestRate")
	return *ret0, err
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() constant returns(uint256)
func (_BaasROI *BaasROISession) InterestRate() (*big.Int, error) {
	return _BaasROI.Contract.InterestRate(&_BaasROI.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) InterestRate() (*big.Int, error) {
	return _BaasROI.Contract.InterestRate(&_BaasROI.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasROI *BaasROICaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasROI *BaasROISession) IsOwner() (bool, error) {
	return _BaasROI.Contract.IsOwner(&_BaasROI.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasROI *BaasROICallerSession) IsOwner() (bool, error) {
	return _BaasROI.Contract.IsOwner(&_BaasROI.CallOpts)
}

// MaxTokensToBeRewarded is a free data retrieval call binding the contract method 0x5f86d790.
//
// Solidity: function maxTokensToBeRewarded(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICaller) MaxTokensToBeRewarded(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "maxTokensToBeRewarded", tokenEuroConversionRate)
	return *ret0, err
}

// MaxTokensToBeRewarded is a free data retrieval call binding the contract method 0x5f86d790.
//
// Solidity: function maxTokensToBeRewarded(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROISession) MaxTokensToBeRewarded(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.MaxTokensToBeRewarded(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// MaxTokensToBeRewarded is a free data retrieval call binding the contract method 0x5f86d790.
//
// Solidity: function maxTokensToBeRewarded(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICallerSession) MaxTokensToBeRewarded(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.MaxTokensToBeRewarded(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(maxTokenToBeRewarded uint256, minPayoutBalance uint256, error uint256)
func (_BaasROI *BaasROICaller) OptimalPayoutDistribution(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (struct {
	MaxTokenToBeRewarded *big.Int
	MinPayoutBalance     *big.Int
	Error                *big.Int
}, error) {
	ret := new(struct {
		MaxTokenToBeRewarded *big.Int
		MinPayoutBalance     *big.Int
		Error                *big.Int
	})
	out := ret
	err := _BaasROI.contract.Call(opts, out, "optimalPayoutDistribution", tokenEuroConversionRate)
	return *ret, err
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(maxTokenToBeRewarded uint256, minPayoutBalance uint256, error uint256)
func (_BaasROI *BaasROISession) OptimalPayoutDistribution(tokenEuroConversionRate *big.Int) (struct {
	MaxTokenToBeRewarded *big.Int
	MinPayoutBalance     *big.Int
	Error                *big.Int
}, error) {
	return _BaasROI.Contract.OptimalPayoutDistribution(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(maxTokenToBeRewarded uint256, minPayoutBalance uint256, error uint256)
func (_BaasROI *BaasROICallerSession) OptimalPayoutDistribution(tokenEuroConversionRate *big.Int) (struct {
	MaxTokenToBeRewarded *big.Int
	MinPayoutBalance     *big.Int
	Error                *big.Int
}, error) {
	return _BaasROI.Contract.OptimalPayoutDistribution(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasROI *BaasROICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasROI *BaasROISession) Owner() (common.Address, error) {
	return _BaasROI.Contract.Owner(&_BaasROI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasROI *BaasROICallerSession) Owner() (common.Address, error) {
	return _BaasROI.Contract.Owner(&_BaasROI.CallOpts)
}

// Roi is a free data retrieval call binding the contract method 0x01b32f0a.
//
// Solidity: function roi(token uint256, tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICaller) Roi(opts *bind.CallOpts, token *big.Int, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "roi", token, tokenEuroConversionRate)
	return *ret0, err
}

// Roi is a free data retrieval call binding the contract method 0x01b32f0a.
//
// Solidity: function roi(token uint256, tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROISession) Roi(token *big.Int, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.Roi(&_BaasROI.CallOpts, token, tokenEuroConversionRate)
}

// Roi is a free data retrieval call binding the contract method 0x01b32f0a.
//
// Solidity: function roi(token uint256, tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICallerSession) Roi(token *big.Int, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.Roi(&_BaasROI.CallOpts, token, tokenEuroConversionRate)
}

// RoiOf is a free data retrieval call binding the contract method 0x1da2649b.
//
// Solidity: function roiOf(wallet address, tokenEuroConversionRate uint256) constant returns(uint256, uint256)
func (_BaasROI *BaasROICaller) RoiOf(opts *bind.CallOpts, wallet common.Address, tokenEuroConversionRate *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BaasROI.contract.Call(opts, out, "roiOf", wallet, tokenEuroConversionRate)
	return *ret0, *ret1, err
}

// RoiOf is a free data retrieval call binding the contract method 0x1da2649b.
//
// Solidity: function roiOf(wallet address, tokenEuroConversionRate uint256) constant returns(uint256, uint256)
func (_BaasROI *BaasROISession) RoiOf(wallet common.Address, tokenEuroConversionRate *big.Int) (*big.Int, *big.Int, error) {
	return _BaasROI.Contract.RoiOf(&_BaasROI.CallOpts, wallet, tokenEuroConversionRate)
}

// RoiOf is a free data retrieval call binding the contract method 0x1da2649b.
//
// Solidity: function roiOf(wallet address, tokenEuroConversionRate uint256) constant returns(uint256, uint256)
func (_BaasROI *BaasROICallerSession) RoiOf(wallet common.Address, tokenEuroConversionRate *big.Int) (*big.Int, *big.Int, error) {
	return _BaasROI.Contract.RoiOf(&_BaasROI.CallOpts, wallet, tokenEuroConversionRate)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasROI *BaasROICaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasROI *BaasROISession) TokenAddress() (common.Address, error) {
	return _BaasROI.Contract.TokenAddress(&_BaasROI.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasROI *BaasROICallerSession) TokenAddress() (common.Address, error) {
	return _BaasROI.Contract.TokenAddress(&_BaasROI.CallOpts)
}

// PayoutAll is a paid mutator transaction binding the contract method 0xdaf38419.
//
// Solidity: function payoutAll(tokenEuroConversionRate uint256) returns(bool)
func (_BaasROI *BaasROITransactor) PayoutAll(opts *bind.TransactOpts, tokenEuroConversionRate *big.Int) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "payoutAll", tokenEuroConversionRate)
}

// PayoutAll is a paid mutator transaction binding the contract method 0xdaf38419.
//
// Solidity: function payoutAll(tokenEuroConversionRate uint256) returns(bool)
func (_BaasROI *BaasROISession) PayoutAll(tokenEuroConversionRate *big.Int) (*types.Transaction, error) {
	return _BaasROI.Contract.PayoutAll(&_BaasROI.TransactOpts, tokenEuroConversionRate)
}

// PayoutAll is a paid mutator transaction binding the contract method 0xdaf38419.
//
// Solidity: function payoutAll(tokenEuroConversionRate uint256) returns(bool)
func (_BaasROI *BaasROITransactorSession) PayoutAll(tokenEuroConversionRate *big.Int) (*types.Transaction, error) {
	return _BaasROI.Contract.PayoutAll(&_BaasROI.TransactOpts, tokenEuroConversionRate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasROI *BaasROITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasROI *BaasROISession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasROI.Contract.RenounceOwnership(&_BaasROI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasROI *BaasROITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasROI.Contract.RenounceOwnership(&_BaasROI.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasROI *BaasROITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasROI *BaasROISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.TransferOwnership(&_BaasROI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasROI *BaasROITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.TransferOwnership(&_BaasROI.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(receiver address) returns(bool)
func (_BaasROI *BaasROITransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "withdraw", receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(receiver address) returns(bool)
func (_BaasROI *BaasROISession) Withdraw(receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.Withdraw(&_BaasROI.TransactOpts, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(receiver address) returns(bool)
func (_BaasROI *BaasROITransactorSession) Withdraw(receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.Withdraw(&_BaasROI.TransactOpts, receiver)
}

// BaasROIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasROI contract.
type BaasROIOwnershipTransferredIterator struct {
	Event *BaasROIOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaasROIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaasROIOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaasROIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIOwnershipTransferred represents a OwnershipTransferred event raised by the BaasROI contract.
type BaasROIOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasROI *BaasROIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasROIOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasROIOwnershipTransferredIterator{contract: _BaasROI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasROI *BaasROIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasROIOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIOwnershipTransferred)
				if err := _BaasROI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BaasROIPaidOutIterator is returned from FilterPaidOut and is used to iterate over the raw logs and unpacked data for PaidOut events raised by the BaasROI contract.
type BaasROIPaidOutIterator struct {
	Event *BaasROIPaidOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaasROIPaidOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIPaidOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaasROIPaidOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaasROIPaidOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIPaidOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIPaidOut represents a PaidOut event raised by the BaasROI contract.
type BaasROIPaidOut struct {
	Receiver        common.Address
	TokensReceived  *big.Int
	TokensPossessed *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaidOut is a free log retrieval operation binding the contract event 0x7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb.
//
// Solidity: e PaidOut(receiver address, tokensReceived uint256, tokensPossessed uint256)
func (_BaasROI *BaasROIFilterer) FilterPaidOut(opts *bind.FilterOpts) (*BaasROIPaidOutIterator, error) {

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PaidOut")
	if err != nil {
		return nil, err
	}
	return &BaasROIPaidOutIterator{contract: _BaasROI.contract, event: "PaidOut", logs: logs, sub: sub}, nil
}

// WatchPaidOut is a free log subscription operation binding the contract event 0x7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb.
//
// Solidity: e PaidOut(receiver address, tokensReceived uint256, tokensPossessed uint256)
func (_BaasROI *BaasROIFilterer) WatchPaidOut(opts *bind.WatchOpts, sink chan<- *BaasROIPaidOut) (event.Subscription, error) {

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PaidOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIPaidOut)
				if err := _BaasROI.contract.UnpackLog(event, "PaidOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BaasROIPaidOutAllIterator is returned from FilterPaidOutAll and is used to iterate over the raw logs and unpacked data for PaidOutAll events raised by the BaasROI contract.
type BaasROIPaidOutAllIterator struct {
	Event *BaasROIPaidOutAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaasROIPaidOutAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIPaidOutAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaasROIPaidOutAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaasROIPaidOutAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIPaidOutAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIPaidOutAll represents a PaidOutAll event raised by the BaasROI contract.
type BaasROIPaidOutAll struct {
	TokensProvided  *big.Int
	TokensPossessed *big.Int
	TokenHolders    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaidOutAll is a free log retrieval operation binding the contract event 0xda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9.
//
// Solidity: e PaidOutAll(tokensProvided uint256, tokensPossessed uint256, tokenHolders uint256)
func (_BaasROI *BaasROIFilterer) FilterPaidOutAll(opts *bind.FilterOpts) (*BaasROIPaidOutAllIterator, error) {

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PaidOutAll")
	if err != nil {
		return nil, err
	}
	return &BaasROIPaidOutAllIterator{contract: _BaasROI.contract, event: "PaidOutAll", logs: logs, sub: sub}, nil
}

// WatchPaidOutAll is a free log subscription operation binding the contract event 0xda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9.
//
// Solidity: e PaidOutAll(tokensProvided uint256, tokensPossessed uint256, tokenHolders uint256)
func (_BaasROI *BaasROIFilterer) WatchPaidOutAll(opts *bind.WatchOpts, sink chan<- *BaasROIPaidOutAll) (event.Subscription, error) {

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PaidOutAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIPaidOutAll)
				if err := _BaasROI.contract.UnpackLog(event, "PaidOutAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BaasROIPayoutOmittedIterator is returned from FilterPayoutOmitted and is used to iterate over the raw logs and unpacked data for PayoutOmitted events raised by the BaasROI contract.
type BaasROIPayoutOmittedIterator struct {
	Event *BaasROIPayoutOmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BaasROIPayoutOmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIPayoutOmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BaasROIPayoutOmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BaasROIPayoutOmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIPayoutOmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIPayoutOmitted represents a PayoutOmitted event raised by the BaasROI contract.
type BaasROIPayoutOmitted struct {
	OmittedReceiver common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPayoutOmitted is a free log retrieval operation binding the contract event 0x1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb72.
//
// Solidity: e PayoutOmitted(omittedReceiver address)
func (_BaasROI *BaasROIFilterer) FilterPayoutOmitted(opts *bind.FilterOpts) (*BaasROIPayoutOmittedIterator, error) {

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PayoutOmitted")
	if err != nil {
		return nil, err
	}
	return &BaasROIPayoutOmittedIterator{contract: _BaasROI.contract, event: "PayoutOmitted", logs: logs, sub: sub}, nil
}

// WatchPayoutOmitted is a free log subscription operation binding the contract event 0x1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb72.
//
// Solidity: e PayoutOmitted(omittedReceiver address)
func (_BaasROI *BaasROIFilterer) WatchPayoutOmitted(opts *bind.WatchOpts, sink chan<- *BaasROIPayoutOmitted) (event.Subscription, error) {

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PayoutOmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIPayoutOmitted)
				if err := _BaasROI.contract.UnpackLog(event, "PayoutOmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
