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

// BaasROIABI is the input ABI used to generate the binding from.
const BaasROIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"omittedReceiver\",\"type\":\"address\"}],\"name\":\"PayoutOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokensReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensPossessed\",\"type\":\"uint256\"}],\"name\":\"PaidOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokensProvided\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensPossessed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenHolders\",\"type\":\"uint256\"}],\"name\":\"PaidOutAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"payoutAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawFromContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eligibleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"contractHasEnoughTokensForPayout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"currentPayoutObligation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"optimalPayoutDistribution\",\"outputs\":[{\"name\":\"contractCanPayMax\",\"type\":\"uint256\"},{\"name\":\"contractShouldHaveMin\",\"type\":\"uint256\"},{\"name\":\"error\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\"},{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"roi\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasROIBin is the compiled bytecode used for deploying new contracts.
const BaasROIBin = `0x608060405234801561001057600080fd5b506040516020806117d18339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506116838061014e6000396000f3fe6080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301b32f0a146100e05780631473ada31461013957806352e26a1514610196578063715018a6146101e95780637c3a00fd146102005780638da5cb5b1461022b5780638f32d59b146102825780639358928b146102b15780639d76ea58146102dc578063b69ef8a814610333578063bb7314671461035e578063be4d0aa214610389578063daf38419146103d8578063f2fde38b1461042b578063f3ff49611461047c575b600080fd5b3480156100ec57600080fd5b506101236004803603604081101561010357600080fd5b8101908080359060200190929190803590602001909291905050506104e5565b6040518082815260200191505060405180910390f35b34801561014557600080fd5b506101726004803603602081101561015c57600080fd5b8101908080359060200190929190505050610527565b60405180848152602001838152602001828152602001935050505060405180910390f35b3480156101a257600080fd5b506101cf600480360360208110156101b957600080fd5b81019080803590602001909291905050506105df565b604051808215151515815260200191505060405180910390f35b3480156101f557600080fd5b506101fe61060f565b005b34801561020c57600080fd5b506102156106e1565b6040518082815260200191505060405180910390f35b34801561023757600080fd5b506102406106ea565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561028e57600080fd5b50610297610713565b604051808215151515815260200191505060405180910390f35b3480156102bd57600080fd5b506102c661076a565b6040518082815260200191505060405180910390f35b3480156102e857600080fd5b506102f1610830565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561033f57600080fd5b5061034861085a565b6040518082815260200191505060405180910390f35b34801561036a57600080fd5b50610373610957565b6040518082815260200191505060405180910390f35b34801561039557600080fd5b506103c2600480360360208110156103ac57600080fd5b8101908080359060200190929190505050610a36565b6040518082815260200191505060405180910390f35b3480156103e457600080fd5b50610411600480360360208110156103fb57600080fd5b8101908080359060200190929190505050610a50565b604051808215151515815260200191505060405180910390f35b34801561043757600080fd5b5061047a6004803603602081101561044e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f50565b005b34801561048857600080fd5b506104cb6004803603602081101561049f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f6f565b604051808215151515815260200191505060405180910390f35b6000808214156104f85760009050610521565b61051e826105106009866111d490919063ffffffff16565b61121290919063ffffffff16565b90505b92915050565b60008060008061054185600961123c90919063ffffffff16565b905061057081610562600961055461076a565b6111d490919063ffffffff16565b61121290919063ffffffff16565b925061059e816105908761058261076a565b6111d490919063ffffffff16565b61121290919063ffffffff16565b93506105cc846105be856105b061076a565b61125d90919063ffffffff16565b61125d90919063ffffffff16565b9150838383935093509350509193909250565b6000806000806105ee85610527565b80935081945082955050505061060261085a565b8211159350505050919050565b610617610713565b151561062257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006009905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639358928b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156107f057600080fd5b505afa158015610804573d6000803e3d6000fd5b505050506040513d602081101561081a57600080fd5b8101908080519060200190929190505050905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561091757600080fd5b505afa15801561092b573d6000803e3d6000fd5b505050506040513d602081101561094157600080fd5b8101908080519060200190929190505050905090565b6000610a3161096461085a565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639358928b6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b1580156109e857600080fd5b505afa1580156109fc573d6000803e3d6000fd5b505050506040513d6020811015610a1257600080fd5b810190808051906020019092919050505061125d90919063ffffffff16565b905090565b6000610a49610a43610957565b836104e5565b9050919050565b6000610a5a610713565b1515610a6557600080fd5b610a6e826105df565b1515610ae2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f636f6e7472616374732062616c616e636520697320746f6f206c6f770000000081525060200191505060405180910390fd5b6000610aec61085a565b9050600080600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a654cfab6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160806040518083038186803b158015610b7857600080fd5b505afa158015610b8c573d6000803e3d6000fd5b505050506040513d6080811015610ba257600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050809450819550829650839750505050506060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637af70c1f6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015610c6957600080fd5b505af1158015610c7d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015610ca757600080fd5b810190808051640100000000811115610cbf57600080fd5b82810190506020810184811115610cd557600080fd5b8151856020820283011164010000000082111715610cf257600080fd5b505092919050505090506000815190506000809050600080905060008090505b83811015610efa5760008582815181101515610d2a57fe5b9060200190602002015190508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161480610d9b57508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80610dd157508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80610e0757508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b80610e3d57503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15610eaa577f1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb7281604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1610eec565b6000610ecd8784815181101515610ebd57fe5b906020019060200201518f61127f565b9050610ee2818661123c90919063ffffffff16565b9450600184019350505b508080600101915050610d12565b507fda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9828a8360405180848152602001838152602001828152602001935050505060405180910390a1505050505050505050919050565b610f58610713565b1515610f6357600080fd5b610f6c8161155d565b50565b6000610f79610713565b1515610f8457600080fd5b6000610f8e61085a565b9050600081111515611008576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f6e6f20746f6b656e73206c65667420696e20636f6e747261637400000000000081525060200191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156110cd57600080fd5b505af11580156110e1573d6000803e3d6000fd5b505050506040513d60208110156110f757600080fd5b8101908080519060200190929190505050151561117c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5826040518082815260200191505060405180910390a26001915050919050565b6000808314156111e7576000905061120c565b600082840290508284828115156111fa57fe5b0414151561120757600080fd5b809150505b92915050565b6000808211151561122257600080fd5b6000828481151561122f57fe5b0490508091505092915050565b600080828401905083811015151561125357600080fd5b8091505092915050565b600082821115151561126e57600080fd5b600082840390508091505092915050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561133d57600080fd5b505afa158015611351573d6000803e3d6000fd5b505050506040513d602081101561136757600080fd5b81019080805190602001909291905050509050600061138682856104e5565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561144d57600080fd5b505af1158015611461573d6000803e3d6000fd5b505050506040513d602081101561147757600080fd5b810190808051906020019092919050505015156114fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f7472616e73666572206f6620746f6b656e73206661696c65640000000000000081525060200191505060405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff167f7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb8284604051808381526020018281526020019250505060405180910390a2809250505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561159957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fea165627a7a72305820443c34c44310441060dc947bb5de1822f1242202039cc7a033c3f20b1715ae6b0029`

