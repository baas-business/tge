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
const BaasEscrowABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CapitalRaised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiseCapital\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCapitalRaiseOngoing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentCapitalRaiseId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"raisedCapitalTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuedTokenTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnedTokenTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"capitalRaiseId\",\"type\":\"uint256\"}],\"name\":\"capitalRaise\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"isFinalized\",\"type\":\"bool\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capitalRaiseIds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasEscrowBin is the compiled bytecode used for deploying new contracts.
const BaasEscrowBin = `0x608060405234801561001057600080fd5b5060405160208061175e8339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506116108061014e6000396000f3fe6080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630a40a5b5146100f65780631d41ad9414610121578063321a83dc1461018d578063378dc3dc146101bc5780634bb278f3146101e75780634f4a7a7914610216578063715018a61461029057806376c7399e146102a7578063867904b4146102d25780638da5cb5b146103455780638f32d59b1461039c5780639d76ea58146103cb578063b69ef8a814610422578063dc0fb4cc1461044d578063f2fde38b14610478578063f605e1ce146104c9578063fbd207ed146104f4575b600080fd5b34801561010257600080fd5b5061010b610551565b6040518082815260200191505060405180910390f35b34801561012d57600080fd5b5061013661055b565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561017957808201518184015260208101905061015e565b505050509050019250505060405180910390f35b34801561019957600080fd5b506101a26105b3565b604051808215151515815260200191505060405180910390f35b3480156101c857600080fd5b506101d16105c0565b6040518082815260200191505060405180910390f35b3480156101f357600080fd5b506101fc6105d3565b604051808215151515815260200191505060405180910390f35b34801561022257600080fd5b5061024f6004803603602081101561023957600080fd5b8101908080359060200190929190505050610923565b604051808781526020018681526020018581526020018481526020018315151515815260200182151515158152602001965050505050505060405180910390f35b34801561029c57600080fd5b506102a56109e7565b005b3480156102b357600080fd5b506102bc610ab9565b6040518082815260200191505060405180910390f35b3480156102de57600080fd5b5061032b600480360360408110156102f557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ac3565b604051808215151515815260200191505060405180910390f35b34801561035157600080fd5b5061035a610ea2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103a857600080fd5b506103b1610ecb565b604051808215151515815260200191505060405180910390f35b3480156103d757600080fd5b506103e0610f22565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561042e57600080fd5b50610437610f4c565b6040518082815260200191505060405180910390f35b34801561045957600080fd5b50610462611049565b6040518082815260200191505060405180910390f35b34801561048457600080fd5b506104c76004803603602081101561049b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611053565b005b3480156104d557600080fd5b506104de611072565b6040518082815260200191505060405180910390f35b34801561050057600080fd5b506105376004803603604081101561051757600080fd5b81019080803590602001909291908035906020019092919050505061107c565b604051808215151515815260200191505060405180910390f35b6000600354905090565b606060028054806020026020016040519081016040528092919081815260200182805480156105a957602002820191906000526020600020905b815481526020019060010190808311610595575b5050505050905090565b6000806006541415905090565b60006a31a17e847807b1bc000000905090565b60006105dd610ecb565b15156105e857600080fd5b6105f06105b3565b1515610664576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e6f206361706974616c207261697365206f6e676f696e67000000000000000081525060200191505060405180910390fd5b61066c6115a9565b60016000600654815260200190815260200160002060c06040519081016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581526020016004820160019054906101000a900460ff1615151515815250509050600061070a8260400151836020015161146c90919063ffffffff16565b905060006006549050600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634cca31b630846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156107d857600080fd5b505af11580156107ec573d6000803e3d6000fd5b505050506040513d602081101561080257600080fd5b81019080805190602001909291905050501515610887576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6275726e696e67206f6620746f6b656e73206661696c6564000000000000000081525060200191505060405180910390fd5b61089c8260055461148e90919063ffffffff16565b6005819055506001806000600654815260200190815260200160002060040160006101000a81548160ff02191690831515021790555060006006819055507fb968440accd1ce5fa60b00de8bb8d8487eb2fda3c3701fb30fea3f69aa910a488183604051808381526020018281526020019250505060405180910390a16001935050505090565b6000806000806000806109346115a9565b6001600089815260200190815260200160002060c06040519081016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581526020016004820160019054906101000a900460ff1615151515815250509050806000015181602001518260400151836060015184608001518560a001519650965096509650965096505091939550919395565b6109ef610ecb565b15156109fa57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600554905090565b6000610acd610ecb565b1515610ad857600080fd5b610ae06105b3565b1515610b54576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e6f206361706974616c207261697365206f6e676f696e67000000000000000081525060200191505060405180910390fd5b610b5c6115a9565b60016000600654815260200190815260200160002060c06040519081016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581526020016004820160019054906101000a900460ff1615151515815250509050610bf88160400151826020015161146c90919063ffffffff16565b83101515610c94576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f616d6f756e742073686f756c64206265206c657373207468656e20776861742081526020017f6973206c6566740000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610d5957600080fd5b505af1158015610d6d573d6000803e3d6000fd5b505050506040513d6020811015610d8357600080fd5b81019080805190602001909291905050501515610e08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f7472616e73666572206661696c6564000000000000000000000000000000000081525060200191505060405180910390fd5b8260016000600654815260200190815260200160002060020160008282540192505081905550610e438360045461148e90919063ffffffff16565b6004819055508373ffffffffffffffffffffffffffffffffffffffff167f21d739f160a7464fddaac4a1d1517d84e76b75618a053943b345c408c4160fe0846040518082815260200191505060405180910390a2600191505092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561100957600080fd5b505afa15801561101d573d6000803e3d6000fd5b505050506040513d602081101561103357600080fd5b8101908080519060200190929190505050905090565b6000600454905090565b61105b610ecb565b151561106657600080fd5b61106f816114af565b50565b6000600654905090565b6000611086610ecb565b151561109157600080fd5b6000821415151561110a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6361706974616c2072616973652069642063616e6e6f7420626520300000000081525060200191505060405180910390fd5b61112a6003546a31a17e847807b1bc00000061146c90919063ffffffff16565b83111515156111c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c8152602001807f6e6f7420656e6f75676820746f6b656e206c65667420666f722074686973206381526020017f61706974616c207261697365000000000000000000000000000000000000000081525060400191505060405180910390fd5b6001600083815260200190815260200160002060040160019054906101000a900460ff16151515611286576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f646f6e277420757365207468652073616d65206361706974616c20726169736581526020017f206964207477696365000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b61128e6105b3565b151515611329576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001807f61206361706974616c207261697365206973206f6e676f696e6720616e64206e81526020017f6f742066696e616c697a6564207965740000000000000000000000000000000081525060400191505060405180910390fd5b60c0604051908101604052808381526020018481526020016000815260200143815260200160001515815260200160011515815250600160008481526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a81548160ff02191690831515021790555060a08201518160040160016101000a81548160ff0219169083151502179055509050506002829080600181540180825580915050906001820390600052602060002001600090919290919091505550816006819055506114248360035461148e90919063ffffffff16565b600381905550817fbc65650fcb1dd5a9e3c92e066533c3425c66a51f1f60fbfe11ce352452c9b08d846040518082815260200191505060405180910390a26001905092915050565b600082821115151561147d57600080fd5b600082840390508091505092915050565b60008082840190508381101515156114a557600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156114eb57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60c06040519081016040528060008152602001600081526020016000815260200160008152602001600015158152602001600015158152509056fea165627a7a723058205b3c6e4e811960982d91e1ca2193fc9462c6d78b1fadee4500dec79c13710a6a0029`

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

