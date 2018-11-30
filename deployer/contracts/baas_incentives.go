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
const BaasIncentivesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncentiveProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stage\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"vestingPeriodInBlocks\",\"type\":\"uint256\"},{\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"provideIncentive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesProvided\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingPeriodInBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getIncentive\",\"outputs\":[{\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"stage\",\"type\":\"uint256\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"stage\",\"type\":\"uint256\"}],\"name\":\"canWithdraw\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"alreadyWithdrawn\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasIncentivesBin is the compiled bytecode used for deploying new contracts.
const BaasIncentivesBin = `0x60806040526000600560006101000a81548160ff02191690831515021790555034801561002b57600080fd5b5060405160208061143183398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506112d68061015b6000396000f3006080604052600436106100e6576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100eb5780630a8fed891461017b578063392e53cd146101b25780633ccfd60b146101e15780635742a36e14610210578063715018a61461023b5780638da5cb5b146102525780638f32d59b146102a95780638f70c2d8146102d85780639d76ea581461034f578063ab77e2d3146103a6578063b69ef8a8146103d1578063cb5595bb146103fc578063cd3a2a3c14610461578063e2ab7cf21461048c578063f2fde38b146104ff575b600080fd5b3480156100f757600080fd5b50610100610542565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610140578082015181840152602081019050610125565b50505050905090810190601f16801561016d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018757600080fd5b506101b0600480360381019080803590602001909291908035906020019092919050505061057f565b005b3480156101be57600080fd5b506101c76105b5565b604051808215151515815260200191505060405180910390f35b3480156101ed57600080fd5b506101f66105cc565b604051808215151515815260200191505060405180910390f35b34801561021c57600080fd5b5061022561098b565b6040518082815260200191505060405180910390f35b34801561024757600080fd5b50610250610995565b005b34801561025e57600080fd5b50610267610a67565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102b557600080fd5b506102be610a90565b604051808215151515815260200191505060405180910390f35b3480156102e457600080fd5b50610319600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ae7565b60405180868152602001858152602001848152602001838152602001821515151581526020019550505050505060405180910390f35b34801561035b57600080fd5b50610364610c0a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103b257600080fd5b506103bb610c34565b6040518082815260200191505060405180910390f35b3480156103dd57600080fd5b506103e6610c3e565b6040518082815260200191505060405180910390f35b34801561040857600080fd5b50610447600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d3d565b604051808215151515815260200191505060405180910390f35b34801561046d57600080fd5b50610476610f67565b6040518082815260200191505060405180910390f35b34801561049857600080fd5b506104d7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f71565b6040518084815260200183815260200182151515158152602001935050505060405180910390f35b34801561050b57600080fd5b50610540600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110c2565b005b60606040805190810160405280600c81526020017f494e43454e544956455320200000000000000000000000000000000000000000815250905090565b600560009054906101000a900460ff1615151561059b57600080fd5b806002819055506000600381905550816004819055505050565b6000600560009054906101000a900460ff16905090565b6000806000806000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160009054906101000a900460ff16151561062f57600080fd5b6004600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481151561067d57fe5b049350600454600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003015443038115156106d157fe5b04925060018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040154840301915061073082856110e190919063ffffffff16565b905060008211151561074157600080fd5b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561080657600080fd5b505af115801561081a573d6000803e3d6000fd5b505050506040513d602081101561083057600080fd5b8101908080519060200190929190505050151561084c57600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600401819055506108e881600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015461111f90919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055503373ffffffffffffffffffffffffffffffffffffffff167f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc68284604051808381526020018281526020019250505060405180910390a25050505090565b6000600354905090565b61099d610a90565b15156109a857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000806000806000610af761125b565b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060c060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff161515151581525050905080602001518160400151826060015183608001518460a00151955095509550955095505091939590929450565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600254905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610cfd57600080fd5b505af1158015610d11573d6000803e3d6000fd5b505050506040513d6020811015610d2757600080fd5b8101908080519060200190929190505050905090565b6000610d47610a90565b1515610d5257600080fd5b6002548211151515610d6357600080fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160009054906101000a900460ff16151515610dbf57600080fd5b610dd48260025461114090919063ffffffff16565b600281905550610def8260035461111f90919063ffffffff16565b60038190555060c0604051908101604052808473ffffffffffffffffffffffffffffffffffffffff168152602001838152602001600081526020014381526020016000815260200160011515815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff0219169083151502179055509050508273ffffffffffffffffffffffffffffffffffffffff167f32dd062c233d622f3ce5c069f841701219c1c992202aef2d7a75159790bcd070836040518082815260200191505060405180910390a26001905092915050565b6000600454905090565b6000806000610f7e61125b565b6000851415610f9d5760008060008292508191509350935093506110ba565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060c060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900460ff16151515158152505090506004816020015181151561109357fe5b04935060045460018603028160600151019250848160800151101591508383839350935093505b509250925092565b6110ca610a90565b15156110d557600080fd5b6110de81611161565b50565b60008060008414156110f65760009150611118565b828402905082848281151561110757fe5b0414151561111457600080fd5b8091505b5092915050565b600080828401905083811015151561113657600080fd5b8091505092915050565b60008083831115151561115257600080fd5b82840390508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561119d57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60c060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000815260200160001515815250905600a165627a7a723058202f8a0b9ced129ae8bc7be13543eba8fb42a8ce908a23d75a714fb7c912153a410029`

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

