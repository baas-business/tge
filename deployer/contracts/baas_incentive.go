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
const BaasIncentiveABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncentiveProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stage\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"vestingPeriodInBlocks\",\"type\":\"uint256\"},{\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideIncentive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesProvided\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingPeriodInBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getIncentive\",\"outputs\":[{\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"stage\",\"type\":\"uint256\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"stage\",\"type\":\"uint256\"}],\"name\":\"canWithdraw\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"alreadyWithdrawn\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BaasIncentiveBin is the compiled bytecode used for deploying new contracts.
const BaasIncentiveBin = `0x60806040526000600560006101000a81548160ff02191690831515021790555034801561002b57600080fd5b5060405160208061135983398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506111fe8061015b6000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630a8fed89146100e0578063392e53cd146101175780633ccfd60b146101465780635742a36e14610175578063715018a6146101a05780638da5cb5b146101b75780638f32d59b1461020e5780638f70c2d81461023d5780639d76ea58146102b4578063ab77e2d31461030b578063b69ef8a814610336578063cb5595bb14610361578063cd3a2a3c146103c6578063e2ab7cf2146103f1578063f2fde38b14610464575b600080fd5b3480156100ec57600080fd5b5061011560048036038101908080359060200190929190803590602001909291905050506104a7565b005b34801561012357600080fd5b5061012c6104dd565b604051808215151515815260200191505060405180910390f35b34801561015257600080fd5b5061015b6104f4565b604051808215151515815260200191505060405180910390f35b34801561018157600080fd5b5061018a6108b3565b6040518082815260200191505060405180910390f35b3480156101ac57600080fd5b506101b56108bd565b005b3480156101c357600080fd5b506101cc61098f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561021a57600080fd5b506102236109b8565b604051808215151515815260200191505060405180910390f35b34801561024957600080fd5b5061027e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a0f565b60405180868152602001858152602001848152602001838152602001821515151581526020019550505050505060405180910390f35b3480156102c057600080fd5b506102c9610b32565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561031757600080fd5b50610320610b5c565b6040518082815260200191505060405180910390f35b34801561034257600080fd5b5061034b610b66565b6040518082815260200191505060405180910390f35b34801561036d57600080fd5b506103ac600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c65565b604051808215151515815260200191505060405180910390f35b3480156103d257600080fd5b506103db610e8f565b6040518082815260200191505060405180910390f35b3480156103fd57600080fd5b5061043c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e99565b6040518084815260200183815260200182151515158152602001935050505060405180910390f35b34801561047057600080fd5b506104a5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fea565b005b600560009054906101000a900460ff161515156104c357600080fd5b806002819055506000600381905550816004819055505050565b6000600560009054906101000a900460ff16905090565b6000806000806000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160009054906101000a900460ff16151561055757600080fd5b6004600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101548115156105a557fe5b049350600454600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015443038115156105f957fe5b04925060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401548403019150610658828561100990919063ffffffff16565b905060008211151561066957600080fd5b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561072e57600080fd5b505af1158015610742573d6000803e3d6000fd5b505050506040513d602081101561075857600080fd5b8101908080519060200190929190505050151561077457600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206004018190555061081081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015461104790919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055503373ffffffffffffffffffffffffffffffffffffffff167f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc68284604051808381526020018281526020019250505060405180910390a25050505090565b6000600354905090565b6108c56109b8565b15156108d057600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000806000806000610a1f611183565b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060c060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050905080602001518160400151826060015183608001518460a00151955095509550955095505091939590929450565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600254905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610c2557600080fd5b505af1158015610c39573d6000803e3d6000fd5b505050506040513d6020811015610c4f57600080fd5b8101908080519060200190929190505050905090565b6000610c6f6109b8565b1515610c7a57600080fd5b6002548211151515610c8b57600080fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160009054906101000a900460ff16151515610ce757600080fd5b610cfc8260025461106890919063ffffffff16565b600281905550610d178260035461104790919063ffffffff16565b60038190555060c0604051908101604052808473ffffffffffffffffffffffffffffffffffffffff168152602001838152602001600081526020014381526020016000815260200160011515815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff0219169083151502179055509050508273ffffffffffffffffffffffffffffffffffffffff167f32dd062c233d622f3ce5c069f841701219c1c992202aef2d7a75159790bcd070836040518082815260200191505060405180910390a26001905092915050565b6000600454905090565b6000806000610ea6611183565b6000851415610ec5576000806000829250819150935093509350610fe2565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060c060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050905060048160200151811515610fbb57fe5b04935060045460018603028160600151019250848160800151101591508383839350935093505b509250925092565b610ff26109b8565b1515610ffd57600080fd5b61100681611089565b50565b600080600084141561101e5760009150611040565b828402905082848281151561102f57fe5b0414151561103c57600080fd5b8091505b5092915050565b600080828401905083811015151561105e57600080fd5b8091505092915050565b60008083831115151561107a57600080fd5b82840390508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156110c557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60c060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000815260200160001515815250905600a165627a7a723058202b45943b1866841464eacc0652da4a194566d5d8affd32ad087c3b8615146f010029`

