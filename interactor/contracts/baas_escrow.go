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
const BaasEscrowABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CapitalRaised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vestingStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingEndBlock\",\"type\":\"uint256\"}],\"name\":\"SetupCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenPrice\",\"type\":\"uint256\"}],\"name\":\"TokenDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"vestingPeriod\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiseCapital\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint256\"}],\"name\":\"provideToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingStartBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vestingEndBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blocknumber\",\"type\":\"uint256\"}],\"name\":\"canRaise\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BaasEscrowBin is the compiled bytecode used for deploying new contracts.
const BaasEscrowBin = `0x60806040526000600760006101000a81548160ff02191690831515021790555034801561002b57600080fd5b5060405160208061125c83398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506111018061015b6000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100e057806314e3a51014610170578063392e53cd146101b55780634313b9e5146101e457806356d17abf14610229578063715018a6146102545780637313ee5a1461026b5780638b854c72146102965780638da5cb5b146103055780638f32d59b1461035c5780639d76ea581461038b5780639edf4e64146103e2578063b69ef8a81461040d578063f2fde38b14610438578063fbd207ed1461047b575b600080fd5b3480156100ec57600080fd5b506100f56104ca565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013557808201518184015260208101905061011a565b50505050905090810190601f1680156101625780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017c57600080fd5b5061019b60048036038101908080359060200190929190505050610507565b604051808215151515815260200191505060405180910390f35b3480156101c157600080fd5b506101ca61051b565b604051808215151515815260200191505060405180910390f35b3480156101f057600080fd5b5061020f60048036038101908080359060200190929190505050610532565b604051808215151515815260200191505060405180910390f35b34801561023557600080fd5b5061023e6105d6565b6040518082815260200191505060405180910390f35b34801561026057600080fd5b506102696105f4565b005b34801561027757600080fd5b506102806106c6565b6040518082815260200191505060405180910390f35b3480156102a257600080fd5b506102eb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291905050506106d0565b604051808215151515815260200191505060405180910390f35b34801561031157600080fd5b5061031a610851565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561036857600080fd5b5061037161087a565b604051808215151515815260200191505060405180910390f35b34801561039757600080fd5b506103a06108d1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103ee57600080fd5b506103f76108fb565b6040518082815260200191505060405180910390f35b34801561041957600080fd5b50610422610905565b6040518082815260200191505060405180910390f35b34801561044457600080fd5b50610479600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a04565b005b34801561048757600080fd5b506104b06004803603810190808035906020019092919080359060200190929190505050610a23565b604051808215151515815260200191505060405180910390f35b60606040805190810160405280600681526020017f455343524f570000000000000000000000000000000000000000000000000000815250905090565b6000816105126105d6565b11159050919050565b6000600760009054906101000a900460ff16905090565b600061053c61087a565b151561054757600080fd5b43600281905550816003819055506001600760006101000a81548160ff0219169083151502179055507f05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd0996002546003546105ae600354600254610fba90919063ffffffff16565b60405180848152602001838152602001828152602001935050505060405180910390a1919050565b60006105ef600354600254610fba90919063ffffffff16565b905090565b6105fc61087a565b151561060757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600354905090565b60006106da61087a565b15156106e557600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156107aa57600080fd5b505af11580156107be573d6000803e3d6000fd5b505050506040513d60208110156107d457600080fd5b810190808051906020019092919050505015156107f057600080fd5b8373ffffffffffffffffffffffffffffffffffffffff167f80438a1ee44a1e699790add4b8fddef1ff8532d2fdb33aa686c90b79bb2684788484604051808381526020018281526020019250505060405180910390a2600190509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600254905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156109c457600080fd5b505af11580156109d8573d6000803e3d6000fd5b505050506040513d60208110156109ee57600080fd5b8101908080519060200190929190505050905090565b610a0c61087a565b1515610a1757600080fd5b610a2081610fdb565b50565b600080600080600060606000806000610a3a61087a565b1515610a4557600080fd5b610a4e43610507565b1515610a5957600080fd5b610a61610905565b8b11151515610a6f57600080fd5b600560008b815260200190815260200160002060020160009054906101000a900460ff16151515610a9f57600080fd5b6060604051908101604052808b81526020018c815260200160011515815250600560008c8152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff021916908315150217905550905050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a654cfab6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401608060405180830381600087803b158015610b8e57600080fd5b505af1158015610ba2573d6000803e3d6000fd5b505050506040513d6080811015610bb857600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050809850819950829a50839b5050505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637af70c1f6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015610c7d57600080fd5b505af1158015610c91573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610cbb57600080fd5b810190808051640100000000811115610cd357600080fd5b82810190506020810184811115610ce957600080fd5b8151856020820283011164010000000082111715610d0657600080fd5b5050929190505050935083519250600091505b82821015610f70578382815181101515610d2f57fe5b9060200190602002015190508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610da057508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80610dd657508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80610e0c57508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610e1657610f63565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610ed357600080fd5b505af1158015610ee7573d6000803e3d6000fd5b505050506040513d6020811015610efd57600080fd5b8101908080519060200190929190505050600660008c815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b8180600101925050610d19565b897fbc65650fcb1dd5a9e3c92e066533c3425c66a51f1f60fbfe11ce352452c9b08d8c6040518082815260200191505060405180910390a260019850505050505050505092915050565b6000808284019050838110151515610fd157600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561101757600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820eb06c6ed0f92b186668db88aa0e7b71d265590f921e9a36d7a1cf1f5fefa7ed80029`

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

