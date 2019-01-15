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

// BaasIncentivesABI is the input ABI used to generate the binding from.
const BaasIncentivesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesIssued\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasIncentivesBin is the compiled bytecode used for deploying new contracts.
const BaasIncentivesBin = `0x608060405234801561001057600080fd5b50604051602080610b2d8339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506a084595161401484a0000006001819055506000600281905550506109c5806101686000396000f3fe6080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063378dc3dc146100a95780633ac94f54146100d4578063715018a6146100ff578063867904b4146101165780638da5cb5b146101895780638f32d59b146101e05780639d76ea581461020f578063ab77e2d314610266578063b69ef8a814610291578063f2fde38b146102bc575b600080fd5b3480156100b557600080fd5b506100be61030d565b6040518082815260200191505060405180910390f35b3480156100e057600080fd5b506100e9610320565b6040518082815260200191505060405180910390f35b34801561010b57600080fd5b5061011461032a565b005b34801561012257600080fd5b5061016f6004803603604081101561013957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506103fc565b604051808215151515815260200191505060405180910390f35b34801561019557600080fd5b5061019e61068c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101ec57600080fd5b506101f56106b5565b604051808215151515815260200191505060405180910390f35b34801561021b57600080fd5b5061022461070c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027257600080fd5b5061027b610736565b6040518082815260200191505060405180910390f35b34801561029d57600080fd5b506102a6610740565b6040518082815260200191505060405180910390f35b3480156102c857600080fd5b5061030b600480360360208110156102df57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061083d565b005b60006a084595161401484a000000905090565b6000600254905090565b6103326106b5565b151561033d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006104066106b5565b151561041157600080fd5b6001548210151561048a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6e6f7420656e6f75676820746f6b656e206c656674000000000000000000000081525060200191505060405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561054f57600080fd5b505af1158015610563573d6000803e3d6000fd5b505050506040513d602081101561057957600080fd5b810190808051906020019092919050505015156105fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b6106138260015461085c90919063ffffffff16565b60018190555061062e8260025461087e90919063ffffffff16565b6002819055508273ffffffffffffffffffffffffffffffffffffffff167f21d739f160a7464fddaac4a1d1517d84e76b75618a053943b345c408c4160fe0836040518082815260200191505060405180910390a26001905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156107fd57600080fd5b505afa158015610811573d6000803e3d6000fd5b505050506040513d602081101561082757600080fd5b8101908080519060200190929190505050905090565b6108456106b5565b151561085057600080fd5b6108598161089f565b50565b600082821115151561086d57600080fd5b600082840390508091505092915050565b600080828401905083811015151561089557600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156108db57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fea165627a7a7230582068a2a5e12cafc348fe7e055683c45a84eadad6c716e088e01efa491f276114050029`

// DeployBaasIncentives deploys a new Ethereum contract, binding an instance of BaasIncentives to it.
func DeployBaasIncentives(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasIncentives, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasIncentivesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasIncentivesBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasIncentives{BaasIncentivesCaller: BaasIncentivesCaller{contract: contract}, BaasIncentivesTransactor: BaasIncentivesTransactor{contract: contract}, BaasIncentivesFilterer: BaasIncentivesFilterer{contract: contract}}, nil
}

// BaasIncentives is an auto generated Go binding around an Ethereum contract.
type BaasIncentives struct {
	BaasIncentivesCaller     // Read-only binding to the contract
	BaasIncentivesTransactor // Write-only binding to the contract
	BaasIncentivesFilterer   // Log filterer for contract events
}

// BaasIncentivesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasIncentivesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasIncentivesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasIncentivesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasIncentivesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasIncentivesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasIncentivesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasIncentivesSession struct {
	Contract     *BaasIncentives   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasIncentivesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasIncentivesCallerSession struct {
	Contract *BaasIncentivesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BaasIncentivesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasIncentivesTransactorSession struct {
	Contract     *BaasIncentivesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BaasIncentivesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasIncentivesRaw struct {
	Contract *BaasIncentives // Generic contract binding to access the raw methods on
}

// BaasIncentivesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasIncentivesCallerRaw struct {
	Contract *BaasIncentivesCaller // Generic read-only contract binding to access the raw methods on
}

// BaasIncentivesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasIncentivesTransactorRaw struct {
	Contract *BaasIncentivesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasIncentives creates a new instance of BaasIncentives, bound to a specific deployed contract.
func NewBaasIncentives(address common.Address, backend bind.ContractBackend) (*BaasIncentives, error) {
	contract, err := bindBaasIncentives(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasIncentives{BaasIncentivesCaller: BaasIncentivesCaller{contract: contract}, BaasIncentivesTransactor: BaasIncentivesTransactor{contract: contract}, BaasIncentivesFilterer: BaasIncentivesFilterer{contract: contract}}, nil
}

// NewBaasIncentivesCaller creates a new read-only instance of BaasIncentives, bound to a specific deployed contract.
func NewBaasIncentivesCaller(address common.Address, caller bind.ContractCaller) (*BaasIncentivesCaller, error) {
	contract, err := bindBaasIncentives(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesCaller{contract: contract}, nil
}

// NewBaasIncentivesTransactor creates a new write-only instance of BaasIncentives, bound to a specific deployed contract.
func NewBaasIncentivesTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasIncentivesTransactor, error) {
	contract, err := bindBaasIncentives(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesTransactor{contract: contract}, nil
}

// NewBaasIncentivesFilterer creates a new log filterer instance of BaasIncentives, bound to a specific deployed contract.
func NewBaasIncentivesFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasIncentivesFilterer, error) {
	contract, err := bindBaasIncentives(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesFilterer{contract: contract}, nil
}

// bindBaasIncentives binds a generic wrapper to an already deployed contract.
func bindBaasIncentives(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasIncentivesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasIncentives *BaasIncentivesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasIncentives.Contract.BaasIncentivesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasIncentives *BaasIncentivesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentives.Contract.BaasIncentivesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasIncentives *BaasIncentivesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasIncentives.Contract.BaasIncentivesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasIncentives *BaasIncentivesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasIncentives.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasIncentives *BaasIncentivesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentives.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasIncentives *BaasIncentivesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasIncentives.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) Balance() (*big.Int, error) {
	return _BaasIncentives.Contract.Balance(&_BaasIncentives.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) Balance() (*big.Int, error) {
	return _BaasIncentives.Contract.Balance(&_BaasIncentives.CallOpts)
}

// IncentivesIssued is a free data retrieval call binding the contract method 0x3ac94f54.
//
// Solidity: function incentivesIssued() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) IncentivesIssued(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "incentivesIssued")
	return *ret0, err
}

// IncentivesIssued is a free data retrieval call binding the contract method 0x3ac94f54.
//
// Solidity: function incentivesIssued() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) IncentivesIssued() (*big.Int, error) {
	return _BaasIncentives.Contract.IncentivesIssued(&_BaasIncentives.CallOpts)
}

// IncentivesIssued is a free data retrieval call binding the contract method 0x3ac94f54.
//
// Solidity: function incentivesIssued() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) IncentivesIssued() (*big.Int, error) {
	return _BaasIncentives.Contract.IncentivesIssued(&_BaasIncentives.CallOpts)
}

// IncentivesLeft is a free data retrieval call binding the contract method 0xab77e2d3.
//
// Solidity: function incentivesLeft() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) IncentivesLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "incentivesLeft")
	return *ret0, err
}

// IncentivesLeft is a free data retrieval call binding the contract method 0xab77e2d3.
//
// Solidity: function incentivesLeft() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) IncentivesLeft() (*big.Int, error) {
	return _BaasIncentives.Contract.IncentivesLeft(&_BaasIncentives.CallOpts)
}

