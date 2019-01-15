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
const BaasIncentivesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncentiveIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tokenNotDelivered\",\"type\":\"uint256\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesProvided\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentives\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getIncentive\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InitialSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasIncentivesBin is the compiled bytecode used for deploying new contracts.
const BaasIncentivesBin = `0x60806040526000600560006101000a81548160ff02191690831515021790555034801561002b57600080fd5b506040516020806112118339810180604052602081101561004b57600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506a084595161401484a00000060038190555060006004819055505061108e806101836000396000f3fe6080604052600436106100d0576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100d55780630d5df7ba146101655780631eb25d13146101d1578063392e53cd146101fc5780635742a36e1461022b578063715018a6146102565780638da5cb5b1461026d5780638f32d59b146102c45780638f70c2d8146102f35780639d76ea581461036a578063ab77e2d3146103c1578063b69ef8a8146103ec578063ea3f068d14610417578063f2fde38b14610494575b600080fd5b3480156100e157600080fd5b506100ea6104e5565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561012a57808201518184015260208101905061010f565b50505050905090810190601f1680156101575780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017157600080fd5b5061017a610522565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101bd5780820151818401526020810190506101a2565b505050509050019250505060405180910390f35b3480156101dd57600080fd5b506101e66105b0565b6040518082815260200191505060405180910390f35b34801561020857600080fd5b506102116105c3565b604051808215151515815260200191505060405180910390f35b34801561023757600080fd5b506102406105da565b6040518082815260200191505060405180910390f35b34801561026257600080fd5b5061026b6105e4565b005b34801561027957600080fd5b506102826106b6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102d057600080fd5b506102d96106df565b604051808215151515815260200191505060405180910390f35b3480156102ff57600080fd5b506103426004803603602081101561031657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610736565b6040518084815260200183815260200182151515158152602001935050505060405180910390f35b34801561037657600080fd5b5061037f610832565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103cd57600080fd5b506103d661085c565b6040518082815260200191505060405180910390f35b3480156103f857600080fd5b50610401610866565b6040518082815260200191505060405180910390f35b34801561042357600080fd5b5061047a6004803603606081101561043a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610963565b604051808215151515815260200191505060405180910390f35b3480156104a057600080fd5b506104e3600480360360208110156104b757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ec2565b005b60606040805190810160405280600a81526020017f494e43454e544956455300000000000000000000000000000000000000000000815250905090565b606060028054806020026020016040519081016040528092919081815260200182805480156105a657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161055c575b5050505050905090565b60006a084595161401484a000000905090565b6000600560009054906101000a900460ff16905090565b6000600454905090565b6105ec6106df565b15156105f757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600080600061074361101e565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020608060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820160009054906101000a900460ff1615151515815250509050806020015181604001518260600151935093509350509193909250565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600354905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561092357600080fd5b505afa158015610937573d6000803e3d6000fd5b505050506040513d602081101561094d57600080fd5b8101908080519060200190929190505050905090565b600061096d6106df565b151561097857600080fd5b600560009054906101000a900460ff1615156109fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f636f6e7472616374206d75737420626520696e697469616c697a65640000000081525060200191505060405180910390fd5b60035483101515610a75576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6e6f7420656e6f75676820746f6b656e206c656674000000000000000000000081525060200191505060405180910390fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160009054906101000a900460ff16151515610b60576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001807f616464726573732073686f756c64206e6f74206265656e20696e63656e74697681526020017f697a65642079657400000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb30856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610c2557600080fd5b505af1158015610c39573d6000803e3d6000fd5b505050506040513d6020811015610c4f57600080fd5b81019080805190602001909291905050501515610cd4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b6080604051908101604052808573ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200160011515815250600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548160ff02191690831515021790555090505060028490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610e4783600354610ee190919063ffffffff16565b600381905550610e6283600454610f0390919063ffffffff16565b600481905550818473ffffffffffffffffffffffffffffffffffffffff167fac41088a6c8849362a3fb6c807ecde81a69bd79ff8ff86484e766d7d16133dcd856040518082815260200191505060405180910390a3600190509392505050565b610eca6106df565b1515610ed557600080fd5b610ede81610f24565b50565b6000828211151515610ef257600080fd5b600082840390508091505092915050565b6000808284019050838110151515610f1a57600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610f6057600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b608060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008019168152602001600015158152509056fea165627a7a72305820bf6b840d166e4a692423fd8bfe30d57166f395c54aa8800e21b023632a3978290029`

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

