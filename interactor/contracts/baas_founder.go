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
const BaasFounderABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FounderWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vestingStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingEndBlock\",\"type\":\"uint256\"}],\"name\":\"SetupCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"vestingPeriod\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingEndBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"canWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasFounderBin is the compiled bytecode used for deploying new contracts.
const BaasFounderBin = `0x60806040526000600360006101000a81548160ff02191690831515021790555034801561002b57600080fd5b50604051602080610c9d83398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610b428061015b6000396000f3006080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100d5578063392e53cd146101655780634313b9e51461019457806356d17abf146101d9578063715018a6146102045780637313ee5a1461021b5780638da5cb5b146102465780638f32d59b1461029d5780639d76ea58146102cc5780639edf4e6414610323578063b69ef8a81461034e578063f2fde38b14610379578063f3fef3a3146103bc578063fbe85f0614610421575b600080fd5b3480156100e157600080fd5b506100ea610466565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561012a57808201518184015260208101905061010f565b50505050905090810190601f1680156101575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017157600080fd5b5061017a6104a3565b604051808215151515815260200191505060405180910390f35b3480156101a057600080fd5b506101bf600480360381019080803590602001909291905050506104ba565b604051808215151515815260200191505060405180910390f35b3480156101e557600080fd5b506101ee61055e565b6040518082815260200191505060405180910390f35b34801561021057600080fd5b5061021961057c565b005b34801561022757600080fd5b5061023061064e565b6040518082815260200191505060405180910390f35b34801561025257600080fd5b5061025b610658565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102a957600080fd5b506102b2610681565b604051808215151515815260200191505060405180910390f35b3480156102d857600080fd5b506102e16106d8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561032f57600080fd5b50610338610702565b6040518082815260200191505060405180910390f35b34801561035a57600080fd5b5061036361070c565b6040518082815260200191505060405180910390f35b34801561038557600080fd5b506103ba600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061080b565b005b3480156103c857600080fd5b50610407600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061082a565b604051808215151515815260200191505060405180910390f35b34801561042d57600080fd5b5061044c600480360381019080803590602001909291905050506109e7565b604051808215151515815260200191505060405180910390f35b60606040805190810160405280600781526020017f464f554e44455200000000000000000000000000000000000000000000000000815250905090565b6000600360009054906101000a900460ff16905090565b60006104c4610681565b15156104cf57600080fd5b43600181905550816002819055506001600360006101000a81548160ff0219169083151502179055507f05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd0996001546002546105366002546001546109fb90919063ffffffff16565b60405180848152602001838152602001828152602001935050505060405180910390a1919050565b60006105776002546001546109fb90919063ffffffff16565b905090565b610584610681565b151561058f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600254905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600154905090565b6000600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156107cb57600080fd5b505af11580156107df573d6000803e3d6000fd5b505050506040513d60208110156107f557600080fd5b8101908080519060200190929190505050905090565b610813610681565b151561081e57600080fd5b61082781610a1c565b50565b6000610834610681565b151561083f57600080fd5b600360009054906101000a900460ff16151561085a57600080fd5b610863436109e7565b151561086e57600080fd5b61087661070c565b821115151561088457600080fd5b600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561094957600080fd5b505af115801561095d573d6000803e3d6000fd5b505050506040513d602081101561097357600080fd5b8101908080519060200190929190505050151561098f57600080fd5b8273ffffffffffffffffffffffffffffffffffffffff167f60ca83f5e5baea21ed476b19b1333e0823cc130945297b532873df0ba8b970c9836040518082815260200191505060405180910390a26001905092915050565b6000816109f261055e565b11159050919050565b6000808284019050838110151515610a1257600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610a5857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058207b8709f70b97b58c3e8645356548823f477f4f35aa084f46a9a5bed8b6c6480f0029`

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

// CanWithdraw is a free data retrieval call binding the contract method 0xfbe85f06.
//
// Solidity: function canWithdraw(blocknumber uint256) constant returns(bool)
func (_BaasFounder *BaasFounderCaller) CanWithdraw(opts *bind.CallOpts, blocknumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "canWithdraw", blocknumber)
	return *ret0, err
}

// CanWithdraw is a free data retrieval call binding the contract method 0xfbe85f06.
//
// Solidity: function canWithdraw(blocknumber uint256) constant returns(bool)
func (_BaasFounder *BaasFounderSession) CanWithdraw(blocknumber *big.Int) (bool, error) {
	return _BaasFounder.Contract.CanWithdraw(&_BaasFounder.CallOpts, blocknumber)
}