// DeployBaasROI deploys a new Ethereum contract, binding an instance of BaasROI to it.
func DeployBaasROI(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasROI, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasROIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasROIBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasROI{BaasROICaller: BaasROICaller{contract: contract}, BaasROITransactor: BaasROITransactor{contract: contract}, BaasROIFilterer: BaasROIFilterer{contract: contract}}, nil
}

// BaasROI is an auto generated Go binding around an Ethereum contract.
type BaasROI struct {
	BaasROICaller     // Read-only binding to the contract
	BaasROITransactor // Write-only binding to the contract
	BaasROIFilterer   // Log filterer for contract events
}

// BaasROICaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasROICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasROITransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasROITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasROIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasROIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasROISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasROISession struct {
	Contract     *BaasROI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasROICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasROICallerSession struct {
	Contract *BaasROICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BaasROITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasROITransactorSession struct {
	Contract     *BaasROITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BaasROIRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasROIRaw struct {
	Contract *BaasROI // Generic contract binding to access the raw methods on
}

// BaasROICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasROICallerRaw struct {
	Contract *BaasROICaller // Generic read-only contract binding to access the raw methods on
}

// BaasROITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasROITransactorRaw struct {
	Contract *BaasROITransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasROI creates a new instance of BaasROI, bound to a specific deployed contract.
func NewBaasROI(address common.Address, backend bind.ContractBackend) (*BaasROI, error) {
	contract, err := bindBaasROI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasROI{BaasROICaller: BaasROICaller{contract: contract}, BaasROITransactor: BaasROITransactor{contract: contract}, BaasROIFilterer: BaasROIFilterer{contract: contract}}, nil
}

// NewBaasROICaller creates a new read-only instance of BaasROI, bound to a specific deployed contract.
func NewBaasROICaller(address common.Address, caller bind.ContractCaller) (*BaasROICaller, error) {
	contract, err := bindBaasROI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasROICaller{contract: contract}, nil
}

// NewBaasROITransactor creates a new write-only instance of BaasROI, bound to a specific deployed contract.
func NewBaasROITransactor(address common.Address, transactor bind.ContractTransactor) (*BaasROITransactor, error) {
	contract, err := bindBaasROI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasROITransactor{contract: contract}, nil
}

// NewBaasROIFilterer creates a new log filterer instance of BaasROI, bound to a specific deployed contract.
func NewBaasROIFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasROIFilterer, error) {
	contract, err := bindBaasROI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasROIFilterer{contract: contract}, nil
}

