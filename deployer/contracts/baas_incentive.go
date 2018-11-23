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

// BaasIncentiveABI is the input ABI used to generate the binding from.
const BaasIncentiveABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"raiseCapital\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BaasIncentiveBin is the compiled bytecode used for deploying new contracts.
const BaasIncentiveBin = `0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a361044c806100cf6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063715018a6146100725780638da5cb5b146100895780638f32d59b146100e0578063c621bf8a1461010f578063f2fde38b14610154575b600080fd5b34801561007e57600080fd5b50610087610197565b005b34801561009557600080fd5b5061009e610269565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100ec57600080fd5b506100f5610292565b604051808215151515815260200191505060405180910390f35b34801561011b57600080fd5b5061013a600480360381019080803590602001909291905050506102e9565b604051808215151515815260200191505060405180910390f35b34801561016057600080fd5b50610195600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610307565b005b61019f610292565b15156101aa57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60006102f3610292565b15156102fe57600080fd5b60019050919050565b61030f610292565b151561031a57600080fd5b61032381610326565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561036257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058205f9c55bc9fa08f4069546c0a792b2f829cae67c728d555342ebd80b58f123ee10029`

// DeployBaasIncentive deploys a new Ethereum contract, binding an instance of BaasIncentive to it.
func DeployBaasIncentive(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaasIncentive, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasIncentiveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasIncentiveBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasIncentive{BaasIncentiveCaller: BaasIncentiveCaller{contract: contract}, BaasIncentiveTransactor: BaasIncentiveTransactor{contract: contract}, BaasIncentiveFilterer: BaasIncentiveFilterer{contract: contract}}, nil
}

// BaasIncentive is an auto generated Go binding around an Ethereum contract.
type BaasIncentive struct {
	BaasIncentiveCaller     // Read-only binding to the contract
	BaasIncentiveTransactor // Write-only binding to the contract
	BaasIncentiveFilterer   // Log filterer for contract events
}

// BaasIncentiveCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasIncentiveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasIncentiveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasIncentiveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasIncentiveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasIncentiveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasIncentiveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasIncentiveSession struct {
	Contract     *BaasIncentive    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasIncentiveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasIncentiveCallerSession struct {
	Contract *BaasIncentiveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BaasIncentiveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasIncentiveTransactorSession struct {
	Contract     *BaasIncentiveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BaasIncentiveRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasIncentiveRaw struct {
	Contract *BaasIncentive // Generic contract binding to access the raw methods on
}

// BaasIncentiveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasIncentiveCallerRaw struct {
	Contract *BaasIncentiveCaller // Generic read-only contract binding to access the raw methods on
}

// BaasIncentiveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasIncentiveTransactorRaw struct {
	Contract *BaasIncentiveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasIncentive creates a new instance of BaasIncentive, bound to a specific deployed contract.
func NewBaasIncentive(address common.Address, backend bind.ContractBackend) (*BaasIncentive, error) {
	contract, err := bindBaasIncentive(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasIncentive{BaasIncentiveCaller: BaasIncentiveCaller{contract: contract}, BaasIncentiveTransactor: BaasIncentiveTransactor{contract: contract}, BaasIncentiveFilterer: BaasIncentiveFilterer{contract: contract}}, nil
}

// NewBaasIncentiveCaller creates a new read-only instance of BaasIncentive, bound to a specific deployed contract.
func NewBaasIncentiveCaller(address common.Address, caller bind.ContractCaller) (*BaasIncentiveCaller, error) {
	contract, err := bindBaasIncentive(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasIncentiveCaller{contract: contract}, nil
}

// NewBaasIncentiveTransactor creates a new write-only instance of BaasIncentive, bound to a specific deployed contract.
func NewBaasIncentiveTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasIncentiveTransactor, error) {
	contract, err := bindBaasIncentive(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasIncentiveTransactor{contract: contract}, nil
}

// NewBaasIncentiveFilterer creates a new log filterer instance of BaasIncentive, bound to a specific deployed contract.
func NewBaasIncentiveFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasIncentiveFilterer, error) {
	contract, err := bindBaasIncentive(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasIncentiveFilterer{contract: contract}, nil
}

// bindBaasIncentive binds a generic wrapper to an already deployed contract.
func bindBaasIncentive(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasIncentiveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasIncentive *BaasIncentiveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasIncentive.Contract.BaasIncentiveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasIncentive *BaasIncentiveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentive.Contract.BaasIncentiveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasIncentive *BaasIncentiveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasIncentive.Contract.BaasIncentiveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasIncentive *BaasIncentiveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasIncentive.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasIncentive *BaasIncentiveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentive.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasIncentive *BaasIncentiveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasIncentive.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasIncentive *BaasIncentiveCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasIncentive *BaasIncentiveSession) IsOwner() (bool, error) {
	return _BaasIncentive.Contract.IsOwner(&_BaasIncentive.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasIncentive *BaasIncentiveCallerSession) IsOwner() (bool, error) {
	return _BaasIncentive.Contract.IsOwner(&_BaasIncentive.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasIncentive *BaasIncentiveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasIncentive *BaasIncentiveSession) Owner() (common.Address, error) {
	return _BaasIncentive.Contract.Owner(&_BaasIncentive.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasIncentive *BaasIncentiveCallerSession) Owner() (common.Address, error) {
	return _BaasIncentive.Contract.Owner(&_BaasIncentive.CallOpts)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xc621bf8a.
//
// Solidity: function raiseCapital(amount uint256) returns(bool)
func (_BaasIncentive *BaasIncentiveTransactor) RaiseCapital(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.contract.Transact(opts, "raiseCapital", amount)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xc621bf8a.
//
// Solidity: function raiseCapital(amount uint256) returns(bool)
func (_BaasIncentive *BaasIncentiveSession) RaiseCapital(amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.Contract.RaiseCapital(&_BaasIncentive.TransactOpts, amount)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xc621bf8a.
//
// Solidity: function raiseCapital(amount uint256) returns(bool)
func (_BaasIncentive *BaasIncentiveTransactorSession) RaiseCapital(amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.Contract.RaiseCapital(&_BaasIncentive.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasIncentive *BaasIncentiveTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentive.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasIncentive *BaasIncentiveSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasIncentive.Contract.RenounceOwnership(&_BaasIncentive.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasIncentive *BaasIncentiveTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasIncentive.Contract.RenounceOwnership(&_BaasIncentive.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasIncentive *BaasIncentiveTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasIncentive.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasIncentive *BaasIncentiveSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasIncentive.Contract.TransferOwnership(&_BaasIncentive.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasIncentive *BaasIncentiveTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasIncentive.Contract.TransferOwnership(&_BaasIncentive.TransactOpts, newOwner)
}

// BaasIncentiveOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasIncentive contract.
type BaasIncentiveOwnershipTransferredIterator struct {
	Event *BaasIncentiveOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasIncentiveOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentiveOwnershipTransferred)
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
		it.Event = new(BaasIncentiveOwnershipTransferred)
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
func (it *BaasIncentiveOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentiveOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentiveOwnershipTransferred represents a OwnershipTransferred event raised by the BaasIncentive contract.
type BaasIncentiveOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasIncentive *BaasIncentiveFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasIncentiveOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasIncentive.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentiveOwnershipTransferredIterator{contract: _BaasIncentive.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasIncentive *BaasIncentiveFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasIncentiveOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasIncentive.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentiveOwnershipTransferred)
				if err := _BaasIncentive.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