// InitialSupply is a free data retrieval call binding the contract method 0x1eb25d13.
//
// Solidity: function InitialSupply() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "InitialSupply")
	return *ret0, err
}

// InitialSupply is a free data retrieval call binding the contract method 0x1eb25d13.
//
// Solidity: function InitialSupply() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) InitialSupply() (*big.Int, error) {
	return _BaasIncentives.Contract.InitialSupply(&_BaasIncentives.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x1eb25d13.
//
// Solidity: function InitialSupply() constant returns(uint256)
func (_BaasIncentives *BaasIncentivesCallerSession) InitialSupply() (*big.Int, error) {
	return _BaasIncentives.Contract.InitialSupply(&_BaasIncentives.CallOpts)
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

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(amount uint256, id bytes32, isValue bool)
func (_BaasIncentives *BaasIncentivesCaller) GetIncentive(opts *bind.CallOpts, account common.Address) (struct {
	Amount  *big.Int
	Id      [32]byte
	IsValue bool
}, error) {
	ret := new(struct {
		Amount  *big.Int
		Id      [32]byte
		IsValue bool
	})
	out := ret
	err := _BaasIncentives.contract.Call(opts, out, "getIncentive", account)
	return *ret, err
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(amount uint256, id bytes32, isValue bool)
func (_BaasIncentives *BaasIncentivesSession) GetIncentive(account common.Address) (struct {
	Amount  *big.Int
	Id      [32]byte
	IsValue bool
}, error) {
	return _BaasIncentives.Contract.GetIncentive(&_BaasIncentives.CallOpts, account)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(amount uint256, id bytes32, isValue bool)
func (_BaasIncentives *BaasIncentivesCallerSession) GetIncentive(account common.Address) (struct {
	Amount  *big.Int
	Id      [32]byte
	IsValue bool
}, error) {
	return _BaasIncentives.Contract.GetIncentive(&_BaasIncentives.CallOpts, account)
}

// Incentives is a free data retrieval call binding the contract method 0x0d5df7ba.
//
// Solidity: function incentives() constant returns(address[])
func (_BaasIncentives *BaasIncentivesCaller) Incentives(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "incentives")
	return *ret0, err
}

// Incentives is a free data retrieval call binding the contract method 0x0d5df7ba.
//
// Solidity: function incentives() constant returns(address[])
func (_BaasIncentives *BaasIncentivesSession) Incentives() ([]common.Address, error) {
	return _BaasIncentives.Contract.Incentives(&_BaasIncentives.CallOpts)
}

// Incentives is a free data retrieval call binding the contract method 0x0d5df7ba.
//
// Solidity: function incentives() constant returns(address[])
func (_BaasIncentives *BaasIncentivesCallerSession) Incentives() ([]common.Address, error) {
	return _BaasIncentives.Contract.Incentives(&_BaasIncentives.CallOpts)
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

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(account address, amount uint256, id bytes32) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Issue(opts *bind.TransactOpts, account common.Address, amount *big.Int, id [32]byte) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "issue", account, amount, id)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(account address, amount uint256, id bytes32) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Issue(account common.Address, amount *big.Int, id [32]byte) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Issue(&_BaasIncentives.TransactOpts, account, amount, id)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(account address, amount uint256, id bytes32) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Issue(account common.Address, amount *big.Int, id [32]byte) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Issue(&_BaasIncentives.TransactOpts, account, amount, id)
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

// BaasIncentivesIncentiveIssuedIterator is returned from FilterIncentiveIssued and is used to iterate over the raw logs and unpacked data for IncentiveIssued events raised by the BaasIncentives contract.
type BaasIncentivesIncentiveIssuedIterator struct {
	Event *BaasIncentivesIncentiveIssued // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesIncentiveIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesIncentiveIssued)
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
		it.Event = new(BaasIncentivesIncentiveIssued)
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
func (it *BaasIncentivesIncentiveIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesIncentiveIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesIncentiveIssued represents a IncentiveIssued event raised by the BaasIncentives contract.
type BaasIncentivesIncentiveIssued struct {
	Account common.Address
	Id      [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentiveIssued is a free log retrieval operation binding the contract event 0xac41088a6c8849362a3fb6c807ecde81a69bd79ff8ff86484e766d7d16133dcd.
//
// Solidity: e IncentiveIssued(account indexed address, id indexed bytes32, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterIncentiveIssued(opts *bind.FilterOpts, account []common.Address, id [][32]byte) (*BaasIncentivesIncentiveIssuedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "IncentiveIssued", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesIncentiveIssuedIterator{contract: _BaasIncentives.contract, event: "IncentiveIssued", logs: logs, sub: sub}, nil
}

// WatchIncentiveIssued is a free log subscription operation binding the contract event 0xac41088a6c8849362a3fb6c807ecde81a69bd79ff8ff86484e766d7d16133dcd.
//
// Solidity: e IncentiveIssued(account indexed address, id indexed bytes32, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchIncentiveIssued(opts *bind.WatchOpts, sink chan<- *BaasIncentivesIncentiveIssued, account []common.Address, id [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "IncentiveIssued", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesIncentiveIssued)
				if err := _BaasIncentives.contract.UnpackLog(event, "IncentiveIssued", log); err != nil {
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

// BaasIncentivesPayoutIterator is returned from FilterPayout and is used to iterate over the raw logs and unpacked data for Payout events raised by the BaasIncentives contract.
type BaasIncentivesPayoutIterator struct {
	Event *BaasIncentivesPayout // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesPayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesPayout)
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
		it.Event = new(BaasIncentivesPayout)
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
func (it *BaasIncentivesPayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesPayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesPayout represents a Payout event raised by the BaasIncentives contract.
type BaasIncentivesPayout struct {
	Account common.Address
	Id      [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayout is a free log retrieval operation binding the contract event 0x212950d8b6a2784adeaa8e088d070610bdfab9eaef5a7127d90118ec1cf34a61.
//
// Solidity: e Payout(account indexed address, id indexed bytes32, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterPayout(opts *bind.FilterOpts, account []common.Address, id [][32]byte) (*BaasIncentivesPayoutIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "Payout", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesPayoutIterator{contract: _BaasIncentives.contract, event: "Payout", logs: logs, sub: sub}, nil
}

// WatchPayout is a free log subscription operation binding the contract event 0x212950d8b6a2784adeaa8e088d070610bdfab9eaef5a7127d90118ec1cf34a61.
//
// Solidity: e Payout(account indexed address, id indexed bytes32, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchPayout(opts *bind.WatchOpts, sink chan<- *BaasIncentivesPayout, account []common.Address, id [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "Payout", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesPayout)
				if err := _BaasIncentives.contract.UnpackLog(event, "Payout", log); err != nil {
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

// BaasIncentivesRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the BaasIncentives contract.
type BaasIncentivesRevokedIterator struct {
	Event *BaasIncentivesRevoked // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesRevoked)
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
		it.Event = new(BaasIncentivesRevoked)
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
func (it *BaasIncentivesRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesRevoked represents a Revoked event raised by the BaasIncentives contract.
type BaasIncentivesRevoked struct {
	Account           common.Address
	Id                [32]byte
	TokenNotDelivered *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0xe490c408aa61d6a8f6ab6dd0b3722b31d05c4ac6dbb268c4472052fb8d89d8df.
//
// Solidity: e Revoked(account indexed address, id indexed bytes32, tokenNotDelivered uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterRevoked(opts *bind.FilterOpts, account []common.Address, id [][32]byte) (*BaasIncentivesRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "Revoked", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesRevokedIterator{contract: _BaasIncentives.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0xe490c408aa61d6a8f6ab6dd0b3722b31d05c4ac6dbb268c4472052fb8d89d8df.
//
// Solidity: e Revoked(account indexed address, id indexed bytes32, tokenNotDelivered uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *BaasIncentivesRevoked, account []common.Address, id [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "Revoked", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesRevoked)
				if err := _BaasIncentives.contract.UnpackLog(event, "Revoked", log); err != nil {
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