// BurnedTokenTotal is a free data retrieval call binding the contract method 0x76c7399e.
//
// Solidity: function burnedTokenTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) BurnedTokenTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "burnedTokenTotal")
	return *ret0, err
}

// BurnedTokenTotal is a free data retrieval call binding the contract method 0x76c7399e.
//
// Solidity: function burnedTokenTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) BurnedTokenTotal() (*big.Int, error) {
	return _BaasEscrow.Contract.BurnedTokenTotal(&_BaasEscrow.CallOpts)
}

// BurnedTokenTotal is a free data retrieval call binding the contract method 0x76c7399e.
//
// Solidity: function burnedTokenTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) BurnedTokenTotal() (*big.Int, error) {
	return _BaasEscrow.Contract.BurnedTokenTotal(&_BaasEscrow.CallOpts)
}

// CapitalRaise is a free data retrieval call binding the contract method 0x4f4a7a79.
//
// Solidity: function capitalRaise(capitalRaiseId uint256) constant returns(id uint256, amount uint256, provided uint256, atBlock uint256, isFinalized bool, isValue bool)
func (_BaasEscrow *BaasEscrowCaller) CapitalRaise(opts *bind.CallOpts, capitalRaiseId *big.Int) (struct {
	Id          *big.Int
	Amount      *big.Int
	Provided    *big.Int
	AtBlock     *big.Int
	IsFinalized bool
	IsValue     bool
}, error) {
	ret := new(struct {
		Id          *big.Int
		Amount      *big.Int
		Provided    *big.Int
		AtBlock     *big.Int
		IsFinalized bool
		IsValue     bool
	})
	out := ret
	err := _BaasEscrow.contract.Call(opts, out, "capitalRaise", capitalRaiseId)
	return *ret, err
}

