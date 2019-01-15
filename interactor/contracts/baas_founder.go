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

// BaasFounderABI is the input ABI used to generate the binding from.
const BaasFounderABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"founderId\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"founderId\",\"type\":\"int256\"}],\"name\":\"issue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"founderId\",\"type\":\"int256\"}],\"name\":\"hasFounderReceivedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasFounderBin is the compiled bytecode used for deploying new contracts.
const BaasFounderBin = `0x608060405234801561001057600080fd5b50604051602080610d598339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610c0b8061014e6000396000f3fe608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461009e57806321cde8c71461012e578063715018a6146101a15780638da5cb5b146101b85780638f32d59b1461020f5780639d76ea581461023e578063b69ef8a814610295578063f2fde38b146102c0578063fc4dfce714610311575b600080fd5b3480156100aa57600080fd5b506100b3610364565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f35780820151818401526020810190506100d8565b50505050905090810190601f1680156101205780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013a57600080fd5b506101876004803603604081101561015157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506103a1565b604051808215151515815260200191505060405180910390f35b3480156101ad57600080fd5b506101b6610802565b005b3480156101c457600080fd5b506101cd6108d4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561021b57600080fd5b506102246108fd565b604051808215151515815260200191505060405180910390f35b34801561024a57600080fd5b50610253610954565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102a157600080fd5b506102aa61097e565b6040518082815260200191505060405180910390f35b3480156102cc57600080fd5b5061030f600480360360208110156102e357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a7b565b005b34801561031d57600080fd5b5061034a6004803603602081101561033457600080fd5b8101908080359060200190929190505050610a9a565b604051808215151515815260200191505060405180910390f35b60606040805190810160405280600781526020017f464f554e44455200000000000000000000000000000000000000000000000000815250905090565b60006103ab6108fd565b15156103b657600080fd5b6001821315151561042f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f666f756e6465724964206e6f74206578697374656e740000000000000000000081525060200191505060405180910390fd5b600080905060008314156104fb57600060149054906101000a900460ff161515156104e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f666f756e646572312068617320616c726561647920726563656976656420746f81526020017f6b656e730000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6a069e10de76676d0800000090506105b5565b600060159054906101000a900460ff161515156105a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f666f756e646572312068617320616c726561647920726563656976656420746f81526020017f6b656e730000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6a01a784379d99db4200000090505b6105bd61097e565b8111151515610634576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e6f7420656e6f75676820746f6b656e73206c6566740000000000000000000081525060200191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156106f957600080fd5b505af115801561070d573d6000803e3d6000fd5b505050506040513d602081101561072357600080fd5b810190808051906020019092919050505015156107a8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f7472616e73666572206f6620746f6b656e73206661696c65640000000000000081525060200191505060405180910390fd5b828473ffffffffffffffffffffffffffffffffffffffff167f545463979f75e44b26e038efbd7b6cb074e9aa8445c341964dfa399d79860cae836040518082815260200191505060405180910390a3600191505092915050565b61080a6108fd565b151561081557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610a3b57600080fd5b505afa158015610a4f573d6000803e3d6000fd5b505050506040513d6020811015610a6557600080fd5b8101908080519060200190929190505050905090565b610a836108fd565b1515610a8e57600080fd5b610a9781610ae5565b50565b600080821415610abb57600060149054906101000a900460ff169050610ae0565b6001821415610adb57600060159054906101000a900460ff169050610ae0565b600090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610b2157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fea165627a7a72305820583ce27583a29c57be80a6c16847fa0d109c886c82d7ab5d842d8d5ff2bf290f0029`

// DeployBaasFounder deploys a new Ethereum contract, binding an instance of BaasFounder to it.
func DeployBaasFounder(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasFounder, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasFounderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasFounderBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasFounder{BaasFounderCaller: BaasFounderCaller{contract: contract}, BaasFounderTransactor: BaasFounderTransactor{contract: contract}, BaasFounderFilterer: BaasFounderFilterer{contract: contract}}, nil
}

// BaasFounder is an auto generated Go binding around an Ethereum contract.
type BaasFounder struct {
	BaasFounderCaller     // Read-only binding to the contract
	BaasFounderTransactor // Write-only binding to the contract
	BaasFounderFilterer   // Log filterer for contract events
}

// BaasFounderCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasFounderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasFounderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasFounderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasFounderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasFounderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasFounderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasFounderSession struct {
	Contract     *BaasFounder      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasFounderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasFounderCallerSession struct {
	Contract *BaasFounderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaasFounderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasFounderTransactorSession struct {
	Contract     *BaasFounderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaasFounderRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasFounderRaw struct {
	Contract *BaasFounder // Generic contract binding to access the raw methods on
}

// BaasFounderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasFounderCallerRaw struct {
	Contract *BaasFounderCaller // Generic read-only contract binding to access the raw methods on
}

// BaasFounderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasFounderTransactorRaw struct {
	Contract *BaasFounderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasFounder creates a new instance of BaasFounder, bound to a specific deployed contract.
func NewBaasFounder(address common.Address, backend bind.ContractBackend) (*BaasFounder, error) {
	contract, err := bindBaasFounder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasFounder{BaasFounderCaller: BaasFounderCaller{contract: contract}, BaasFounderTransactor: BaasFounderTransactor{contract: contract}, BaasFounderFilterer: BaasFounderFilterer{contract: contract}}, nil
}

// NewBaasFounderCaller creates a new read-only instance of BaasFounder, bound to a specific deployed contract.
func NewBaasFounderCaller(address common.Address, caller bind.ContractCaller) (*BaasFounderCaller, error) {
	contract, err := bindBaasFounder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasFounderCaller{contract: contract}, nil
}

// NewBaasFounderTransactor creates a new write-only instance of BaasFounder, bound to a specific deployed contract.
func NewBaasFounderTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasFounderTransactor, error) {
	contract, err := bindBaasFounder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasFounderTransactor{contract: contract}, nil
}

// NewBaasFounderFilterer creates a new log filterer instance of BaasFounder, bound to a specific deployed contract.
func NewBaasFounderFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasFounderFilterer, error) {
	contract, err := bindBaasFounder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasFounderFilterer{contract: contract}, nil
}

// bindBaasFounder binds a generic wrapper to an already deployed contract.
func bindBaasFounder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasFounderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasFounder *BaasFounderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasFounder.Contract.BaasFounderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasFounder *BaasFounderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasFounder.Contract.BaasFounderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasFounder *BaasFounderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasFounder.Contract.BaasFounderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasFounder *BaasFounderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasFounder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasFounder *BaasFounderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasFounder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasFounder *BaasFounderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasFounder.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasFounder *BaasFounderCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasFounder *BaasFounderSession) Balance() (*big.Int, error) {
	return _BaasFounder.Contract.Balance(&_BaasFounder.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasFounder *BaasFounderCallerSession) Balance() (*big.Int, error) {
	return _BaasFounder.Contract.Balance(&_BaasFounder.CallOpts)
}

// HasFounderReceivedTokens is a free data retrieval call binding the contract method 0xfc4dfce7.
//
// Solidity: function hasFounderReceivedTokens(founderId int256) constant returns(bool)
func (_BaasFounder *BaasFounderCaller) HasFounderReceivedTokens(opts *bind.CallOpts, founderId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "hasFounderReceivedTokens", founderId)
	return *ret0, err
}

// HasFounderReceivedTokens is a free data retrieval call binding the contract method 0xfc4dfce7.
//
// Solidity: function hasFounderReceivedTokens(founderId int256) constant returns(bool)
func (_BaasFounder *BaasFounderSession) HasFounderReceivedTokens(founderId *big.Int) (bool, error) {
	return _BaasFounder.Contract.HasFounderReceivedTokens(&_BaasFounder.CallOpts, founderId)
}