// CanRaise is a free data retrieval call binding the contract method 0x14e3a510.
//
// Solidity: function canRaise(blocknumber uint256) constant returns(bool)
func (_BaasEscrow *BaasEscrowCaller) CanRaise(opts *bind.CallOpts, blocknumber *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "canRaise", blocknumber)
	return *ret0, err
}

// CanRaise is a free data retrieval call binding the contract method 0x14e3a510.
//
// Solidity: function canRaise(blocknumber uint256) constant returns(bool)
func (_BaasEscrow *BaasEscrowSession) CanRaise(blocknumber *big.Int) (bool, error) {
	return _BaasEscrow.Contract.CanRaise(&_BaasEscrow.CallOpts, blocknumber)
}

// CanRaise is a free data retrieval call binding the contract method 0x14e3a510.
//
// Solidity: function canRaise(blocknumber uint256) constant returns(bool)
func (_BaasEscrow *BaasEscrowCallerSession) CanRaise(blocknumber *big.Int) (bool, error) {
	return _BaasEscrow.Contract.CanRaise(&_BaasEscrow.CallOpts, blocknumber)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasEscrow *BaasEscrowCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasEscrow *BaasEscrowSession) IsInitialized() (bool, error) {
	return _BaasEscrow.Contract.IsInitialized(&_BaasEscrow.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasEscrow *BaasEscrowCallerSession) IsInitialized() (bool, error) {
	return _BaasEscrow.Contract.IsInitialized(&_BaasEscrow.CallOpts)
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

// VestingEndBlock is a free data retrieval call binding the contract method 0x56d17abf.
//
// Solidity: function vestingEndBlock() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) VestingEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "vestingEndBlock")
	return *ret0, err
}

// VestingEndBlock is a free data retrieval call binding the contract method 0x56d17abf.
//
// Solidity: function vestingEndBlock() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) VestingEndBlock() (*big.Int, error) {
	return _BaasEscrow.Contract.VestingEndBlock(&_BaasEscrow.CallOpts)
}

// VestingEndBlock is a free data retrieval call binding the contract method 0x56d17abf.
//
// Solidity: function vestingEndBlock() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) VestingEndBlock() (*big.Int, error) {
	return _BaasEscrow.Contract.VestingEndBlock(&_BaasEscrow.CallOpts)
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) VestingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "vestingPeriod")
	return *ret0, err
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) VestingPeriod() (*big.Int, error) {
	return _BaasEscrow.Contract.VestingPeriod(&_BaasEscrow.CallOpts)
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) VestingPeriod() (*big.Int, error) {
	return _BaasEscrow.Contract.VestingPeriod(&_BaasEscrow.CallOpts)
}