// DeployBaasIncentive deploys a new Ethereum contract, binding an instance of BaasIncentive to it.
func DeployBaasIncentive(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasIncentive, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasIncentiveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasIncentiveBin), backend, token)
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

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveSession) Balance() (*big.Int, error) {
	return _BaasIncentive.Contract.Balance(&_BaasIncentive.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCallerSession) Balance() (*big.Int, error) {
	return _BaasIncentive.Contract.Balance(&_BaasIncentive.CallOpts)
}

// CanWithdraw is a free data retrieval call binding the contract method 0xe2ab7cf2.
//
// Solidity: function canWithdraw(account address, stage uint256) constant returns(amount uint256, atBlock uint256, alreadyWithdrawn bool)
func (_BaasIncentive *BaasIncentiveCaller) CanWithdraw(opts *bind.CallOpts, account common.Address, stage *big.Int) (struct {
	Amount           *big.Int
	AtBlock          *big.Int
	AlreadyWithdrawn bool
}, error) {
	ret := new(struct {
		Amount           *big.Int
		AtBlock          *big.Int
		AlreadyWithdrawn bool
	})
	out := ret
	err := _BaasIncentive.contract.Call(opts, out, "canWithdraw", account, stage)
	return *ret, err
}

// CanWithdraw is a free data retrieval call binding the contract method 0xe2ab7cf2.
//
// Solidity: function canWithdraw(account address, stage uint256) constant returns(amount uint256, atBlock uint256, alreadyWithdrawn bool)
func (_BaasIncentive *BaasIncentiveSession) CanWithdraw(account common.Address, stage *big.Int) (struct {
	Amount           *big.Int
	AtBlock          *big.Int
	AlreadyWithdrawn bool
}, error) {
	return _BaasIncentive.Contract.CanWithdraw(&_BaasIncentive.CallOpts, account, stage)
}

// CanWithdraw is a free data retrieval call binding the contract method 0xe2ab7cf2.
//
// Solidity: function canWithdraw(account address, stage uint256) constant returns(amount uint256, atBlock uint256, alreadyWithdrawn bool)
func (_BaasIncentive *BaasIncentiveCallerSession) CanWithdraw(account common.Address, stage *big.Int) (struct {
	Amount           *big.Int
	AtBlock          *big.Int
	AlreadyWithdrawn bool
}, error) {
	return _BaasIncentive.Contract.CanWithdraw(&_BaasIncentive.CallOpts, account, stage)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, stage uint256, isValue bool)
func (_BaasIncentive *BaasIncentiveCaller) GetIncentive(opts *bind.CallOpts, account common.Address) (struct {
	InitialAmount *big.Int
	Withdrawn     *big.Int
	AtBlock       *big.Int
	Stage         *big.Int
	IsValue       bool
}, error) {
	ret := new(struct {
		InitialAmount *big.Int
		Withdrawn     *big.Int
		AtBlock       *big.Int
		Stage         *big.Int
		IsValue       bool
	})
	out := ret
	err := _BaasIncentive.contract.Call(opts, out, "getIncentive", account)
	return *ret, err
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, stage uint256, isValue bool)
func (_BaasIncentive *BaasIncentiveSession) GetIncentive(account common.Address) (struct {
	InitialAmount *big.Int
	Withdrawn     *big.Int
	AtBlock       *big.Int
	Stage         *big.Int
	IsValue       bool
}, error) {
	return _BaasIncentive.Contract.GetIncentive(&_BaasIncentive.CallOpts, account)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, stage uint256, isValue bool)
func (_BaasIncentive *BaasIncentiveCallerSession) GetIncentive(account common.Address) (struct {
	InitialAmount *big.Int
	Withdrawn     *big.Int
	AtBlock       *big.Int
	Stage         *big.Int
	IsValue       bool
}, error) {
	return _BaasIncentive.Contract.GetIncentive(&_BaasIncentive.CallOpts, account)
}

// IncentivesLeft is a free data retrieval call binding the contract method 0xab77e2d3.
//
// Solidity: function incentivesLeft() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCaller) IncentivesLeft(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "incentivesLeft")
	return *ret0, err
}