// HasFounderReceivedTokens is a free data retrieval call binding the contract method 0xfc4dfce7.
//
// Solidity: function hasFounderReceivedTokens(founderId int256) constant returns(bool)
func (_BaasFounder *BaasFounderCallerSession) HasFounderReceivedTokens(founderId *big.Int) (bool, error) {
	return _BaasFounder.Contract.HasFounderReceivedTokens(&_BaasFounder.CallOpts, founderId)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasFounder *BaasFounderCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasFounder *BaasFounderSession) IsOwner() (bool, error) {
	return _BaasFounder.Contract.IsOwner(&_BaasFounder.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasFounder *BaasFounderCallerSession) IsOwner() (bool, error) {
	return _BaasFounder.Contract.IsOwner(&_BaasFounder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasFounder *BaasFounderCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasFounder *BaasFounderSession) Name() (string, error) {
	return _BaasFounder.Contract.Name(&_BaasFounder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasFounder *BaasFounderCallerSession) Name() (string, error) {
	return _BaasFounder.Contract.Name(&_BaasFounder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasFounder *BaasFounderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasFounder *BaasFounderSession) Owner() (common.Address, error) {
	return _BaasFounder.Contract.Owner(&_BaasFounder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasFounder *BaasFounderCallerSession) Owner() (common.Address, error) {
	return _BaasFounder.Contract.Owner(&_BaasFounder.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasFounder *BaasFounderCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasFounder *BaasFounderSession) TokenAddress() (common.Address, error) {
	return _BaasFounder.Contract.TokenAddress(&_BaasFounder.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasFounder *BaasFounderCallerSession) TokenAddress() (common.Address, error) {
	return _BaasFounder.Contract.TokenAddress(&_BaasFounder.CallOpts)
}

// Issue is a paid mutator transaction binding the contract method 0x21cde8c7.
//
// Solidity: function issue(receiver address, founderId int256) returns(bool)
func (_BaasFounder *BaasFounderTransactor) Issue(opts *bind.TransactOpts, receiver common.Address, founderId *big.Int) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "issue", receiver, founderId)
}

// Issue is a paid mutator transaction binding the contract method 0x21cde8c7.
//
// Solidity: function issue(receiver address, founderId int256) returns(bool)
func (_BaasFounder *BaasFounderSession) Issue(receiver common.Address, founderId *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Issue(&_BaasFounder.TransactOpts, receiver, founderId)
}

// Issue is a paid mutator transaction binding the contract method 0x21cde8c7.
//
// Solidity: function issue(receiver address, founderId int256) returns(bool)
func (_BaasFounder *BaasFounderTransactorSession) Issue(receiver common.Address, founderId *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Issue(&_BaasFounder.TransactOpts, receiver, founderId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasFounder *BaasFounderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasFounder *BaasFounderSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasFounder.Contract.RenounceOwnership(&_BaasFounder.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasFounder *BaasFounderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasFounder.Contract.RenounceOwnership(&_BaasFounder.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasFounder *BaasFounderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasFounder *BaasFounderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasFounder.Contract.TransferOwnership(&_BaasFounder.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasFounder *BaasFounderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasFounder.Contract.TransferOwnership(&_BaasFounder.TransactOpts, newOwner)
}

// BaasFounderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasFounder contract.
type BaasFounderOwnershipTransferredIterator struct {
	Event *BaasFounderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasFounderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasFounderOwnershipTransferred)
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
		it.Event = new(BaasFounderOwnershipTransferred)
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
func (it *BaasFounderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasFounderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasFounderOwnershipTransferred represents a OwnershipTransferred event raised by the BaasFounder contract.
type BaasFounderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasFounder *BaasFounderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasFounderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasFounder.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasFounderOwnershipTransferredIterator{contract: _BaasFounder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasFounder *BaasFounderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasFounderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasFounder.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasFounderOwnershipTransferred)
				if err := _BaasFounder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaasFounderTokensIssuedIterator is returned from FilterTokensIssued and is used to iterate over the raw logs and unpacked data for TokensIssued events raised by the BaasFounder contract.
type BaasFounderTokensIssuedIterator struct {
	Event *BaasFounderTokensIssued // Event containing the contract specifics and raw log

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
func (it *BaasFounderTokensIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasFounderTokensIssued)
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
		it.Event = new(BaasFounderTokensIssued)
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
func (it *BaasFounderTokensIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasFounderTokensIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasFounderTokensIssued represents a TokensIssued event raised by the BaasFounder contract.
type BaasFounderTokensIssued struct {
	Receiver  common.Address
	FounderId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensIssued is a free log retrieval operation binding the contract event 0x545463979f75e44b26e038efbd7b6cb074e9aa8445c341964dfa399d79860cae.
//
// Solidity: e TokensIssued(receiver indexed address, founderId indexed int256, amount uint256)
func (_BaasFounder *BaasFounderFilterer) FilterTokensIssued(opts *bind.FilterOpts, receiver []common.Address, founderId []*big.Int) (*BaasFounderTokensIssuedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var founderIdRule []interface{}
	for _, founderIdItem := range founderId {
		founderIdRule = append(founderIdRule, founderIdItem)
	}

	logs, sub, err := _BaasFounder.contract.FilterLogs(opts, "TokensIssued", receiverRule, founderIdRule)
	if err != nil {
		return nil, err
	}
	return &BaasFounderTokensIssuedIterator{contract: _BaasFounder.contract, event: "TokensIssued", logs: logs, sub: sub}, nil
}

// WatchTokensIssued is a free log subscription operation binding the contract event 0x545463979f75e44b26e038efbd7b6cb074e9aa8445c341964dfa399d79860cae.
//
// Solidity: e TokensIssued(receiver indexed address, founderId indexed int256, amount uint256)
func (_BaasFounder *BaasFounderFilterer) WatchTokensIssued(opts *bind.WatchOpts, sink chan<- *BaasFounderTokensIssued, receiver []common.Address, founderId []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var founderIdRule []interface{}
	for _, founderIdItem := range founderId {
		founderIdRule = append(founderIdRule, founderIdItem)
	}

	logs, sub, err := _BaasFounder.contract.WatchLogs(opts, "TokensIssued", receiverRule, founderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasFounderTokensIssued)
				if err := _BaasFounder.contract.UnpackLog(event, "TokensIssued", log); err != nil {
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
