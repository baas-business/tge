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

// BaasPPABI is the input ABI used to generate the binding from.
const BaasPPABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint8\"}],\"name\":\"deliverToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"conversionRate\",\"type\":\"uint8\"}],\"name\":\"TokenDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint8\"}],\"name\":\"provideToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BaasPPBin is the compiled bytecode used for deploying new contracts.
const BaasPPBin = `0x`

// DeployBaasPP deploys a new Ethereum contract, binding an instance of BaasPP to it.
func DeployBaasPP(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasPP, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasPPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasPPBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasPP{BaasPPCaller: BaasPPCaller{contract: contract}, BaasPPTransactor: BaasPPTransactor{contract: contract}, BaasPPFilterer: BaasPPFilterer{contract: contract}}, nil
}

// BaasPP is an auto generated Go binding around an Ethereum contract.
type BaasPP struct {
	BaasPPCaller     // Read-only binding to the contract
	BaasPPTransactor // Write-only binding to the contract
	BaasPPFilterer   // Log filterer for contract events
}

// BaasPPCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasPPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasPPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasPPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasPPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasPPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasPPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasPPSession struct {
	Contract     *BaasPP           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasPPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasPPCallerSession struct {
	Contract *BaasPPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaasPPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasPPTransactorSession struct {
	Contract     *BaasPPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasPPRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasPPRaw struct {
	Contract *BaasPP // Generic contract binding to access the raw methods on
}

// BaasPPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasPPCallerRaw struct {
	Contract *BaasPPCaller // Generic read-only contract binding to access the raw methods on
}

// BaasPPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasPPTransactorRaw struct {
	Contract *BaasPPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasPP creates a new instance of BaasPP, bound to a specific deployed contract.
func NewBaasPP(address common.Address, backend bind.ContractBackend) (*BaasPP, error) {
	contract, err := bindBaasPP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasPP{BaasPPCaller: BaasPPCaller{contract: contract}, BaasPPTransactor: BaasPPTransactor{contract: contract}, BaasPPFilterer: BaasPPFilterer{contract: contract}}, nil
}

// NewBaasPPCaller creates a new read-only instance of BaasPP, bound to a specific deployed contract.
func NewBaasPPCaller(address common.Address, caller bind.ContractCaller) (*BaasPPCaller, error) {
	contract, err := bindBaasPP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasPPCaller{contract: contract}, nil
}

// NewBaasPPTransactor creates a new write-only instance of BaasPP, bound to a specific deployed contract.
func NewBaasPPTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasPPTransactor, error) {
	contract, err := bindBaasPP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasPPTransactor{contract: contract}, nil
}

// NewBaasPPFilterer creates a new log filterer instance of BaasPP, bound to a specific deployed contract.
func NewBaasPPFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasPPFilterer, error) {
	contract, err := bindBaasPP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasPPFilterer{contract: contract}, nil
}