// IncentivesLeft is a free data retrieval call binding the contract method 0xab77e2d3.
//
// Solidity: function incentivesLeft() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveSession) IncentivesLeft() (*big.Int, error) {
	return _BaasIncentive.Contract.IncentivesLeft(&_BaasIncentive.CallOpts)
}

// IncentivesLeft is a free data retrieval call binding the contract method 0xab77e2d3.
//
// Solidity: function incentivesLeft() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCallerSession) IncentivesLeft() (*big.Int, error) {
	return _BaasIncentive.Contract.IncentivesLeft(&_BaasIncentive.CallOpts)
}

// IncentivesProvided is a free data retrieval call binding the contract method 0x5742a36e.
//
// Solidity: function incentivesProvided() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCaller) IncentivesProvided(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "incentivesProvided")
	return *ret0, err
}

// IncentivesProvided is a free data retrieval call binding the contract method 0x5742a36e.
//
// Solidity: function incentivesProvided() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveSession) IncentivesProvided() (*big.Int, error) {
	return _BaasIncentive.Contract.IncentivesProvided(&_BaasIncentive.CallOpts)
}

// IncentivesProvided is a free data retrieval call binding the contract method 0x5742a36e.
//
// Solidity: function incentivesProvided() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCallerSession) IncentivesProvided() (*big.Int, error) {
	return _BaasIncentive.Contract.IncentivesProvided(&_BaasIncentive.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasIncentive *BaasIncentiveCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasIncentive *BaasIncentiveSession) IsInitialized() (bool, error) {
	return _BaasIncentive.Contract.IsInitialized(&_BaasIncentive.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasIncentive *BaasIncentiveCallerSession) IsInitialized() (bool, error) {
	return _BaasIncentive.Contract.IsInitialized(&_BaasIncentive.CallOpts)
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

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasIncentive *BaasIncentiveCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasIncentive *BaasIncentiveSession) TokenAddress() (common.Address, error) {
	return _BaasIncentive.Contract.TokenAddress(&_BaasIncentive.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasIncentive *BaasIncentiveCallerSession) TokenAddress() (common.Address, error) {
	return _BaasIncentive.Contract.TokenAddress(&_BaasIncentive.CallOpts)
}

// VestingPeriodInBlocks is a free data retrieval call binding the contract method 0xcd3a2a3c.
//
// Solidity: function vestingPeriodInBlocks() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCaller) VestingPeriodInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentive.contract.Call(opts, out, "vestingPeriodInBlocks")
	return *ret0, err
}

// VestingPeriodInBlocks is a free data retrieval call binding the contract method 0xcd3a2a3c.
//
// Solidity: function vestingPeriodInBlocks() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveSession) VestingPeriodInBlocks() (*big.Int, error) {
	return _BaasIncentive.Contract.VestingPeriodInBlocks(&_BaasIncentive.CallOpts)
}

// VestingPeriodInBlocks is a free data retrieval call binding the contract method 0xcd3a2a3c.
//
// Solidity: function vestingPeriodInBlocks() constant returns(uint256)
func (_BaasIncentive *BaasIncentiveCallerSession) VestingPeriodInBlocks() (*big.Int, error) {
	return _BaasIncentive.Contract.VestingPeriodInBlocks(&_BaasIncentive.CallOpts)
}

// ProvideIncentive is a paid mutator transaction binding the contract method 0xcb5595bb.
//
// Solidity: function provideIncentive(account address, amount uint256) returns(bool)
func (_BaasIncentive *BaasIncentiveTransactor) ProvideIncentive(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.contract.Transact(opts, "provideIncentive", account, amount)
}

// ProvideIncentive is a paid mutator transaction binding the contract method 0xcb5595bb.
//
// Solidity: function provideIncentive(account address, amount uint256) returns(bool)
func (_BaasIncentive *BaasIncentiveSession) ProvideIncentive(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.Contract.ProvideIncentive(&_BaasIncentive.TransactOpts, account, amount)
}

