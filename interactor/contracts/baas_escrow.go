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

// BaasEscrowABI is the input ABI used to generate the binding from.
const BaasEscrowABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CapitalRaised\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"raiseCapital\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasEscrowBin is the compiled bytecode used for deploying new contracts.
const BaasEscrowBin = `0x608060405234801561001057600080fd5b5060405160208061088883398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610748806101406000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610093578063715018a6146101235780638da5cb5b1461013a5780638f32d59b146101915780639d76ea58146101c0578063b69ef8a814610217578063c621bf8a14610242578063f2fde38b14610287575b600080fd5b34801561009f57600080fd5b506100a86102ca565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100e85780820151818401526020810190506100cd565b50505050905090810190601f1680156101155780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012f57600080fd5b50610138610307565b005b34801561014657600080fd5b5061014f6103d9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561019d57600080fd5b506101a6610402565b604051808215151515815260200191505060405180910390f35b3480156101cc57600080fd5b506101d5610459565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561022357600080fd5b5061022c610483565b6040518082815260200191505060405180910390f35b34801561024e57600080fd5b5061026d60048036038101908080359060200190929190505050610582565b604051808215151515815260200191505060405180910390f35b34801561029357600080fd5b506102c8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610603565b005b60606040805190810160405280600681526020017f455343524f570000000000000000000000000000000000000000000000000000815250905090565b61030f610402565b151561031a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561054257600080fd5b505af1158015610556573d6000803e3d6000fd5b505050506040513d602081101561056c57600080fd5b8101908080519060200190929190505050905090565b600061058c610402565b151561059757600080fd5b61059f610483565b82111515156105ad57600080fd5b6105b5610483565b82111515156105c357600080fd5b7f300283bdc9ca661e983285eff22a4decd9d0eb7ab78d357ce8e0f6533f727fa4826040518082815260200191505060405180910390a160019050919050565b61060b610402565b151561061657600080fd5b61061f81610622565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561065e57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820dfb3e9c28d89f5fc13956ccaa456d6c90c557cbb302cc65131babf7071a977800029`

// DeployBaasEscrow deploys a new Ethereum contract, binding an instance of BaasEscrow to it.
func DeployBaasEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasEscrow, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasEscrowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasEscrowBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasEscrow{BaasEscrowCaller: BaasEscrowCaller{contract: contract}, BaasEscrowTransactor: BaasEscrowTransactor{contract: contract}, BaasEscrowFilterer: BaasEscrowFilterer{contract: contract}}, nil
}

// BaasEscrow is an auto generated Go binding around an Ethereum contract.
type BaasEscrow struct {
	BaasEscrowCaller     // Read-only binding to the contract
	BaasEscrowTransactor // Write-only binding to the contract
	BaasEscrowFilterer   // Log filterer for contract events
}

// BaasEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasEscrowSession struct {
	Contract     *BaasEscrow       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasEscrowCallerSession struct {
	Contract *BaasEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BaasEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasEscrowTransactorSession struct {
	Contract     *BaasEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BaasEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasEscrowRaw struct {
	Contract *BaasEscrow // Generic contract binding to access the raw methods on
}

// BaasEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasEscrowCallerRaw struct {
	Contract *BaasEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// BaasEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasEscrowTransactorRaw struct {
	Contract *BaasEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasEscrow creates a new instance of BaasEscrow, bound to a specific deployed contract.
func NewBaasEscrow(address common.Address, backend bind.ContractBackend) (*BaasEscrow, error) {
	contract, err := bindBaasEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasEscrow{BaasEscrowCaller: BaasEscrowCaller{contract: contract}, BaasEscrowTransactor: BaasEscrowTransactor{contract: contract}, BaasEscrowFilterer: BaasEscrowFilterer{contract: contract}}, nil
}

// NewBaasEscrowCaller creates a new read-only instance of BaasEscrow, bound to a specific deployed contract.
func NewBaasEscrowCaller(address common.Address, caller bind.ContractCaller) (*BaasEscrowCaller, error) {
	contract, err := bindBaasEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowCaller{contract: contract}, nil
}

// NewBaasEscrowTransactor creates a new write-only instance of BaasEscrow, bound to a specific deployed contract.
func NewBaasEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasEscrowTransactor, error) {
	contract, err := bindBaasEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowTransactor{contract: contract}, nil
}

// NewBaasEscrowFilterer creates a new log filterer instance of BaasEscrow, bound to a specific deployed contract.
func NewBaasEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasEscrowFilterer, error) {
	contract, err := bindBaasEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowFilterer{contract: contract}, nil
}

