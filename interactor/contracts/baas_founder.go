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
const BaasFounderABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"FounderChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"founder1\",\"type\":\"address\"},{\"name\":\"founder2\",\"type\":\"address\"},{\"name\":\"vestingStart\",\"type\":\"uint256\"},{\"name\":\"vestingPeriod\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddress\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"changeFounder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Founder1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Founder2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasFounderBin is the compiled bytecode used for deploying new contracts.
const BaasFounderBin = `0x608060405234801561001057600080fd5b50604051602080610c2183398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610ae1806101406000396000f3006080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100b457806316ecc686146101445780636ce6cc0a146101d3578063715018a614610238578063880cab291461024f5780638da5cb5b146102a65780638f32d59b146102fd5780639d76ea581461032c578063a192a42514610383578063b69ef8a8146103da578063f2fde38b14610405575b600080fd5b3480156100c057600080fd5b506100c9610448565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101095780820151818401526020810190506100ee565b50505050905090810190601f1680156101365780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015057600080fd5b506101b9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610485565b604051808215151515815260200191505060405180910390f35b3480156101df57600080fd5b5061021e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610532565b604051808215151515815260200191505060405180910390f35b34801561024457600080fd5b5061024d6106cd565b005b34801561025b57600080fd5b5061026461079f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102b257600080fd5b506102bb6107c9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561030957600080fd5b506103126107f2565b604051808215151515815260200191505060405180910390f35b34801561033857600080fd5b50610341610849565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561038f57600080fd5b50610398610873565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103e657600080fd5b506103ef61089d565b6040518082815260200191505060405180910390f35b34801561041157600080fd5b50610446600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061099c565b005b60606040805190810160405280600781526020017f464f554e44455200000000000000000000000000000000000000000000000000815250905090565b600061048f6107f2565b151561049a57600080fd5b84600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260038190555081600481905550949350505050565b60008061053d6107f2565b151561054857600080fd5b60008314156105bc57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610623565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b7fe1bca451dd7dde7f7e5ea0fce980d0d8f9b19ed289454fa173631c0d07760a94818585604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1600191505092915050565b6106d56107f2565b15156106e057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561095c57600080fd5b505af1158015610970573d6000803e3d6000fd5b505050506040513d602081101561098657600080fd5b8101908080519060200190929190505050905090565b6109a46107f2565b15156109af57600080fd5b6109b8816109bb565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156109f757600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582036868a7551137c0e885138da34de0cc3772be79a58cee9b898202fb5ab280ad70029`

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

// Founder1 is a free data retrieval call binding the contract method 0xa192a425.
//
// Solidity: function Founder1() constant returns(address)
func (_BaasFounder *BaasFounderCaller) Founder1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "Founder1")
	return *ret0, err
}

// Founder1 is a free data retrieval call binding the contract method 0xa192a425.
//
// Solidity: function Founder1() constant returns(address)
func (_BaasFounder *BaasFounderSession) Founder1() (common.Address, error) {
	return _BaasFounder.Contract.Founder1(&_BaasFounder.CallOpts)
}

// Founder1 is a free data retrieval call binding the contract method 0xa192a425.
//
// Solidity: function Founder1() constant returns(address)
func (_BaasFounder *BaasFounderCallerSession) Founder1() (common.Address, error) {
	return _BaasFounder.Contract.Founder1(&_BaasFounder.CallOpts)
}

// Founder2 is a free data retrieval call binding the contract method 0x880cab29.
//
// Solidity: function Founder2() constant returns(address)
func (_BaasFounder *BaasFounderCaller) Founder2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "Founder2")
	return *ret0, err
}

// Founder2 is a free data retrieval call binding the contract method 0x880cab29.
//
// Solidity: function Founder2() constant returns(address)
func (_BaasFounder *BaasFounderSession) Founder2() (common.Address, error) {
	return _BaasFounder.Contract.Founder2(&_BaasFounder.CallOpts)
}

// Founder2 is a free data retrieval call binding the contract method 0x880cab29.
//
// Solidity: function Founder2() constant returns(address)
func (_BaasFounder *BaasFounderCallerSession) Founder2() (common.Address, error) {
	return _BaasFounder.Contract.Founder2(&_BaasFounder.CallOpts)
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

// ChangeFounder is a paid mutator transaction binding the contract method 0x6ce6cc0a.
//
// Solidity: function changeFounder(newAddress address, index int256) returns(bool)
func (_BaasFounder *BaasFounderTransactor) ChangeFounder(opts *bind.TransactOpts, newAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "changeFounder", newAddress, index)
}