// CapitalRaise is a free data retrieval call binding the contract method 0x4f4a7a79.
//
// Solidity: function capitalRaise(capitalRaiseId uint256) constant returns(id uint256, amount uint256, provided uint256, atBlock uint256, isFinalized bool, isValue bool)
func (_BaasEscrow *BaasEscrowSession) CapitalRaise(capitalRaiseId *big.Int) (struct {
	Id          *big.Int
	Amount      *big.Int
	Provided    *big.Int
	AtBlock     *big.Int
	IsFinalized bool
	IsValue     bool
}, error) {
	return _BaasEscrow.Contract.CapitalRaise(&_BaasEscrow.CallOpts, capitalRaiseId)
}

// CapitalRaise is a free data retrieval call binding the contract method 0x4f4a7a79.
//
// Solidity: function capitalRaise(capitalRaiseId uint256) constant returns(id uint256, amount uint256, provided uint256, atBlock uint256, isFinalized bool, isValue bool)
func (_BaasEscrow *BaasEscrowCallerSession) CapitalRaise(capitalRaiseId *big.Int) (struct {
	Id          *big.Int
	Amount      *big.Int
	Provided    *big.Int
	AtBlock     *big.Int
	IsFinalized bool
	IsValue     bool
}, error) {
	return _BaasEscrow.Contract.CapitalRaise(&_BaasEscrow.CallOpts, capitalRaiseId)
}

// CapitalRaiseIds is a free data retrieval call binding the contract method 0x1d41ad94.
//
// Solidity: function capitalRaiseIds() constant returns(uint256[])
func (_BaasEscrow *BaasEscrowCaller) CapitalRaiseIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "capitalRaiseIds")
	return *ret0, err
}

// CapitalRaiseIds is a free data retrieval call binding the contract method 0x1d41ad94.
//
// Solidity: function capitalRaiseIds() constant returns(uint256[])
func (_BaasEscrow *BaasEscrowSession) CapitalRaiseIds() ([]*big.Int, error) {
	return _BaasEscrow.Contract.CapitalRaiseIds(&_BaasEscrow.CallOpts)
}

// CapitalRaiseIds is a free data retrieval call binding the contract method 0x1d41ad94.
//
// Solidity: function capitalRaiseIds() constant returns(uint256[])
func (_BaasEscrow *BaasEscrowCallerSession) CapitalRaiseIds() ([]*big.Int, error) {
	return _BaasEscrow.Contract.CapitalRaiseIds(&_BaasEscrow.CallOpts)
}