// bindBaasROI binds a generic wrapper to an already deployed contract.
func bindBaasROI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasROIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasROI *BaasROIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasROI.Contract.BaasROICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasROI *BaasROIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasROI.Contract.BaasROITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasROI *BaasROIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasROI.Contract.BaasROITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasROI *BaasROICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasROI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasROI *BaasROITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasROI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasROI *BaasROITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasROI.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasROI *BaasROICaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasROI *BaasROISession) Balance() (*big.Int, error) {
	return _BaasROI.Contract.Balance(&_BaasROI.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) Balance() (*big.Int, error) {
	return _BaasROI.Contract.Balance(&_BaasROI.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasROI *BaasROICaller) CirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "circulatingSupply")
	return *ret0, err
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasROI *BaasROISession) CirculatingSupply() (*big.Int, error) {
	return _BaasROI.Contract.CirculatingSupply(&_BaasROI.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) CirculatingSupply() (*big.Int, error) {
	return _BaasROI.Contract.CirculatingSupply(&_BaasROI.CallOpts)
}

// ContractHasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x52e26a15.
//
// Solidity: function contractHasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROICaller) ContractHasEnoughTokensForPayout(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "contractHasEnoughTokensForPayout", tokenEuroConversionRate)
	return *ret0, err
}

// ContractHasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x52e26a15.
//
// Solidity: function contractHasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROISession) ContractHasEnoughTokensForPayout(tokenEuroConversionRate *big.Int) (bool, error) {
	return _BaasROI.Contract.ContractHasEnoughTokensForPayout(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// ContractHasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x52e26a15.
//
// Solidity: function contractHasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROICallerSession) ContractHasEnoughTokensForPayout(tokenEuroConversionRate *big.Int) (bool, error) {
	return _BaasROI.Contract.ContractHasEnoughTokensForPayout(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// CurrentPayoutObligation is a free data retrieval call binding the contract method 0xbe4d0aa2.
//
// Solidity: function currentPayoutObligation(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICaller) CurrentPayoutObligation(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "currentPayoutObligation", tokenEuroConversionRate)
	return *ret0, err
}

// CurrentPayoutObligation is a free data retrieval call binding the contract method 0xbe4d0aa2.
//
// Solidity: function currentPayoutObligation(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROISession) CurrentPayoutObligation(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.CurrentPayoutObligation(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// CurrentPayoutObligation is a free data retrieval call binding the contract method 0xbe4d0aa2.
//
// Solidity: function currentPayoutObligation(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICallerSession) CurrentPayoutObligation(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.CurrentPayoutObligation(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// EligibleToken is a free data retrieval call binding the contract method 0xbb731467.
//
// Solidity: function eligibleToken() constant returns(uint256)
func (_BaasROI *BaasROICaller) EligibleToken(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "eligibleToken")
	return *ret0, err
}

// EligibleToken is a free data retrieval call binding the contract method 0xbb731467.
//
// Solidity: function eligibleToken() constant returns(uint256)
func (_BaasROI *BaasROISession) EligibleToken() (*big.Int, error) {
	return _BaasROI.Contract.EligibleToken(&_BaasROI.CallOpts)
}