// CanWithdraw is a free data retrieval call binding the contract method 0xfbe85f06.
//
// Solidity: function canWithdraw(blocknumber uint256) constant returns(bool)
func (_BaasFounder *BaasFounderCallerSession) CanWithdraw(blocknumber *big.Int) (bool, error) {
	return _BaasFounder.Contract.CanWithdraw(&_BaasFounder.CallOpts, blocknumber)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasFounder *BaasFounderCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasFounder *BaasFounderSession) IsInitialized() (bool, error) {
	return _BaasFounder.Contract.IsInitialized(&_BaasFounder.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasFounder *BaasFounderCallerSession) IsInitialized() (bool, error) {
	return _BaasFounder.Contract.IsInitialized(&_BaasFounder.CallOpts)
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

// VestingEndBlock is a free data retrieval call binding the contract method 0x56d17abf.
//
// Solidity: function vestingEndBlock() constant returns(uint256)
func (_BaasFounder *BaasFounderCaller) VestingEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "vestingEndBlock")
	return *ret0, err
}

// VestingEndBlock is a free data retrieval call binding the contract method 0x56d17abf.
//
// Solidity: function vestingEndBlock() constant returns(uint256)
func (_BaasFounder *BaasFounderSession) VestingEndBlock() (*big.Int, error) {
	return _BaasFounder.Contract.VestingEndBlock(&_BaasFounder.CallOpts)
}

// VestingEndBlock is a free data retrieval call binding the contract method 0x56d17abf.
//
// Solidity: function vestingEndBlock() constant returns(uint256)
func (_BaasFounder *BaasFounderCallerSession) VestingEndBlock() (*big.Int, error) {
	return _BaasFounder.Contract.VestingEndBlock(&_BaasFounder.CallOpts)
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() constant returns(uint256)
func (_BaasFounder *BaasFounderCaller) VestingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "vestingPeriod")
	return *ret0, err
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() constant returns(uint256)
func (_BaasFounder *BaasFounderSession) VestingPeriod() (*big.Int, error) {
	return _BaasFounder.Contract.VestingPeriod(&_BaasFounder.CallOpts)
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() constant returns(uint256)
func (_BaasFounder *BaasFounderCallerSession) VestingPeriod() (*big.Int, error) {
	return _BaasFounder.Contract.VestingPeriod(&_BaasFounder.CallOpts)
}

// VestingStartBlock is a free data retrieval call binding the contract method 0x9edf4e64.
//
// Solidity: function vestingStartBlock() constant returns(uint256)
func (_BaasFounder *BaasFounderCaller) VestingStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasFounder.contract.Call(opts, out, "vestingStartBlock")
	return *ret0, err
}

// VestingStartBlock is a free data retrieval call binding the contract method 0x9edf4e64.
//
// Solidity: function vestingStartBlock() constant returns(uint256)
func (_BaasFounder *BaasFounderSession) VestingStartBlock() (*big.Int, error) {
	return _BaasFounder.Contract.VestingStartBlock(&_BaasFounder.CallOpts)
}

// VestingStartBlock is a free data retrieval call binding the contract method 0x9edf4e64.
//
// Solidity: function vestingStartBlock() constant returns(uint256)
func (_BaasFounder *BaasFounderCallerSession) VestingStartBlock() (*big.Int, error) {
	return _BaasFounder.Contract.VestingStartBlock(&_BaasFounder.CallOpts)
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

// Setup is a paid mutator transaction binding the contract method 0x4313b9e5.
//
// Solidity: function setup(vestingPeriod uint256) returns(bool)
func (_BaasFounder *BaasFounderTransactor) Setup(opts *bind.TransactOpts, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "setup", vestingPeriod)
}

// Setup is a paid mutator transaction binding the contract method 0x4313b9e5.
//
// Solidity: function setup(vestingPeriod uint256) returns(bool)
func (_BaasFounder *BaasFounderSession) Setup(vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Setup(&_BaasFounder.TransactOpts, vestingPeriod)
}