// CurrentCapitalRaiseId is a free data retrieval call binding the contract method 0xf605e1ce.
//
// Solidity: function currentCapitalRaiseId() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) CurrentCapitalRaiseId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "currentCapitalRaiseId")
	return *ret0, err
}

// CurrentCapitalRaiseId is a free data retrieval call binding the contract method 0xf605e1ce.
//
// Solidity: function currentCapitalRaiseId() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) CurrentCapitalRaiseId() (*big.Int, error) {
	return _BaasEscrow.Contract.CurrentCapitalRaiseId(&_BaasEscrow.CallOpts)
}

// CurrentCapitalRaiseId is a free data retrieval call binding the contract method 0xf605e1ce.
//
// Solidity: function currentCapitalRaiseId() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) CurrentCapitalRaiseId() (*big.Int, error) {
	return _BaasEscrow.Contract.CurrentCapitalRaiseId(&_BaasEscrow.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "initialSupply")
	return *ret0, err
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) InitialSupply() (*big.Int, error) {
	return _BaasEscrow.Contract.InitialSupply(&_BaasEscrow.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) InitialSupply() (*big.Int, error) {
	return _BaasEscrow.Contract.InitialSupply(&_BaasEscrow.CallOpts)
}

// IsCapitalRaiseOngoing is a free data retrieval call binding the contract method 0x321a83dc.
//
// Solidity: function isCapitalRaiseOngoing() constant returns(bool)
func (_BaasEscrow *BaasEscrowCaller) IsCapitalRaiseOngoing(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "isCapitalRaiseOngoing")
	return *ret0, err
}

// IsCapitalRaiseOngoing is a free data retrieval call binding the contract method 0x321a83dc.
//
// Solidity: function isCapitalRaiseOngoing() constant returns(bool)
func (_BaasEscrow *BaasEscrowSession) IsCapitalRaiseOngoing() (bool, error) {
	return _BaasEscrow.Contract.IsCapitalRaiseOngoing(&_BaasEscrow.CallOpts)
}

// IsCapitalRaiseOngoing is a free data retrieval call binding the contract method 0x321a83dc.
//
// Solidity: function isCapitalRaiseOngoing() constant returns(bool)
func (_BaasEscrow *BaasEscrowCallerSession) IsCapitalRaiseOngoing() (bool, error) {
	return _BaasEscrow.Contract.IsCapitalRaiseOngoing(&_BaasEscrow.CallOpts)
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

// IssuedTokenTotal is a free data retrieval call binding the contract method 0xdc0fb4cc.
//
// Solidity: function issuedTokenTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) IssuedTokenTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "issuedTokenTotal")
	return *ret0, err
}

// IssuedTokenTotal is a free data retrieval call binding the contract method 0xdc0fb4cc.
//
// Solidity: function issuedTokenTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) IssuedTokenTotal() (*big.Int, error) {
	return _BaasEscrow.Contract.IssuedTokenTotal(&_BaasEscrow.CallOpts)
}

// IssuedTokenTotal is a free data retrieval call binding the contract method 0xdc0fb4cc.
//
// Solidity: function issuedTokenTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) IssuedTokenTotal() (*big.Int, error) {
	return _BaasEscrow.Contract.IssuedTokenTotal(&_BaasEscrow.CallOpts)
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

// RaisedCapitalTotal is a free data retrieval call binding the contract method 0x0a40a5b5.
//
// Solidity: function raisedCapitalTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) RaisedCapitalTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "raisedCapitalTotal")
	return *ret0, err
}

// RaisedCapitalTotal is a free data retrieval call binding the contract method 0x0a40a5b5.
//
// Solidity: function raisedCapitalTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) RaisedCapitalTotal() (*big.Int, error) {
	return _BaasEscrow.Contract.RaisedCapitalTotal(&_BaasEscrow.CallOpts)
}