// IncentivesLeft is a free data retrieval call binding the contract method 0xab77e2d3.
//
// Solidity: function incentivesLeft() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) IncentivesLeft() (*big.Int, error) {
	return _BaasIncentives.Contract.IncentivesLeft(&_BaasIncentives.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "initialSupply")
	return *ret0, err
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) InitialSupply() (*big.Int, error) {
	return _BaasIncentives.Contract.InitialSupply(&_BaasIncentives.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) InitialSupply() (*big.Int, error) {
	return _BaasIncentives.Contract.InitialSupply(&_BaasIncentives.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasIncentives *BaasIncentivesCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasIncentives *BaasIncentivesSession) IsOwner() (bool, error) {
	return _BaasIncentives.Contract.IsOwner(&_BaasIncentives.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasIncentives *BaasIncentivesCallerSession) IsOwner() (bool, error) {
	return _BaasIncentives.Contract.IsOwner(&_BaasIncentives.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasIncentives *BaasIncentivesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasIncentives *BaasIncentivesSession) Owner() (common.Address, error) {
	return _BaasIncentives.Contract.Owner(&_BaasIncentives.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasIncentives *BaasIncentivesCallerSession) Owner() (common.Address, error) {
	return _BaasIncentives.Contract.Owner(&_BaasIncentives.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasIncentives *BaasIncentivesCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasIncentives *BaasIncentivesSession) TokenAddress() (common.Address, error) {
	return _BaasIncentives.Contract.TokenAddress(&_BaasIncentives.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasIncentives *BaasIncentivesCallerSession) TokenAddress() (common.Address, error) {
	return _BaasIncentives.Contract.TokenAddress(&_BaasIncentives.CallOpts)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(account address, amount uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Issue(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "issue", account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(account address, amount uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Issue(&_BaasIncentives.TransactOpts, account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(account address, amount uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Issue(&_BaasIncentives.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasIncentives *BaasIncentivesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasIncentives *BaasIncentivesSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasIncentives.Contract.RenounceOwnership(&_BaasIncentives.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasIncentives *BaasIncentivesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasIncentives.Contract.RenounceOwnership(&_BaasIncentives.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasIncentives *BaasIncentivesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasIncentives *BaasIncentivesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasIncentives.Contract.TransferOwnership(&_BaasIncentives.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasIncentives *BaasIncentivesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasIncentives.Contract.TransferOwnership(&_BaasIncentives.TransactOpts, newOwner)
}

// BaasIncentivesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasIncentives contract.
type BaasIncentivesOwnershipTransferredIterator struct {
	Event *BaasIncentivesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesOwnershipTransferred)
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
		it.Event = new(BaasIncentivesOwnershipTransferred)
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
func (it *BaasIncentivesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesOwnershipTransferred represents a OwnershipTransferred event raised by the BaasIncentives contract.
type BaasIncentivesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasIncentives *BaasIncentivesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasIncentivesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesOwnershipTransferredIterator{contract: _BaasIncentives.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasIncentives *BaasIncentivesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasIncentivesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesOwnershipTransferred)
				if err := _BaasIncentives.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaasIncentivesTokensIssuedIterator is returned from FilterTokensIssued and is used to iterate over the raw logs and unpacked data for TokensIssued events raised by the BaasIncentives contract.
type BaasIncentivesTokensIssuedIterator struct {
	Event *BaasIncentivesTokensIssued // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesTokensIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesTokensIssued)
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
		it.Event = new(BaasIncentivesTokensIssued)
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
func (it *BaasIncentivesTokensIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesTokensIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesTokensIssued represents a TokensIssued event raised by the BaasIncentives contract.
type BaasIncentivesTokensIssued struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensIssued is a free log retrieval operation binding the contract event 0x21d739f160a7464fddaac4a1d1517d84e76b75618a053943b345c408c4160fe0.
//
// Solidity: e TokensIssued(account indexed address, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterTokensIssued(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesTokensIssuedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "TokensIssued", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesTokensIssuedIterator{contract: _BaasIncentives.contract, event: "TokensIssued", logs: logs, sub: sub}, nil
}

// WatchTokensIssued is a free log subscription operation binding the contract event 0x21d739f160a7464fddaac4a1d1517d84e76b75618a053943b345c408c4160fe0.
//
// Solidity: e TokensIssued(account indexed address, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchTokensIssued(opts *bind.WatchOpts, sink chan<- *BaasIncentivesTokensIssued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "TokensIssued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesTokensIssued)
				if err := _BaasIncentives.contract.UnpackLog(event, "TokensIssued", log); err != nil {
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