// ChangeFounder is a paid mutator transaction binding the contract method 0x6ce6cc0a.
//
// Solidity: function changeFounder(newAddress address, index int256) returns(bool)
func (_BaasFounder *BaasFounderSession) ChangeFounder(newAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.ChangeFounder(&_BaasFounder.TransactOpts, newAddress, index)
}

// ChangeFounder is a paid mutator transaction binding the contract method 0x6ce6cc0a.
//
// Solidity: function changeFounder(newAddress address, index int256) returns(bool)
func (_BaasFounder *BaasFounderTransactorSession) ChangeFounder(newAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.ChangeFounder(&_BaasFounder.TransactOpts, newAddress, index)
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

// Setup is a paid mutator transaction binding the contract method 0x16ecc686.
//
// Solidity: function setup(founder1 address, founder2 address, vestingStart uint256, vestingPeriod uint256) returns(bool)
func (_BaasFounder *BaasFounderTransactor) Setup(opts *bind.TransactOpts, founder1 common.Address, founder2 common.Address, vestingStart *big.Int, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "setup", founder1, founder2, vestingStart, vestingPeriod)
}

// Setup is a paid mutator transaction binding the contract method 0x16ecc686.
//
// Solidity: function setup(founder1 address, founder2 address, vestingStart uint256, vestingPeriod uint256) returns(bool)
func (_BaasFounder *BaasFounderSession) Setup(founder1 common.Address, founder2 common.Address, vestingStart *big.Int, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Setup(&_BaasFounder.TransactOpts, founder1, founder2, vestingStart, vestingPeriod)
}

// Setup is a paid mutator transaction binding the contract method 0x16ecc686.
//
// Solidity: function setup(founder1 address, founder2 address, vestingStart uint256, vestingPeriod uint256) returns(bool)
func (_BaasFounder *BaasFounderTransactorSession) Setup(founder1 common.Address, founder2 common.Address, vestingStart *big.Int, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Setup(&_BaasFounder.TransactOpts, founder1, founder2, vestingStart, vestingPeriod)
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

// BaasFounderFounderChangedIterator is returned from FilterFounderChanged and is used to iterate over the raw logs and unpacked data for FounderChanged events raised by the BaasFounder contract.
type BaasFounderFounderChangedIterator struct {
	Event *BaasFounderFounderChanged // Event containing the contract specifics and raw log

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
func (it *BaasFounderFounderChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasFounderFounderChanged)
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
		it.Event = new(BaasFounderFounderChanged)
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
func (it *BaasFounderFounderChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasFounderFounderChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasFounderFounderChanged represents a FounderChanged event raised by the BaasFounder contract.
type BaasFounderFounderChanged struct {
	OldAddress common.Address
	NewAddress common.Address
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFounderChanged is a free log retrieval operation binding the contract event 0xe1bca451dd7dde7f7e5ea0fce980d0d8f9b19ed289454fa173631c0d07760a94.
//
// Solidity: e FounderChanged(oldAddress address, newAddress address, index int256)
func (_BaasFounder *BaasFounderFilterer) FilterFounderChanged(opts *bind.FilterOpts) (*BaasFounderFounderChangedIterator, error) {

	logs, sub, err := _BaasFounder.contract.FilterLogs(opts, "FounderChanged")
	if err != nil {
		return nil, err
	}
	return &BaasFounderFounderChangedIterator{contract: _BaasFounder.contract, event: "FounderChanged", logs: logs, sub: sub}, nil
}

// WatchFounderChanged is a free log subscription operation binding the contract event 0xe1bca451dd7dde7f7e5ea0fce980d0d8f9b19ed289454fa173631c0d07760a94.
//
// Solidity: e FounderChanged(oldAddress address, newAddress address, index int256)
func (_BaasFounder *BaasFounderFilterer) WatchFounderChanged(opts *bind.WatchOpts, sink chan<- *BaasFounderFounderChanged) (event.Subscription, error) {

	logs, sub, err := _BaasFounder.contract.WatchLogs(opts, "FounderChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasFounderFounderChanged)
				if err := _BaasFounder.contract.UnpackLog(event, "FounderChanged", log); err != nil {
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