// VestingStartBlock is a free data retrieval call binding the contract method 0x9edf4e64.
//
// Solidity: function vestingStartBlock() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) VestingStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "vestingStartBlock")
	return *ret0, err
}

// VestingStartBlock is a free data retrieval call binding the contract method 0x9edf4e64.
//
// Solidity: function vestingStartBlock() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) VestingStartBlock() (*big.Int, error) {
	return _BaasEscrow.Contract.VestingStartBlock(&_BaasEscrow.CallOpts)
}

// VestingStartBlock is a free data retrieval call binding the contract method 0x9edf4e64.
//
// Solidity: function vestingStartBlock() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) VestingStartBlock() (*big.Int, error) {
	return _BaasEscrow.Contract.VestingStartBlock(&_BaasEscrow.CallOpts)
}

// ProvideToken is a paid mutator transaction binding the contract method 0x8b854c72.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) ProvideToken(opts *bind.TransactOpts, account common.Address, amount *big.Int, conversionRate *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "provideToken", account, amount, conversionRate)
}

// ProvideToken is a paid mutator transaction binding the contract method 0x8b854c72.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint256) returns(bool)
func (_BaasEscrow *BaasEscrowSession) ProvideToken(account common.Address, amount *big.Int, conversionRate *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.ProvideToken(&_BaasEscrow.TransactOpts, account, amount, conversionRate)
}

// ProvideToken is a paid mutator transaction binding the contract method 0x8b854c72.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) ProvideToken(account common.Address, amount *big.Int, conversionRate *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.ProvideToken(&_BaasEscrow.TransactOpts, account, amount, conversionRate)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xfbd207ed.
//
// Solidity: function raiseCapital(amount uint256, id uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) RaiseCapital(opts *bind.TransactOpts, amount *big.Int, id *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "raiseCapital", amount, id)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xfbd207ed.
//
// Solidity: function raiseCapital(amount uint256, id uint256) returns(bool)
func (_BaasEscrow *BaasEscrowSession) RaiseCapital(amount *big.Int, id *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.RaiseCapital(&_BaasEscrow.TransactOpts, amount, id)
}

// RaiseCapital is a paid mutator transaction binding the contract method 0xfbd207ed.
//
// Solidity: function raiseCapital(amount uint256, id uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) RaiseCapital(amount *big.Int, id *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.RaiseCapital(&_BaasEscrow.TransactOpts, amount, id)
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

// Setup is a paid mutator transaction binding the contract method 0x4313b9e5.
//
// Solidity: function setup(vestingPeriod uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) Setup(opts *bind.TransactOpts, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "setup", vestingPeriod)
}

// Setup is a paid mutator transaction binding the contract method 0x4313b9e5.
//
// Solidity: function setup(vestingPeriod uint256) returns(bool)
func (_BaasEscrow *BaasEscrowSession) Setup(vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.Setup(&_BaasEscrow.TransactOpts, vestingPeriod)
}

// Setup is a paid mutator transaction binding the contract method 0x4313b9e5.
//
// Solidity: function setup(vestingPeriod uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) Setup(vestingPeriod *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.Setup(&_BaasEscrow.TransactOpts, vestingPeriod)
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
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCapitalRaised is a free log retrieval operation binding the contract event 0xbc65650fcb1dd5a9e3c92e066533c3425c66a51f1f60fbfe11ce352452c9b08d.
//
// Solidity: e CapitalRaised(id indexed uint256, amount uint256)
func (_BaasEscrow *BaasEscrowFilterer) FilterCapitalRaised(opts *bind.FilterOpts, id []*big.Int) (*BaasEscrowCapitalRaisedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "CapitalRaised", idRule)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowCapitalRaisedIterator{contract: _BaasEscrow.contract, event: "CapitalRaised", logs: logs, sub: sub}, nil
}

// WatchCapitalRaised is a free log subscription operation binding the contract event 0xbc65650fcb1dd5a9e3c92e066533c3425c66a51f1f60fbfe11ce352452c9b08d.
//
// Solidity: e CapitalRaised(id indexed uint256, amount uint256)
func (_BaasEscrow *BaasEscrowFilterer) WatchCapitalRaised(opts *bind.WatchOpts, sink chan<- *BaasEscrowCapitalRaised, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "CapitalRaised", idRule)
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