// EligibleToken is a free data retrieval call binding the contract method 0xbb731467.
//
// Solidity: function eligibleToken() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) EligibleToken() (*big.Int, error) {
	return _BaasROI.Contract.EligibleToken(&_BaasROI.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() constant returns(uint256)
func (_BaasROI *BaasROICaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "interestRate")
	return *ret0, err
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() constant returns(uint256)
func (_BaasROI *BaasROISession) InterestRate() (*big.Int, error) {
	return _BaasROI.Contract.InterestRate(&_BaasROI.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() constant returns(uint256)
func (_BaasROI *BaasROICallerSession) InterestRate() (*big.Int, error) {
	return _BaasROI.Contract.InterestRate(&_BaasROI.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasROI *BaasROICaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasROI *BaasROISession) IsOwner() (bool, error) {
	return _BaasROI.Contract.IsOwner(&_BaasROI.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasROI *BaasROICallerSession) IsOwner() (bool, error) {
	return _BaasROI.Contract.IsOwner(&_BaasROI.CallOpts)
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(contractCanPayMax uint256, contractShouldHaveMin uint256, error uint256)
func (_BaasROI *BaasROICaller) OptimalPayoutDistribution(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (struct {
	ContractCanPayMax     *big.Int
	ContractShouldHaveMin *big.Int
	Error                 *big.Int
}, error) {
	ret := new(struct {
		ContractCanPayMax     *big.Int
		ContractShouldHaveMin *big.Int
		Error                 *big.Int
	})
	out := ret
	err := _BaasROI.contract.Call(opts, out, "optimalPayoutDistribution", tokenEuroConversionRate)
	return *ret, err
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(contractCanPayMax uint256, contractShouldHaveMin uint256, error uint256)
func (_BaasROI *BaasROISession) OptimalPayoutDistribution(tokenEuroConversionRate *big.Int) (struct {
	ContractCanPayMax     *big.Int
	ContractShouldHaveMin *big.Int
	Error                 *big.Int
}, error) {
	return _BaasROI.Contract.OptimalPayoutDistribution(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(contractCanPayMax uint256, contractShouldHaveMin uint256, error uint256)
func (_BaasROI *BaasROICallerSession) OptimalPayoutDistribution(tokenEuroConversionRate *big.Int) (struct {
	ContractCanPayMax     *big.Int
	ContractShouldHaveMin *big.Int
	Error                 *big.Int
}, error) {
	return _BaasROI.Contract.OptimalPayoutDistribution(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasROI *BaasROICaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasROI *BaasROISession) Owner() (common.Address, error) {
	return _BaasROI.Contract.Owner(&_BaasROI.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasROI *BaasROICallerSession) Owner() (common.Address, error) {
	return _BaasROI.Contract.Owner(&_BaasROI.CallOpts)
}

// Roi is a free data retrieval call binding the contract method 0x01b32f0a.
//
// Solidity: function roi(token uint256, tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICaller) Roi(opts *bind.CallOpts, token *big.Int, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "roi", token, tokenEuroConversionRate)
	return *ret0, err
}

// Roi is a free data retrieval call binding the contract method 0x01b32f0a.
//
// Solidity: function roi(token uint256, tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROISession) Roi(token *big.Int, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.Roi(&_BaasROI.CallOpts, token, tokenEuroConversionRate)
}

// Roi is a free data retrieval call binding the contract method 0x01b32f0a.
//
// Solidity: function roi(token uint256, tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICallerSession) Roi(token *big.Int, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.Roi(&_BaasROI.CallOpts, token, tokenEuroConversionRate)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasROI *BaasROICaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasROI *BaasROISession) TokenAddress() (common.Address, error) {
	return _BaasROI.Contract.TokenAddress(&_BaasROI.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasROI *BaasROICallerSession) TokenAddress() (common.Address, error) {
	return _BaasROI.Contract.TokenAddress(&_BaasROI.CallOpts)
}

// PayoutAll is a paid mutator transaction binding the contract method 0xdaf38419.
//
// Solidity: function payoutAll(tokenEuroConversionRate uint256) returns(bool)
func (_BaasROI *BaasROITransactor) PayoutAll(opts *bind.TransactOpts, tokenEuroConversionRate *big.Int) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "payoutAll", tokenEuroConversionRate)
}

// PayoutAll is a paid mutator transaction binding the contract method 0xdaf38419.
//
// Solidity: function payoutAll(tokenEuroConversionRate uint256) returns(bool)
func (_BaasROI *BaasROISession) PayoutAll(tokenEuroConversionRate *big.Int) (*types.Transaction, error) {
	return _BaasROI.Contract.PayoutAll(&_BaasROI.TransactOpts, tokenEuroConversionRate)
}

// PayoutAll is a paid mutator transaction binding the contract method 0xdaf38419.
//
// Solidity: function payoutAll(tokenEuroConversionRate uint256) returns(bool)
func (_BaasROI *BaasROITransactorSession) PayoutAll(tokenEuroConversionRate *big.Int) (*types.Transaction, error) {
	return _BaasROI.Contract.PayoutAll(&_BaasROI.TransactOpts, tokenEuroConversionRate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasROI *BaasROITransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasROI *BaasROISession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasROI.Contract.RenounceOwnership(&_BaasROI.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasROI *BaasROITransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasROI.Contract.RenounceOwnership(&_BaasROI.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasROI *BaasROITransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasROI *BaasROISession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.TransferOwnership(&_BaasROI.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasROI *BaasROITransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.TransferOwnership(&_BaasROI.TransactOpts, newOwner)
}

// WithdrawFromContract is a paid mutator transaction binding the contract method 0xf3ff4961.
//
// Solidity: function withdrawFromContract(receiver address) returns(bool)
func (_BaasROI *BaasROITransactor) WithdrawFromContract(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "withdrawFromContract", receiver)
}

// WithdrawFromContract is a paid mutator transaction binding the contract method 0xf3ff4961.
//
// Solidity: function withdrawFromContract(receiver address) returns(bool)
func (_BaasROI *BaasROISession) WithdrawFromContract(receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.WithdrawFromContract(&_BaasROI.TransactOpts, receiver)
}

// WithdrawFromContract is a paid mutator transaction binding the contract method 0xf3ff4961.
//
// Solidity: function withdrawFromContract(receiver address) returns(bool)
func (_BaasROI *BaasROITransactorSession) WithdrawFromContract(receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.WithdrawFromContract(&_BaasROI.TransactOpts, receiver)
}

// BaasROIOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasROI contract.
type BaasROIOwnershipTransferredIterator struct {
	Event *BaasROIOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasROIOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIOwnershipTransferred)
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
		it.Event = new(BaasROIOwnershipTransferred)
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
func (it *BaasROIOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIOwnershipTransferred represents a OwnershipTransferred event raised by the BaasROI contract.
type BaasROIOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasROI *BaasROIFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasROIOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasROIOwnershipTransferredIterator{contract: _BaasROI.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasROI *BaasROIFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasROIOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIOwnershipTransferred)
				if err := _BaasROI.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaasROIPaidOutIterator is returned from FilterPaidOut and is used to iterate over the raw logs and unpacked data for PaidOut events raised by the BaasROI contract.
type BaasROIPaidOutIterator struct {
	Event *BaasROIPaidOut // Event containing the contract specifics and raw log

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
func (it *BaasROIPaidOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIPaidOut)
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
		it.Event = new(BaasROIPaidOut)
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
func (it *BaasROIPaidOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIPaidOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIPaidOut represents a PaidOut event raised by the BaasROI contract.
type BaasROIPaidOut struct {
	Receiver        common.Address
	TokensReceived  *big.Int
	TokensPossessed *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaidOut is a free log retrieval operation binding the contract event 0x7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb.
//
// Solidity: e PaidOut(receiver indexed address, tokensReceived uint256, tokensPossessed uint256)
func (_BaasROI *BaasROIFilterer) FilterPaidOut(opts *bind.FilterOpts, receiver []common.Address) (*BaasROIPaidOutIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PaidOut", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BaasROIPaidOutIterator{contract: _BaasROI.contract, event: "PaidOut", logs: logs, sub: sub}, nil
}

// WatchPaidOut is a free log subscription operation binding the contract event 0x7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb.
//
// Solidity: e PaidOut(receiver indexed address, tokensReceived uint256, tokensPossessed uint256)
func (_BaasROI *BaasROIFilterer) WatchPaidOut(opts *bind.WatchOpts, sink chan<- *BaasROIPaidOut, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PaidOut", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIPaidOut)
				if err := _BaasROI.contract.UnpackLog(event, "PaidOut", log); err != nil {
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

// BaasROIPaidOutAllIterator is returned from FilterPaidOutAll and is used to iterate over the raw logs and unpacked data for PaidOutAll events raised by the BaasROI contract.
type BaasROIPaidOutAllIterator struct {
	Event *BaasROIPaidOutAll // Event containing the contract specifics and raw log

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
func (it *BaasROIPaidOutAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIPaidOutAll)
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
		it.Event = new(BaasROIPaidOutAll)
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
func (it *BaasROIPaidOutAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIPaidOutAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIPaidOutAll represents a PaidOutAll event raised by the BaasROI contract.
type BaasROIPaidOutAll struct {
	TokensProvided  *big.Int
	TokensPossessed *big.Int
	TokenHolders    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaidOutAll is a free log retrieval operation binding the contract event 0xda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9.
//
// Solidity: e PaidOutAll(tokensProvided uint256, tokensPossessed uint256, tokenHolders uint256)
func (_BaasROI *BaasROIFilterer) FilterPaidOutAll(opts *bind.FilterOpts) (*BaasROIPaidOutAllIterator, error) {

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PaidOutAll")
	if err != nil {
		return nil, err
	}
	return &BaasROIPaidOutAllIterator{contract: _BaasROI.contract, event: "PaidOutAll", logs: logs, sub: sub}, nil
}

// WatchPaidOutAll is a free log subscription operation binding the contract event 0xda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9.
//
// Solidity: e PaidOutAll(tokensProvided uint256, tokensPossessed uint256, tokenHolders uint256)
func (_BaasROI *BaasROIFilterer) WatchPaidOutAll(opts *bind.WatchOpts, sink chan<- *BaasROIPaidOutAll) (event.Subscription, error) {

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PaidOutAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIPaidOutAll)
				if err := _BaasROI.contract.UnpackLog(event, "PaidOutAll", log); err != nil {
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

// BaasROIPayoutOmittedIterator is returned from FilterPayoutOmitted and is used to iterate over the raw logs and unpacked data for PayoutOmitted events raised by the BaasROI contract.
type BaasROIPayoutOmittedIterator struct {
	Event *BaasROIPayoutOmitted // Event containing the contract specifics and raw log

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
func (it *BaasROIPayoutOmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIPayoutOmitted)
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
		it.Event = new(BaasROIPayoutOmitted)
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
func (it *BaasROIPayoutOmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIPayoutOmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIPayoutOmitted represents a PayoutOmitted event raised by the BaasROI contract.
type BaasROIPayoutOmitted struct {
	OmittedReceiver common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPayoutOmitted is a free log retrieval operation binding the contract event 0x1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb72.
//
// Solidity: e PayoutOmitted(omittedReceiver address)
func (_BaasROI *BaasROIFilterer) FilterPayoutOmitted(opts *bind.FilterOpts) (*BaasROIPayoutOmittedIterator, error) {

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PayoutOmitted")
	if err != nil {
		return nil, err
	}
	return &BaasROIPayoutOmittedIterator{contract: _BaasROI.contract, event: "PayoutOmitted", logs: logs, sub: sub}, nil
}

// WatchPayoutOmitted is a free log subscription operation binding the contract event 0x1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb72.
//
// Solidity: e PayoutOmitted(omittedReceiver address)
func (_BaasROI *BaasROIFilterer) WatchPayoutOmitted(opts *bind.WatchOpts, sink chan<- *BaasROIPayoutOmitted) (event.Subscription, error) {

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PayoutOmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIPayoutOmitted)
				if err := _BaasROI.contract.UnpackLog(event, "PayoutOmitted", log); err != nil {
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

// BaasROIWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BaasROI contract.
type BaasROIWithdrawnIterator struct {
	Event *BaasROIWithdrawn // Event containing the contract specifics and raw log

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
func (it *BaasROIWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasROIWithdrawn)
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
		it.Event = new(BaasROIWithdrawn)
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
func (it *BaasROIWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasROIWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasROIWithdrawn represents a Withdrawn event raised by the BaasROI contract.
type BaasROIWithdrawn struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(receiver indexed address, amount uint256)
func (_BaasROI *BaasROIFilterer) FilterWithdrawn(opts *bind.FilterOpts, receiver []common.Address) (*BaasROIWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "Withdrawn", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BaasROIWithdrawnIterator{contract: _BaasROI.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(receiver indexed address, amount uint256)
func (_BaasROI *BaasROIFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BaasROIWithdrawn, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "Withdrawn", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasROIWithdrawn)
				if err := _BaasROI.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