// ProvideIncentive is a paid mutator transaction binding the contract method 0xcb5595bb.
//
// Solidity: function provideIncentive(account address, amount uint256) returns(bool)
func (_BaasIncentive *BaasIncentiveTransactorSession) ProvideIncentive(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.Contract.ProvideIncentive(&_BaasIncentive.TransactOpts, account, amount)
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

// Setup is a paid mutator transaction binding the contract method 0x0a8fed89.
//
// Solidity: function setup(vestingPeriodInBlocks uint256, supply uint256) returns()
func (_BaasIncentive *BaasIncentiveTransactor) Setup(opts *bind.TransactOpts, vestingPeriodInBlocks *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.contract.Transact(opts, "setup", vestingPeriodInBlocks, supply)
}

// Setup is a paid mutator transaction binding the contract method 0x0a8fed89.
//
// Solidity: function setup(vestingPeriodInBlocks uint256, supply uint256) returns()
func (_BaasIncentive *BaasIncentiveSession) Setup(vestingPeriodInBlocks *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.Contract.Setup(&_BaasIncentive.TransactOpts, vestingPeriodInBlocks, supply)
}

// Setup is a paid mutator transaction binding the contract method 0x0a8fed89.
//
// Solidity: function setup(vestingPeriodInBlocks uint256, supply uint256) returns()
func (_BaasIncentive *BaasIncentiveTransactorSession) Setup(vestingPeriodInBlocks *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _BaasIncentive.Contract.Setup(&_BaasIncentive.TransactOpts, vestingPeriodInBlocks, supply)
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

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_BaasIncentive *BaasIncentiveTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentive.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_BaasIncentive *BaasIncentiveSession) Withdraw() (*types.Transaction, error) {
	return _BaasIncentive.Contract.Withdraw(&_BaasIncentive.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_BaasIncentive *BaasIncentiveTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BaasIncentive.Contract.Withdraw(&_BaasIncentive.TransactOpts)
}

// BaasIncentiveIncentiveProvidedIterator is returned from FilterIncentiveProvided and is used to iterate over the raw logs and unpacked data for IncentiveProvided events raised by the BaasIncentive contract.
type BaasIncentiveIncentiveProvidedIterator struct {
	Event *BaasIncentiveIncentiveProvided // Event containing the contract specifics and raw log

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
func (it *BaasIncentiveIncentiveProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentiveIncentiveProvided)
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
		it.Event = new(BaasIncentiveIncentiveProvided)
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
func (it *BaasIncentiveIncentiveProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentiveIncentiveProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentiveIncentiveProvided represents a IncentiveProvided event raised by the BaasIncentive contract.
type BaasIncentiveIncentiveProvided struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentiveProvided is a free log retrieval operation binding the contract event 0x32dd062c233d622f3ce5c069f841701219c1c992202aef2d7a75159790bcd070.
//
// Solidity: e IncentiveProvided(account indexed address, amount uint256)
func (_BaasIncentive *BaasIncentiveFilterer) FilterIncentiveProvided(opts *bind.FilterOpts, account []common.Address) (*BaasIncentiveIncentiveProvidedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentive.contract.FilterLogs(opts, "IncentiveProvided", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentiveIncentiveProvidedIterator{contract: _BaasIncentive.contract, event: "IncentiveProvided", logs: logs, sub: sub}, nil
}

// WatchIncentiveProvided is a free log subscription operation binding the contract event 0x32dd062c233d622f3ce5c069f841701219c1c992202aef2d7a75159790bcd070.
//
// Solidity: e IncentiveProvided(account indexed address, amount uint256)
func (_BaasIncentive *BaasIncentiveFilterer) WatchIncentiveProvided(opts *bind.WatchOpts, sink chan<- *BaasIncentiveIncentiveProvided, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentive.contract.WatchLogs(opts, "IncentiveProvided", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentiveIncentiveProvided)
				if err := _BaasIncentive.contract.UnpackLog(event, "IncentiveProvided", log); err != nil {
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

// BaasIncentiveWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BaasIncentive contract.
type BaasIncentiveWithdrawnIterator struct {
	Event *BaasIncentiveWithdrawn // Event containing the contract specifics and raw log

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
func (it *BaasIncentiveWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentiveWithdrawn)
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
		it.Event = new(BaasIncentiveWithdrawn)
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
func (it *BaasIncentiveWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentiveWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentiveWithdrawn represents a Withdrawn event raised by the BaasIncentive contract.
type BaasIncentiveWithdrawn struct {
	Account common.Address
	Amount  *big.Int
	Stage   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: e Withdrawn(account indexed address, amount uint256, stage uint256)
func (_BaasIncentive *BaasIncentiveFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*BaasIncentiveWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentive.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentiveWithdrawnIterator{contract: _BaasIncentive.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: e Withdrawn(account indexed address, amount uint256, stage uint256)
func (_BaasIncentive *BaasIncentiveFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BaasIncentiveWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentive.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentiveWithdrawn)
				if err := _BaasIncentive.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