// bindBaasPP binds a generic wrapper to an already deployed contract.
func bindBaasPP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasPPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasPP *BaasPPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasPP.Contract.BaasPPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasPP *BaasPPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.Contract.BaasPPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasPP *BaasPPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasPP.Contract.BaasPPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasPP *BaasPPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasPP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasPP *BaasPPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasPP *BaasPPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasPP.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasPP *BaasPPCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasPP *BaasPPSession) Balance() (*big.Int, error) {
	return _BaasPP.Contract.Balance(&_BaasPP.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasPP *BaasPPCallerSession) Balance() (*big.Int, error) {
	return _BaasPP.Contract.Balance(&_BaasPP.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasPP *BaasPPCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasPP *BaasPPSession) IsOwner() (bool, error) {
	return _BaasPP.Contract.IsOwner(&_BaasPP.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasPP *BaasPPCallerSession) IsOwner() (bool, error) {
	return _BaasPP.Contract.IsOwner(&_BaasPP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasPP *BaasPPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasPP *BaasPPSession) Owner() (common.Address, error) {
	return _BaasPP.Contract.Owner(&_BaasPP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasPP *BaasPPCallerSession) Owner() (common.Address, error) {
	return _BaasPP.Contract.Owner(&_BaasPP.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasPP *BaasPPCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasPP *BaasPPSession) TokenAddress() (common.Address, error) {
	return _BaasPP.Contract.TokenAddress(&_BaasPP.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasPP *BaasPPCallerSession) TokenAddress() (common.Address, error) {
	return _BaasPP.Contract.TokenAddress(&_BaasPP.CallOpts)
}

// DeliverToken is a paid mutator transaction binding the contract method 0x887d5643.
//
// Solidity: function deliverToken(account address, amount uint256, conversionRate uint8) returns(bool)
func (_BaasPP *BaasPPTransactor) DeliverToken(opts *bind.TransactOpts, account common.Address, amount *big.Int, conversionRate uint8) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "deliverToken", account, amount, conversionRate)
}

// DeliverToken is a paid mutator transaction binding the contract method 0x887d5643.
//
// Solidity: function deliverToken(account address, amount uint256, conversionRate uint8) returns(bool)
func (_BaasPP *BaasPPSession) DeliverToken(account common.Address, amount *big.Int, conversionRate uint8) (*types.Transaction, error) {
	return _BaasPP.Contract.DeliverToken(&_BaasPP.TransactOpts, account, amount, conversionRate)
}

// DeliverToken is a paid mutator transaction binding the contract method 0x887d5643.
//
// Solidity: function deliverToken(account address, amount uint256, conversionRate uint8) returns(bool)
func (_BaasPP *BaasPPTransactorSession) DeliverToken(account common.Address, amount *big.Int, conversionRate uint8) (*types.Transaction, error) {
	return _BaasPP.Contract.DeliverToken(&_BaasPP.TransactOpts, account, amount, conversionRate)
}

// ProvideToken is a paid mutator transaction binding the contract method 0xfcbcac1d.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint8) returns(bool)
func (_BaasPP *BaasPPTransactor) ProvideToken(opts *bind.TransactOpts, account common.Address, amount *big.Int, conversionRate uint8) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "provideToken", account, amount, conversionRate)
}

// ProvideToken is a paid mutator transaction binding the contract method 0xfcbcac1d.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint8) returns(bool)
func (_BaasPP *BaasPPSession) ProvideToken(account common.Address, amount *big.Int, conversionRate uint8) (*types.Transaction, error) {
	return _BaasPP.Contract.ProvideToken(&_BaasPP.TransactOpts, account, amount, conversionRate)
}

// ProvideToken is a paid mutator transaction binding the contract method 0xfcbcac1d.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint8) returns(bool)
func (_BaasPP *BaasPPTransactorSession) ProvideToken(account common.Address, amount *big.Int, conversionRate uint8) (*types.Transaction, error) {
	return _BaasPP.Contract.ProvideToken(&_BaasPP.TransactOpts, account, amount, conversionRate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasPP *BaasPPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasPP *BaasPPSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasPP.Contract.RenounceOwnership(&_BaasPP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasPP *BaasPPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasPP.Contract.RenounceOwnership(&_BaasPP.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasPP *BaasPPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasPP *BaasPPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasPP.Contract.TransferOwnership(&_BaasPP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasPP *BaasPPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasPP.Contract.TransferOwnership(&_BaasPP.TransactOpts, newOwner)
}

// BaasPPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasPP contract.
type BaasPPOwnershipTransferredIterator struct {
	Event *BaasPPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasPPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasPPOwnershipTransferred)
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
		it.Event = new(BaasPPOwnershipTransferred)
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
func (it *BaasPPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasPPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasPPOwnershipTransferred represents a OwnershipTransferred event raised by the BaasPP contract.
type BaasPPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasPP *BaasPPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasPPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasPP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasPPOwnershipTransferredIterator{contract: _BaasPP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasPP *BaasPPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasPPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasPP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasPPOwnershipTransferred)
				if err := _BaasPP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaasPPTokenDeliveredIterator is returned from FilterTokenDelivered and is used to iterate over the raw logs and unpacked data for TokenDelivered events raised by the BaasPP contract.
type BaasPPTokenDeliveredIterator struct {
	Event *BaasPPTokenDelivered // Event containing the contract specifics and raw log

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
func (it *BaasPPTokenDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasPPTokenDelivered)
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
		it.Event = new(BaasPPTokenDelivered)
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
func (it *BaasPPTokenDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasPPTokenDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasPPTokenDelivered represents a TokenDelivered event raised by the BaasPP contract.
type BaasPPTokenDelivered struct {
	To             common.Address
	Amount         *big.Int
	ConversionRate uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenDelivered is a free log retrieval operation binding the contract event 0xa254c04eefa2b3459dff6e70a8a28dc5876401ccfb247ad27c56f0c32970cf06.
//
// Solidity: e TokenDelivered(to indexed address, amount uint256, conversionRate uint8)
func (_BaasPP *BaasPPFilterer) FilterTokenDelivered(opts *bind.FilterOpts, to []common.Address) (*BaasPPTokenDeliveredIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaasPP.contract.FilterLogs(opts, "TokenDelivered", toRule)
	if err != nil {
		return nil, err
	}
	return &BaasPPTokenDeliveredIterator{contract: _BaasPP.contract, event: "TokenDelivered", logs: logs, sub: sub}, nil
}

// WatchTokenDelivered is a free log subscription operation binding the contract event 0xa254c04eefa2b3459dff6e70a8a28dc5876401ccfb247ad27c56f0c32970cf06.
//
// Solidity: e TokenDelivered(to indexed address, amount uint256, conversionRate uint8)
func (_BaasPP *BaasPPFilterer) WatchTokenDelivered(opts *bind.WatchOpts, sink chan<- *BaasPPTokenDelivered, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaasPP.contract.WatchLogs(opts, "TokenDelivered", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasPPTokenDelivered)
				if err := _BaasPP.contract.UnpackLog(event, "TokenDelivered", log); err != nil {
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