// BaasEscrowSetupCompletedIterator is returned from FilterSetupCompleted and is used to iterate over the raw logs and unpacked data for SetupCompleted events raised by the BaasEscrow contract.
type BaasEscrowSetupCompletedIterator struct {
	Event *BaasEscrowSetupCompleted // Event containing the contract specifics and raw log

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
func (it *BaasEscrowSetupCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasEscrowSetupCompleted)
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
		it.Event = new(BaasEscrowSetupCompleted)
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
func (it *BaasEscrowSetupCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasEscrowSetupCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasEscrowSetupCompleted represents a SetupCompleted event raised by the BaasEscrow contract.
type BaasEscrowSetupCompleted struct {
	VestingStartBlock *big.Int
	VestingPeriod     *big.Int
	VestingEndBlock   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetupCompleted is a free log retrieval operation binding the contract event 0x05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd099.
//
// Solidity: e SetupCompleted(vestingStartBlock uint256, vestingPeriod uint256, vestingEndBlock uint256)
func (_BaasEscrow *BaasEscrowFilterer) FilterSetupCompleted(opts *bind.FilterOpts) (*BaasEscrowSetupCompletedIterator, error) {

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return &BaasEscrowSetupCompletedIterator{contract: _BaasEscrow.contract, event: "SetupCompleted", logs: logs, sub: sub}, nil
}

// WatchSetupCompleted is a free log subscription operation binding the contract event 0x05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd099.
//
// Solidity: e SetupCompleted(vestingStartBlock uint256, vestingPeriod uint256, vestingEndBlock uint256)
func (_BaasEscrow *BaasEscrowFilterer) WatchSetupCompleted(opts *bind.WatchOpts, sink chan<- *BaasEscrowSetupCompleted) (event.Subscription, error) {

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasEscrowSetupCompleted)
				if err := _BaasEscrow.contract.UnpackLog(event, "SetupCompleted", log); err != nil {
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

// BaasEscrowTokenDeliveredIterator is returned from FilterTokenDelivered and is used to iterate over the raw logs and unpacked data for TokenDelivered events raised by the BaasEscrow contract.
type BaasEscrowTokenDeliveredIterator struct {
	Event *BaasEscrowTokenDelivered // Event containing the contract specifics and raw log

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
func (it *BaasEscrowTokenDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasEscrowTokenDelivered)
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
		it.Event = new(BaasEscrowTokenDelivered)
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
func (it *BaasEscrowTokenDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasEscrowTokenDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasEscrowTokenDelivered represents a TokenDelivered event raised by the BaasEscrow contract.
type BaasEscrowTokenDelivered struct {
	To         common.Address
	Amount     *big.Int
	TokenPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenDelivered is a free log retrieval operation binding the contract event 0x80438a1ee44a1e699790add4b8fddef1ff8532d2fdb33aa686c90b79bb268478.
//
// Solidity: e TokenDelivered(to indexed address, amount uint256, tokenPrice uint256)
func (_BaasEscrow *BaasEscrowFilterer) FilterTokenDelivered(opts *bind.FilterOpts, to []common.Address) (*BaasEscrowTokenDeliveredIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "TokenDelivered", toRule)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowTokenDeliveredIterator{contract: _BaasEscrow.contract, event: "TokenDelivered", logs: logs, sub: sub}, nil
}

// WatchTokenDelivered is a free log subscription operation binding the contract event 0x80438a1ee44a1e699790add4b8fddef1ff8532d2fdb33aa686c90b79bb268478.
//
// Solidity: e TokenDelivered(to indexed address, amount uint256, tokenPrice uint256)
func (_BaasEscrow *BaasEscrowFilterer) WatchTokenDelivered(opts *bind.WatchOpts, sink chan<- *BaasEscrowTokenDelivered, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "TokenDelivered", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasEscrowTokenDelivered)
				if err := _BaasEscrow.contract.UnpackLog(event, "TokenDelivered", log); err != nil {
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