// CanWithdraw is a free data retrieval call binding the contract method 0xe2ab7cf2.
//
// Solidity: function canWithdraw(account address, stage uint256) constant returns(amount uint256, atBlock uint256, alreadyWithdrawn bool)
func (_BaasIncentives *BaasIncentivesCaller) CanWithdraw(opts *bind.CallOpts, account common.Address, stage *big.Int) (struct {
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
	err := _BaasIncentives.contract.Call(opts, out, "canWithdraw", account, stage)
	return *ret, err
}

// CanWithdraw is a free data retrieval call binding the contract method 0xe2ab7cf2.
//
// Solidity: function canWithdraw(account address, stage uint256) constant returns(amount uint256, atBlock uint256, alreadyWithdrawn bool)
func (_BaasIncentives *BaasIncentivesSession) CanWithdraw(account common.Address, stage *big.Int) (struct {
	Amount           *big.Int
	AtBlock          *big.Int
	AlreadyWithdrawn bool
}, error) {
	return _BaasIncentives.Contract.CanWithdraw(&_BaasIncentives.CallOpts, account, stage)
}

// CanWithdraw is a free data retrieval call binding the contract method 0xe2ab7cf2.
//
// Solidity: function canWithdraw(account address, stage uint256) constant returns(amount uint256, atBlock uint256, alreadyWithdrawn bool)
func (_BaasIncentives *BaasIncentivesCallerSession) CanWithdraw(account common.Address, stage *big.Int) (struct {
	Amount           *big.Int
	AtBlock          *big.Int
	AlreadyWithdrawn bool
}, error) {
	return _BaasIncentives.Contract.CanWithdraw(&_BaasIncentives.CallOpts, account, stage)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, stage uint256, isValue bool)
func (_BaasIncentives *BaasIncentivesCaller) GetIncentive(opts *bind.CallOpts, account common.Address) (struct {
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
	err := _BaasIncentives.contract.Call(opts, out, "getIncentive", account)
	return *ret, err
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, stage uint256, isValue bool)
func (_BaasIncentives *BaasIncentivesSession) GetIncentive(account common.Address) (struct {
	InitialAmount *big.Int
	Withdrawn     *big.Int
	AtBlock       *big.Int
	Stage         *big.Int
	IsValue       bool
}, error) {
	return _BaasIncentives.Contract.GetIncentive(&_BaasIncentives.CallOpts, account)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, stage uint256, isValue bool)
func (_BaasIncentives *BaasIncentivesCallerSession) GetIncentive(account common.Address) (struct {
	InitialAmount *big.Int
	Withdrawn     *big.Int
	AtBlock       *big.Int
	Stage         *big.Int
	IsValue       bool
}, error) {
	return _BaasIncentives.Contract.GetIncentive(&_BaasIncentives.CallOpts, account)
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

// IncentivesProvided is a free data retrieval call binding the contract method 0x5742a36e.
//
// Solidity: function incentivesProvided() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) IncentivesProvided(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "incentivesProvided")
	return *ret0, err
}

// IncentivesProvided is a free data retrieval call binding the contract method 0x5742a36e.
//
// Solidity: function incentivesProvided() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) IncentivesProvided() (*big.Int, error) {
	return _BaasIncentives.Contract.IncentivesProvided(&_BaasIncentives.CallOpts)
}

// IncentivesProvided is a free data retrieval call binding the contract method 0x5742a36e.
//
// Solidity: function incentivesProvided() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) IncentivesProvided() (*big.Int, error) {
	return _BaasIncentives.Contract.IncentivesProvided(&_BaasIncentives.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasIncentives *BaasIncentivesCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasIncentives *BaasIncentivesSession) IsInitialized() (bool, error) {
	return _BaasIncentives.Contract.IsInitialized(&_BaasIncentives.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasIncentives *BaasIncentivesCallerSession) IsInitialized() (bool, error) {
	return _BaasIncentives.Contract.IsInitialized(&_BaasIncentives.CallOpts)
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

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasIncentives *BaasIncentivesCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasIncentives *BaasIncentivesSession) Name() (string, error) {
	return _BaasIncentives.Contract.Name(&_BaasIncentives.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasIncentives *BaasIncentivesCallerSession) Name() (string, error) {
	return _BaasIncentives.Contract.Name(&_BaasIncentives.CallOpts)
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

// VestingPeriodInBlocks is a free data retrieval call binding the contract method 0xcd3a2a3c.
//
// Solidity: function vestingPeriodInBlocks() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) VestingPeriodInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "vestingPeriodInBlocks")
	return *ret0, err
}

// VestingPeriodInBlocks is a free data retrieval call binding the contract method 0xcd3a2a3c.
//
// Solidity: function vestingPeriodInBlocks() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) VestingPeriodInBlocks() (*big.Int, error) {
	return _BaasIncentives.Contract.VestingPeriodInBlocks(&_BaasIncentives.CallOpts)
}

// VestingPeriodInBlocks is a free data retrieval call binding the contract method 0xcd3a2a3c.
//
// Solidity: function vestingPeriodInBlocks() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) VestingPeriodInBlocks() (*big.Int, error) {
	return _BaasIncentives.Contract.VestingPeriodInBlocks(&_BaasIncentives.CallOpts)
}

// ProvideIncentive is a paid mutator transaction binding the contract method 0xcb5595bb.
//
// Solidity: function provideIncentive(account address, amount uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) ProvideIncentive(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "provideIncentive", account, amount)
}