// RaisedCapitalTotal is a free data retrieval call binding the contract method 0x0a40a5b5.
//
// Solidity: function raisedCapitalTotal() constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) RaisedCapitalTotal() (*big.Int, error) {
	return _BaasEscrow.Contract.RaisedCapitalTotal(&_BaasEscrow.CallOpts)
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

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_BaasEscrow *BaasEscrowSession) Finalize() (*types.Transaction, error) {
	return _BaasEscrow.Contract.Finalize(&_BaasEscrow.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) Finalize() (*types.Transaction, error) {
	return _BaasEscrow.Contract.Finalize(&_BaasEscrow.TransactOpts)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(account address, amount uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) Issue(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "issue", account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(account address, amount uint256) returns(bool)
func (_BaasEscrow *BaasEscrowSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.Issue(&_BaasEscrow.TransactOpts, account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(account address, amount uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.Issue(&_BaasEscrow.TransactOpts, account, amount)
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

// BaasEscrowFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the BaasEscrow contract.
type BaasEscrowFinalizedIterator struct {
	Event *BaasEscrowFinalized // Event containing the contract specifics and raw log

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
func (it *BaasEscrowFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasEscrowFinalized)
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
		it.Event = new(BaasEscrowFinalized)
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
func (it *BaasEscrowFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasEscrowFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasEscrowFinalized represents a Finalized event raised by the BaasEscrow contract.
type BaasEscrowFinalized struct {
	Id           *big.Int
	BurnedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0xb968440accd1ce5fa60b00de8bb8d8487eb2fda3c3701fb30fea3f69aa910a48.
//
// Solidity: e Finalized(id uint256, burnedAmount uint256)
func (_BaasEscrow *BaasEscrowFilterer) FilterFinalized(opts *bind.FilterOpts) (*BaasEscrowFinalizedIterator, error) {

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &BaasEscrowFinalizedIterator{contract: _BaasEscrow.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0xb968440accd1ce5fa60b00de8bb8d8487eb2fda3c3701fb30fea3f69aa910a48.
//
// Solidity: e Finalized(id uint256, burnedAmount uint256)
func (_BaasEscrow *BaasEscrowFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *BaasEscrowFinalized) (event.Subscription, error) {

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasEscrowFinalized)
				if err := _BaasEscrow.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// BaasEscrowTokensIssuedIterator is returned from FilterTokensIssued and is used to iterate over the raw logs and unpacked data for TokensIssued events raised by the BaasEscrow contract.
type BaasEscrowTokensIssuedIterator struct {
	Event *BaasEscrowTokensIssued // Event containing the contract specifics and raw log

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
func (it *BaasEscrowTokensIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasEscrowTokensIssued)
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
		it.Event = new(BaasEscrowTokensIssued)
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
func (it *BaasEscrowTokensIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasEscrowTokensIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasEscrowTokensIssued represents a TokensIssued event raised by the BaasEscrow contract.
type BaasEscrowTokensIssued struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensIssued is a free log retrieval operation binding the contract event 0x21d739f160a7464fddaac4a1d1517d84e76b75618a053943b345c408c4160fe0.
//
// Solidity: e TokensIssued(account indexed address, amount uint256)
func (_BaasEscrow *BaasEscrowFilterer) FilterTokensIssued(opts *bind.FilterOpts, account []common.Address) (*BaasEscrowTokensIssuedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasEscrow.contract.FilterLogs(opts, "TokensIssued", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasEscrowTokensIssuedIterator{contract: _BaasEscrow.contract, event: "TokensIssued", logs: logs, sub: sub}, nil
}

// WatchTokensIssued is a free log subscription operation binding the contract event 0x21d739f160a7464fddaac4a1d1517d84e76b75618a053943b345c408c4160fe0.
//
// Solidity: e TokensIssued(account indexed address, amount uint256)
func (_BaasEscrow *BaasEscrowFilterer) WatchTokensIssued(opts *bind.WatchOpts, sink chan<- *BaasEscrowTokensIssued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasEscrow.contract.WatchLogs(opts, "TokensIssued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasEscrowTokensIssued)
				if err := _BaasEscrow.contract.UnpackLog(event, "TokensIssued", log); err != nil {
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