// bindBaasEscrow binds a generic wrapper to an already deployed contract.
func bindBaasEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasEscrow *BaasEscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasEscrow.Contract.BaasEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasEscrow *BaasEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasEscrow.Contract.BaasEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasEscrow *BaasEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasEscrow.Contract.BaasEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasEscrow *BaasEscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasEscrow *BaasEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasEscrow *BaasEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasEscrow.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) Balance() (*big.Int, error) {
	return _BaasEscrow.Contract.Balance(&_BaasEscrow.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) Balance() (*big.Int, error) {
	return _BaasEscrow.Contract.Balance(&_BaasEscrow.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasEscrow *BaasEscrowCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasEscrow *BaasEscrowSession) IsOwner() (bool, error) {
	return _BaasEscrow.Contract.IsOwner(&_BaasEscrow.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasEscrow *BaasEscrowCallerSession) IsOwner() (bool, error) {
	return _BaasEscrow.Contract.IsOwner(&_BaasEscrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasEscrow *BaasEscrowCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasEscrow *BaasEscrowSession) Name() (string, error) {
	return _BaasEscrow.Contract.Name(&_BaasEscrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasEscrow *BaasEscrowCallerSession) Name() (string, error) {
	return _BaasEscrow.Contract.Name(&_BaasEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasEscrow *BaasEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasEscrow *BaasEscrowSession) Owner() (common.Address, error) {
	return _BaasEscrow.Contract.Owner(&_BaasEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasEscrow *BaasEscrowCallerSession) Owner() (common.Address, error) {
	return _BaasEscrow.Contract.Owner(&_BaasEscrow.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasEscrow *BaasEscrowCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasEscrow *BaasEscrowSession) TokenAddress() (common.Address, error) {
	return _BaasEscrow.Contract.TokenAddress(&_BaasEscrow.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasEscrow *BaasEscrowCallerSession) TokenAddress() (common.Address, error) {
	return _BaasEscrow.Contract.TokenAddress(&_BaasEscrow.CallOpts)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xc621bf8a.
//
// Solidity: function raiseCapital(amount uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) RaiseCapital(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "raiseCapital", amount)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xc621bf8a.
//
// Solidity: function raiseCapital(amount uint256) returns(bool)
func (_BaasEscrow *BaasEscrowSession) RaiseCapital(amount *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.RaiseCapital(&_BaasEscrow.TransactOpts, amount)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xc621bf8a.
//
// Solidity: function raiseCapital(amount uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) RaiseCapital(amount *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.RaiseCapital(&_BaasEscrow.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasEscrow *BaasEscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasEscrow *BaasEscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasEscrow.Contract.RenounceOwnership(&_BaasEscrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasEscrow *BaasEscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasEscrow.Contract.RenounceOwnership(&_BaasEscrow.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasEscrow *BaasEscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasEscrow *BaasEscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasEscrow.Contract.TransferOwnership(&_BaasEscrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasEscrow *BaasEscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasEscrow.Contract.TransferOwnership(&_BaasEscrow.TransactOpts, newOwner)
}

// BaasEscrowCapitalRaisedIterator is returned from FilterCapitalRaised and is used to iterate over the raw logs and unpacked data for CapitalRaised events raised by the BaasEscrow contract.
type BaasEscrowCapitalRaisedIterator struct {
	Event *BaasEscrowCapitalRaised // Event containing the contract specifics and raw log

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
func (it *BaasEscrowCapitalRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasEscrowCapitalRaised)
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
		it.Event = new(BaasEscrowCapitalRaised)
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
func (it *BaasEscrowCapitalRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasEscrowCapitalRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasEscrowCapitalRaised represents a CapitalRaised event raised by the BaasEscrow contract.
type BaasEscrowCapitalRaised struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCapitalRaised is a free log retrieval operation binding the contract event 0x300283bdc9ca661e983285eff22a4decd9d0eb7ab78d357ce8e0f6533f727fa4.
//
// Solidity: e CapitalRaised(amount uint256)
func (_BaasEscrow *BaasEscrowFilterer) FilterCapitalRaised(opts *bind.FilterOpts) (*BaasEscrowCapitalRaisedIterator, error) {

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "CapitalRaised")
	if err != nil {
		return nil, err
	}
	return &BaasEscrowCapitalRaisedIterator{contract: _BaasEscrow.contract, event: "CapitalRaised", logs: logs, sub: sub}, nil
}

// WatchCapitalRaised is a free log subscription operation binding the contract event 0x300283bdc9ca661e983285eff22a4decd9d0eb7ab78d357ce8e0f6533f727fa4.
//
// Solidity: e CapitalRaised(amount uint256)
func (_BaasEscrow *BaasEscrowFilterer) WatchCapitalRaised(opts *bind.WatchOpts, sink chan<- *BaasEscrowCapitalRaised) (event.Subscription, error) {

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "CapitalRaised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasEscrowCapitalRaised)
				if err := _BaasEscrow.contract.UnpackLog(event, "CapitalRaised", log); err != nil {
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

// BaasEscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasEscrow contract.
type BaasEscrowOwnershipTransferredIterator struct {
	Event *BaasEscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasEscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasEscrowOwnershipTransferred)
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
		it.Event = new(BaasEscrowOwnershipTransferred)
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
func (it *BaasEscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasEscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasEscrowOwnershipTransferred represents a OwnershipTransferred event raised by the BaasEscrow contract.
type BaasEscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasEscrow *BaasEscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasEscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowOwnershipTransferredIterator{contract: _BaasEscrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasEscrow *BaasEscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasEscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasEscrowOwnershipTransferred)
				if err := _BaasEscrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