// ProvideIncentive is a paid mutator transaction binding the contract method 0xcb5595bb.
//
// Solidity: function provideIncentive(account address, amount uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) ProvideIncentive(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.ProvideIncentive(&_BaasIncentives.TransactOpts, account, amount)
}

// ProvideIncentive is a paid mutator transaction binding the contract method 0xcb5595bb.
//
// Solidity: function provideIncentive(account address, amount uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) ProvideIncentive(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.ProvideIncentive(&_BaasIncentives.TransactOpts, account, amount)
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

// Setup is a paid mutator transaction binding the contract method 0x0a8fed89.
//
// Solidity: function setup(vestingPeriodInBlocks uint256, supply uint256) returns()
func (_BaasIncentives *BaasIncentivesTransactor) Setup(opts *bind.TransactOpts, vestingPeriodInBlocks *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "setup", vestingPeriodInBlocks, supply)
}

// Setup is a paid mutator transaction binding the contract method 0x0a8fed89.
//
// Solidity: function setup(vestingPeriodInBlocks uint256, supply uint256) returns()
func (_BaasIncentives *BaasIncentivesSession) Setup(vestingPeriodInBlocks *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Setup(&_BaasIncentives.TransactOpts, vestingPeriodInBlocks, supply)
}

// Setup is a paid mutator transaction binding the contract method 0x0a8fed89.
//
// Solidity: function setup(vestingPeriodInBlocks uint256, supply uint256) returns()
func (_BaasIncentives *BaasIncentivesTransactorSession) Setup(vestingPeriodInBlocks *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Setup(&_BaasIncentives.TransactOpts, vestingPeriodInBlocks, supply)
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

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Withdraw() (*types.Transaction, error) {
	return _BaasIncentives.Contract.Withdraw(&_BaasIncentives.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BaasIncentives.Contract.Withdraw(&_BaasIncentives.TransactOpts)
}

// BaasIncentivesIncentiveProvidedIterator is returned from FilterIncentiveProvided and is used to iterate over the raw logs and unpacked data for IncentiveProvided events raised by the BaasIncentives contract.
type BaasIncentivesIncentiveProvidedIterator struct {
	Event *BaasIncentivesIncentiveProvided // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesIncentiveProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesIncentiveProvided)
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
		it.Event = new(BaasIncentivesIncentiveProvided)
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
func (it *BaasIncentivesIncentiveProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesIncentiveProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesIncentiveProvided represents a IncentiveProvided event raised by the BaasIncentives contract.
type BaasIncentivesIncentiveProvided struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentiveProvided is a free log retrieval operation binding the contract event 0x32dd062c233d622f3ce5c069f841701219c1c992202aef2d7a75159790bcd070.
//
// Solidity: e IncentiveProvided(account indexed address, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterIncentiveProvided(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesIncentiveProvidedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "IncentiveProvided", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesIncentiveProvidedIterator{contract: _BaasIncentives.contract, event: "IncentiveProvided", logs: logs, sub: sub}, nil
}

// WatchIncentiveProvided is a free log subscription operation binding the contract event 0x32dd062c233d622f3ce5c069f841701219c1c992202aef2d7a75159790bcd070.
//
// Solidity: e IncentiveProvided(account indexed address, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchIncentiveProvided(opts *bind.WatchOpts, sink chan<- *BaasIncentivesIncentiveProvided, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "IncentiveProvided", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesIncentiveProvided)
				if err := _BaasIncentives.contract.UnpackLog(event, "IncentiveProvided", log); err != nil {
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

// BaasIncentivesWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BaasIncentives contract.
type BaasIncentivesWithdrawnIterator struct {
	Event *BaasIncentivesWithdrawn // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesWithdrawn)
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
		it.Event = new(BaasIncentivesWithdrawn)
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
func (it *BaasIncentivesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesWithdrawn represents a Withdrawn event raised by the BaasIncentives contract.
type BaasIncentivesWithdrawn struct {
	Account common.Address
	Amount  *big.Int
	Stage   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: e Withdrawn(account indexed address, amount uint256, stage uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesWithdrawnIterator{contract: _BaasIncentives.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: e Withdrawn(account indexed address, amount uint256, stage uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BaasIncentivesWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesWithdrawn)
				if err := _BaasIncentives.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
