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
const BaasROIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"omittedReceiver\",\"type\":\"address\"}],\"name\":\"PayoutOmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokensReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensPossessed\",\"type\":\"uint256\"}],\"name\":\"PaidOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokensProvided\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensPossessed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenHolders\",\"type\":\"uint256\"}],\"name\":\"PaidOutAll\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"payoutAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eligibleToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"hasEnoughTokensForPayout\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"uint256\"},{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"roi\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\"},{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"roiOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"currentPayoutObligation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"maxTokensToBeRewarded\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenEuroConversionRate\",\"type\":\"uint256\"}],\"name\":\"optimalPayoutDistribution\",\"outputs\":[{\"name\":\"maxTokenToBeRewarded\",\"type\":\"uint256\"},{\"name\":\"minPayoutBalance\",\"type\":\"uint256\"},{\"name\":\"error\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasROIBin is the compiled bytecode used for deploying new contracts.
const BaasROIBin = `0x608060405234801561001057600080fd5b5060405160208061179983398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611659806101406000396000f3006080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806301b32f0a1461010157806306fdde031461014c5780631473ada3146101dc5780631da2649b1461022b57806351cff8d9146102935780635f86d790146102ee578063715018a61461032f5780637c3a00fd146103465780638d10ca26146103715780638da5cb5b146103b65780638f32d59b1461040d5780639358928b1461043c5780639d76ea5814610467578063b69ef8a8146104be578063bb731467146104e9578063be4d0aa214610514578063daf3841914610555578063f2fde38b1461059a575b600080fd5b34801561010d57600080fd5b5061013660048036038101908080359060200190929190803590602001909291905050506105dd565b6040518082815260200191505060405180910390f35b34801561015857600080fd5b5061016161061f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101a1578082015181840152602081019050610186565b50505050905090810190601f1680156101ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101e857600080fd5b506102076004803603810190808035906020019092919050505061065c565b60405180848152602001838152602001828152602001935050505060405180910390f35b34801561023757600080fd5b50610276600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610714565b604051808381526020018281526020019250505060405180910390f35b34801561029f57600080fd5b506102d4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061082a565b604051808215151515815260200191505060405180910390f35b3480156102fa57600080fd5b506103196004803603810190808035906020019092919050505061096e565b6040518082815260200191505060405180910390f35b34801561033b57600080fd5b506103446109bb565b005b34801561035257600080fd5b5061035b610a8d565b6040518082815260200191505060405180910390f35b34801561037d57600080fd5b5061039c60048036038101908080359060200190929190505050610a96565b604051808215151515815260200191505060405180910390f35b3480156103c257600080fd5b506103cb610ac6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561041957600080fd5b50610422610aef565b604051808215151515815260200191505060405180910390f35b34801561044857600080fd5b50610451610b46565b6040518082815260200191505060405180910390f35b34801561047357600080fd5b5061047c610c0e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104ca57600080fd5b506104d3610c38565b6040518082815260200191505060405180910390f35b3480156104f557600080fd5b506104fe610d37565b6040518082815260200191505060405180910390f35b34801561052057600080fd5b5061053f60048036038101908080359060200190929190505050610e18565b6040518082815260200191505060405180910390f35b34801561056157600080fd5b5061058060048036038101908080359060200190929190505050610e32565b604051808215151515815260200191505060405180910390f35b3480156105a657600080fd5b506105db600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112ca565b005b6000808214156105f05760009050610619565b610616826106086009866112e990919063ffffffff16565b61132790919063ffffffff16565b90505b92915050565b60606040805190810160405280601481526020017f52455455524e204f4620494e564553544d454e54000000000000000000000000815250905090565b60008060008061067685600961135190919063ffffffff16565b90506106a5816106976009610689610b46565b6112e990919063ffffffff16565b61132790919063ffffffff16565b92506106d3816106c5876106b7610b46565b6112e990919063ffffffff16565b61132790919063ffffffff16565b9350610701846106f3856106e5610b46565b61137290919063ffffffff16565b61137290919063ffffffff16565b9150838383935093509350509193909250565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156107d657600080fd5b505af11580156107ea573d6000803e3d6000fd5b505050506040513d602081101561080057600080fd5b8101908080519060200190929190505050905061081d81856105dd565b8192509250509250929050565b600080610835610aef565b151561084057600080fd5b610848610c38565b905060008111151561085957600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561091e57600080fd5b505af1158015610932573d6000803e3d6000fd5b505050506040513d602081101561094857600080fd5b8101908080519060200190929190505050151561096457600080fd5b6001915050919050565b60008061098560098461135190919063ffffffff16565b90506109b3816109a585610997610b46565b6112e990919063ffffffff16565b61132790919063ffffffff16565b915050919050565b6109c3610aef565b15156109ce57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006009905090565b600080600080610aa58561065c565b809350819450829550505050610ab9610c38565b8211159350505050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639358928b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610bce57600080fd5b505af1158015610be2573d6000803e3d6000fd5b505050506040513d6020811015610bf857600080fd5b8101908080519060200190929190505050905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610cf757600080fd5b505af1158015610d0b573d6000803e3d6000fd5b505050506040513d6020811015610d2157600080fd5b8101908080519060200190929190505050905090565b6000610e13610d44610c38565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639358928b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610dca57600080fd5b505af1158015610dde573d6000803e3d6000fd5b505050506040513d6020811015610df457600080fd5b810190808051906020019092919050505061137290919063ffffffff16565b905090565b6000610e2b610e25610d37565b836105dd565b9050919050565b6000806000806000806060600080600080600080610e4e610aef565b1515610e5957600080fd5b610e628e610a96565b1515610e6d57600080fd5b610e75610c38565b9b50600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a654cfab6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401608060405180830381600087803b158015610efd57600080fd5b505af1158015610f11573d6000803e3d6000fd5b505050506040513d6080811015610f2757600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050809b50819c50829d50839e5050505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637af70c1f6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015610fec57600080fd5b505af1158015611000573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561102a57600080fd5b81019080805164010000000081111561104257600080fd5b8281019050602081018481111561105857600080fd5b815185602082028301116401000000008211171561107557600080fd5b50509291905050509650865195506000945060009350600092505b858310156112725786838151811015156110a657fe5b9060200190602002015191508a73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16148061111757508773ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b8061114d57508873ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b8061118357508973ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b806111b957503073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b15611226577f1ce49881e4d47088c220804148559e4103d5df83e4fc9e98101ff01eb74ecb7282604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1611265565b611247878481518110151561123757fe5b906020019060200201518f611393565b905061125c818661135190919063ffffffff16565b94506001840193505b8280600101935050611090565b7fda72e69b624518c6d2f93807ec73b6a8a638ae29f7f0a8fa353946027d72fba9858d8660405180848152602001838152602001828152602001935050505060405180910390a1505050505050505050505050919050565b6112d2610aef565b15156112dd57600080fd5b6112e681611533565b50565b60008060008414156112fe5760009150611320565b828402905082848281151561130f57fe5b0414151561131c57600080fd5b8091505b5092915050565b60008060008311151561133957600080fd5b828481151561134457fe5b0490508091505092915050565b600080828401905083811015151561136857600080fd5b8091505092915050565b60008083831115151561138457600080fd5b82840390508091505092915050565b60008060006113a28585610714565b8092508193505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561146f57600080fd5b505af1158015611483573d6000803e3d6000fd5b505050506040513d602081101561149957600080fd5b810190808051906020019092919050505015156114b557600080fd5b7f7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb858383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a1819250505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561156f57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a72305820a1bead558a03710576a3c742cccb52eb9a20756eb3e71555862940c3e1d858250029`

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

// HasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x8d10ca26.
//
// Solidity: function hasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROICaller) HasEnoughTokensForPayout(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "hasEnoughTokensForPayout", tokenEuroConversionRate)
	return *ret0, err
}

// HasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x8d10ca26.
//
// Solidity: function hasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROISession) HasEnoughTokensForPayout(tokenEuroConversionRate *big.Int) (bool, error) {
	return _BaasROI.Contract.HasEnoughTokensForPayout(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// HasEnoughTokensForPayout is a free data retrieval call binding the contract method 0x8d10ca26.
//
// Solidity: function hasEnoughTokensForPayout(tokenEuroConversionRate uint256) constant returns(bool)
func (_BaasROI *BaasROICallerSession) HasEnoughTokensForPayout(tokenEuroConversionRate *big.Int) (bool, error) {
	return _BaasROI.Contract.HasEnoughTokensForPayout(&_BaasROI.CallOpts, tokenEuroConversionRate)
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

// MaxTokensToBeRewarded is a free data retrieval call binding the contract method 0x5f86d790.
//
// Solidity: function maxTokensToBeRewarded(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICaller) MaxTokensToBeRewarded(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "maxTokensToBeRewarded", tokenEuroConversionRate)
	return *ret0, err
}

// MaxTokensToBeRewarded is a free data retrieval call binding the contract method 0x5f86d790.
//
// Solidity: function maxTokensToBeRewarded(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROISession) MaxTokensToBeRewarded(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.MaxTokensToBeRewarded(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// MaxTokensToBeRewarded is a free data retrieval call binding the contract method 0x5f86d790.
//
// Solidity: function maxTokensToBeRewarded(tokenEuroConversionRate uint256) constant returns(uint256)
func (_BaasROI *BaasROICallerSession) MaxTokensToBeRewarded(tokenEuroConversionRate *big.Int) (*big.Int, error) {
	return _BaasROI.Contract.MaxTokensToBeRewarded(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasROI *BaasROICaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaasROI.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasROI *BaasROISession) Name() (string, error) {
	return _BaasROI.Contract.Name(&_BaasROI.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasROI *BaasROICallerSession) Name() (string, error) {
	return _BaasROI.Contract.Name(&_BaasROI.CallOpts)
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(maxTokenToBeRewarded uint256, minPayoutBalance uint256, error uint256)
func (_BaasROI *BaasROICaller) OptimalPayoutDistribution(opts *bind.CallOpts, tokenEuroConversionRate *big.Int) (struct {
	MaxTokenToBeRewarded *big.Int
	MinPayoutBalance     *big.Int
	Error                *big.Int
}, error) {
	ret := new(struct {
		MaxTokenToBeRewarded *big.Int
		MinPayoutBalance     *big.Int
		Error                *big.Int
	})
	out := ret
	err := _BaasROI.contract.Call(opts, out, "optimalPayoutDistribution", tokenEuroConversionRate)
	return *ret, err
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(maxTokenToBeRewarded uint256, minPayoutBalance uint256, error uint256)
func (_BaasROI *BaasROISession) OptimalPayoutDistribution(tokenEuroConversionRate *big.Int) (struct {
	MaxTokenToBeRewarded *big.Int
	MinPayoutBalance     *big.Int
	Error                *big.Int
}, error) {
	return _BaasROI.Contract.OptimalPayoutDistribution(&_BaasROI.CallOpts, tokenEuroConversionRate)
}

// OptimalPayoutDistribution is a free data retrieval call binding the contract method 0x1473ada3.
//
// Solidity: function optimalPayoutDistribution(tokenEuroConversionRate uint256) constant returns(maxTokenToBeRewarded uint256, minPayoutBalance uint256, error uint256)
func (_BaasROI *BaasROICallerSession) OptimalPayoutDistribution(tokenEuroConversionRate *big.Int) (struct {
	MaxTokenToBeRewarded *big.Int
	MinPayoutBalance     *big.Int
	Error                *big.Int
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

// RoiOf is a free data retrieval call binding the contract method 0x1da2649b.
//
// Solidity: function roiOf(wallet address, tokenEuroConversionRate uint256) constant returns(uint256, uint256)
func (_BaasROI *BaasROICaller) RoiOf(opts *bind.CallOpts, wallet common.Address, tokenEuroConversionRate *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BaasROI.contract.Call(opts, out, "roiOf", wallet, tokenEuroConversionRate)
	return *ret0, *ret1, err
}

// RoiOf is a free data retrieval call binding the contract method 0x1da2649b.
//
// Solidity: function roiOf(wallet address, tokenEuroConversionRate uint256) constant returns(uint256, uint256)
func (_BaasROI *BaasROISession) RoiOf(wallet common.Address, tokenEuroConversionRate *big.Int) (*big.Int, *big.Int, error) {
	return _BaasROI.Contract.RoiOf(&_BaasROI.CallOpts, wallet, tokenEuroConversionRate)
}

// RoiOf is a free data retrieval call binding the contract method 0x1da2649b.
//
// Solidity: function roiOf(wallet address, tokenEuroConversionRate uint256) constant returns(uint256, uint256)
func (_BaasROI *BaasROICallerSession) RoiOf(wallet common.Address, tokenEuroConversionRate *big.Int) (*big.Int, *big.Int, error) {
	return _BaasROI.Contract.RoiOf(&_BaasROI.CallOpts, wallet, tokenEuroConversionRate)
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

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(receiver address) returns(bool)
func (_BaasROI *BaasROITransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.contract.Transact(opts, "withdraw", receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(receiver address) returns(bool)
func (_BaasROI *BaasROISession) Withdraw(receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.Withdraw(&_BaasROI.TransactOpts, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(receiver address) returns(bool)
func (_BaasROI *BaasROITransactorSession) Withdraw(receiver common.Address) (*types.Transaction, error) {
	return _BaasROI.Contract.Withdraw(&_BaasROI.TransactOpts, receiver)
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
// Solidity: e PaidOut(receiver address, tokensReceived uint256, tokensPossessed uint256)
func (_BaasROI *BaasROIFilterer) FilterPaidOut(opts *bind.FilterOpts) (*BaasROIPaidOutIterator, error) {

	logs, sub, err := _BaasROI.contract.FilterLogs(opts, "PaidOut")
	if err != nil {
		return nil, err
	}
	return &BaasROIPaidOutIterator{contract: _BaasROI.contract, event: "PaidOut", logs: logs, sub: sub}, nil
}

// WatchPaidOut is a free log subscription operation binding the contract event 0x7ca7469714f3e1d8732b3a67b0599fba3be82b826137fcfa805c19afc2b20aeb.
//
// Solidity: e PaidOut(receiver address, tokensReceived uint256, tokensPossessed uint256)
func (_BaasROI *BaasROIFilterer) WatchPaidOut(opts *bind.WatchOpts, sink chan<- *BaasROIPaidOut) (event.Subscription, error) {

	logs, sub, err := _BaasROI.contract.WatchLogs(opts, "PaidOut")
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