// Setup is a paid mutator transaction binding the contract method 0x4313b9e5.
//
// Solidity: function setup(vestingPeriod uint256) returns(bool)
func (_BaasFounder *BaasFounderTransactorSession) Setup(vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Setup(&_BaasFounder.TransactOpts, vestingPeriod)
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

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns(bool)
func (_BaasFounder *BaasFounderTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasFounder.contract.Transact(opts, "withdraw", receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns(bool)
func (_BaasFounder *BaasFounderSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Withdraw(&_BaasFounder.TransactOpts, receiver, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, amount uint256) returns(bool)
func (_BaasFounder *BaasFounderTransactorSession) Withdraw(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasFounder.Contract.Withdraw(&_BaasFounder.TransactOpts, receiver, amount)
}

// BaasFounderFounderWithdrawIterator is returned from FilterFounderWithdraw and is used to iterate over the raw logs and unpacked data for FounderWithdraw events raised by the BaasFounder contract.
type BaasFounderFounderWithdrawIterator struct {
	Event *BaasFounderFounderWithdraw // Event containing the contract specifics and raw log

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
func (it *BaasFounderFounderWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasFounderFounderWithdraw)
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
		it.Event = new(BaasFounderFounderWithdraw)
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
func (it *BaasFounderFounderWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasFounderFounderWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasFounderFounderWithdraw represents a FounderWithdraw event raised by the BaasFounder contract.
type BaasFounderFounderWithdraw struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFounderWithdraw is a free log retrieval operation binding the contract event 0x60ca83f5e5baea21ed476b19b1333e0823cc130945297b532873df0ba8b970c9.
//
// Solidity: e FounderWithdraw(receiver indexed address, amount uint256)
func (_BaasFounder *BaasFounderFilterer) FilterFounderWithdraw(opts *bind.FilterOpts, receiver []common.Address) (*BaasFounderFounderWithdrawIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaasFounder.contract.FilterLogs(opts, "FounderWithdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BaasFounderFounderWithdrawIterator{contract: _BaasFounder.contract, event: "FounderWithdraw", logs: logs, sub: sub}, nil
}

// WatchFounderWithdraw is a free log subscription operation binding the contract event 0x60ca83f5e5baea21ed476b19b1333e0823cc130945297b532873df0ba8b970c9.
//
// Solidity: e FounderWithdraw(receiver indexed address, amount uint256)
func (_BaasFounder *BaasFounderFilterer) WatchFounderWithdraw(opts *bind.WatchOpts, sink chan<- *BaasFounderFounderWithdraw, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaasFounder.contract.WatchLogs(opts, "FounderWithdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasFounderFounderWithdraw)
				if err := _BaasFounder.contract.UnpackLog(event, "FounderWithdraw", log); err != nil {
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

// BaasFounderSetupCompletedIterator is returned from FilterSetupCompleted and is used to iterate over the raw logs and unpacked data for SetupCompleted events raised by the BaasFounder contract.
type BaasFounderSetupCompletedIterator struct {
	Event *BaasFounderSetupCompleted // Event containing the contract specifics and raw log

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
func (it *BaasFounderSetupCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasFounderSetupCompleted)
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
		it.Event = new(BaasFounderSetupCompleted)
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
func (it *BaasFounderSetupCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasFounderSetupCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasFounderSetupCompleted represents a SetupCompleted event raised by the BaasFounder contract.
type BaasFounderSetupCompleted struct {
	VestingStartBlock *big.Int
	VestingPeriod     *big.Int
	VestingEndBlock   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetupCompleted is a free log retrieval operation binding the contract event 0x05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd099.
//
// Solidity: e SetupCompleted(vestingStartBlock uint256, vestingPeriod uint256, vestingEndBlock uint256)
func (_BaasFounder *BaasFounderFilterer) FilterSetupCompleted(opts *bind.FilterOpts) (*BaasFounderSetupCompletedIterator, error) {

	logs, sub, err := _BaasFounder.contract.FilterLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return &BaasFounderSetupCompletedIterator{contract: _BaasFounder.contract, event: "SetupCompleted", logs: logs, sub: sub}, nil
}

// WatchSetupCompleted is a free log subscription operation binding the contract event 0x05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd099.
//
// Solidity: e SetupCompleted(vestingStartBlock uint256, vestingPeriod uint256, vestingEndBlock uint256)
func (_BaasFounder *BaasFounderFilterer) WatchSetupCompleted(opts *bind.WatchOpts, sink chan<- *BaasFounderSetupCompleted) (event.Subscription, error) {

	logs, sub, err := _BaasFounder.contract.WatchLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasFounderSetupCompleted)
				if err := _BaasFounder.contract.UnpackLog(event, "SetupCompleted", log); err != nil {
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
