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
const BaasIncentivesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncentiveProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Payout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"tokenNotDelivered\",\"type\":\"uint256\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"SetupCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"reserve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesProvided\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentives\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getIncentive\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InitialSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasIncentivesBin is the compiled bytecode used for deploying new contracts.
const BaasIncentivesBin = `0x60806040526000600560006101000a81548160ff02191690831515021790555034801561002b57600080fd5b5060405160208061185c83398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506117018061015b6000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100f65780630d5df7ba14610186578063117de2fd146101f25780631eb25d1314610253578063392e53cd1461027e578063452869d6146102ad5780635742a36e14610320578063715018a61461034b57806374a8f103146103625780638da5cb5b146103bd5780638f32d59b146104145780638f70c2d8146104435780639d76ea58146104bb578063ab77e2d314610512578063b69ef8a81461053d578063ba0bba4014610568578063f2fde38b1461057f575b600080fd5b34801561010257600080fd5b5061010b6105c2565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014b578082015181840152602081019050610130565b50505050905090810190601f1680156101785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019257600080fd5b5061019b6105ff565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101de5780820151818401526020810190506101c3565b505050509050019250505060405180910390f35b3480156101fe57600080fd5b5061023d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061068d565b6040518082815260200191505060405180910390f35b34801561025f57600080fd5b50610268610ab4565b6040518082815260200191505060405180910390f35b34801561028a57600080fd5b50610293610ac7565b604051808215151515815260200191505060405180910390f35b3480156102b957600080fd5b50610306600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035600019169060200190929190505050610ade565b604051808215151515815260200191505060405180910390f35b34801561032c57600080fd5b50610335610e6f565b6040518082815260200191505060405180910390f35b34801561035757600080fd5b50610360610e79565b005b34801561036e57600080fd5b506103a3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f4b565b604051808215151515815260200191505060405180910390f35b3480156103c957600080fd5b506103d26111de565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561042057600080fd5b50610429611207565b604051808215151515815260200191505060405180910390f35b34801561044f57600080fd5b50610484600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061125e565b6040518085815260200184815260200183600019166000191681526020018215151515815260200194505050505060405180910390f35b3480156104c757600080fd5b506104d0611374565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561051e57600080fd5b5061052761139e565b6040518082815260200191505060405180910390f35b34801561054957600080fd5b506105526113a8565b6040518082815260200191505060405180910390f35b34801561057457600080fd5b5061057d6114a7565b005b34801561058b57600080fd5b506105c0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061152f565b005b60606040805190810160405280600a81526020017f494e43454e544956455300000000000000000000000000000000000000000000815250905090565b6060600280548060200260200160405190810160405280929190818152602001828054801561068357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610639575b5050505050905090565b600080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff161515610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6164647265737320776173206e6576657220696e63656e746976697a6564000081525060200191505060405180910390fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101540390508083101515610854576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6e6f7420656e6f75676820746f6b656e206c656674000000000000000000000081525060200191505060405180910390fd5b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb30856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561091957600080fd5b505af115801561092d573d6000803e3d6000fd5b505050506040513d602081101561094357600080fd5b810190808051906020019092919050505015156109c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b82600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008282540192505081905550600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154600019168473ffffffffffffffffffffffffffffffffffffffff167f212950d8b6a2784adeaa8e088d070610bdfab9eaef5a7127d90118ec1cf34a61856040518082815260200191505060405180910390a35092915050565b60006a084595161401484a000000905090565b6000600560009054906101000a900460ff16905090565b6000610ae8611207565b1515610af357600080fd5b600560009054906101000a900460ff161515610b77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f636f6e7472616374206d75737420626520696e697469616c697a65640000000081525060200191505060405180910390fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff16151515610c62576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001807f616464726573732073686f756c64206e6f74206265656e20696e63656e74697681526020017f697a65642079657400000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60a0604051908101604052808573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001600081526020018360001916815260200160011515815250600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003019060001916905560808201518160040160006101000a81548160ff02191690831515021790555090505060028490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610df08360035461154e90919063ffffffff16565b600381905550610e0b8360045461156f90919063ffffffff16565b60048190555081600019168473ffffffffffffffffffffffffffffffffffffffff167f8c5cf7bdeff0f0e688b7f7ff9a9d4be7c19cb375b45a8e03e2b1f130cc6f3d6a856040518082815260200191505060405180910390a3600190509392505050565b6000600454905090565b610e81611207565b1515610e8c57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080610f56611207565b1515610f6157600080fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160009054906101000a900460ff161515611025576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f6164647265737320776173206e6576657220696e63656e746976697a6564000081525060200191505060405180910390fd5b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160006101000a81548160ff021916908315150217905550600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020154600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015403905061111e8160035461156f90919063ffffffff16565b6003819055506111398160045461154e90919063ffffffff16565b600481905550600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030154600019168373ffffffffffffffffffffffffffffffffffffffff167fe490c408aa61d6a8f6ab6dd0b3722b31d05c4ac6dbb268c4472052fb8d89d8df836040518082815260200191505060405180910390a36001915050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60008060008061126c61168a565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060a060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820154600019166000191681526020016004820160009054906101000a900460ff161515151581525050905080602001518160400151826060015183608001519450945094509450509193509193565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600354905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561146757600080fd5b505af115801561147b573d6000803e3d6000fd5b505050506040513d602081101561149157600080fd5b8101908080519060200190929190505050905090565b600560009054906101000a900460ff161515156114c357600080fd5b6114cb6113a8565b60038190555060006004819055506001600560006101000a81548160ff0219169083151502179055507fd406aba267fc52b718da0f82dc70773ae07bd549015dbd3fe6858094443b95326003546040518082815260200191505060405180910390a1565b611537611207565b151561154257600080fd5b61154b81611590565b50565b60008083831115151561156057600080fd5b82840390508091505092915050565b600080828401905083811015151561158657600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156115cc57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60a060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000801916815260200160001515815250905600a165627a7a723058205cf6291180ba7350b9ad2e5eb5ecc1232e659a37f4f5cfc8fa2cd0f65b9414a40029`

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
// Solidity: function getIncentive(account address) constant returns(amount uint256, provided uint256, id bytes32, isValue bool)
func (_BaasIncentives *BaasIncentivesCaller) GetIncentive(opts *bind.CallOpts, account common.Address) (struct {
	Amount   *big.Int
	Provided *big.Int
	Id       [32]byte
	IsValue  bool
}, error) {
	ret := new(struct {
		Amount   *big.Int
		Provided *big.Int
		Id       [32]byte
		IsValue  bool
	})
	out := ret
	err := _BaasIncentives.contract.Call(opts, out, "getIncentive", account)
	return *ret, err
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(amount uint256, provided uint256, id bytes32, isValue bool)
func (_BaasIncentives *BaasIncentivesSession) GetIncentive(account common.Address) (struct {
	Amount   *big.Int
	Provided *big.Int
	Id       [32]byte
	IsValue  bool
}, error) {
	return _BaasIncentives.Contract.GetIncentive(&_BaasIncentives.CallOpts, account)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(amount uint256, provided uint256, id bytes32, isValue bool)
func (_BaasIncentives *BaasIncentivesCallerSession) GetIncentive(account common.Address) (struct {
	Amount   *big.Int
	Provided *big.Int
	Id       [32]byte
	IsValue  bool
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

// Payout is a paid mutator transaction binding the contract method 0x117de2fd.
//
// Solidity: function payout(account address, amount uint256) returns(uint256)
func (_BaasIncentives *BaasIncentivesTransactor) Payout(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "payout", account, amount)
}

// Payout is a paid mutator transaction binding the contract method 0x117de2fd.
//
// Solidity: function payout(account address, amount uint256) returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) Payout(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Payout(&_BaasIncentives.TransactOpts, account, amount)
}

// Payout is a paid mutator transaction binding the contract method 0x117de2fd.
//
// Solidity: function payout(account address, amount uint256) returns(uint256)
func (_BaasIncentives *BaasIncentivesTransactorSession) Payout(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Payout(&_BaasIncentives.TransactOpts, account, amount)
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

// Reserve is a paid mutator transaction binding the contract method 0x452869d6.
//
// Solidity: function reserve(account address, amount uint256, id bytes32) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Reserve(opts *bind.TransactOpts, account common.Address, amount *big.Int, id [32]byte) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "reserve", account, amount, id)
}

// Reserve is a paid mutator transaction binding the contract method 0x452869d6.
//
// Solidity: function reserve(account address, amount uint256, id bytes32) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Reserve(account common.Address, amount *big.Int, id [32]byte) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Reserve(&_BaasIncentives.TransactOpts, account, amount, id)
}

// Reserve is a paid mutator transaction binding the contract method 0x452869d6.
//
// Solidity: function reserve(account address, amount uint256, id bytes32) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Reserve(account common.Address, amount *big.Int, id [32]byte) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Reserve(&_BaasIncentives.TransactOpts, account, amount, id)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(account address) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Revoke(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "revoke", account)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(account address) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Revoke(account common.Address) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Revoke(&_BaasIncentives.TransactOpts, account)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(account address) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Revoke(account common.Address) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Revoke(&_BaasIncentives.TransactOpts, account)
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_BaasIncentives *BaasIncentivesTransactor) Setup(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "setup")
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_BaasIncentives *BaasIncentivesSession) Setup() (*types.Transaction, error) {
	return _BaasIncentives.Contract.Setup(&_BaasIncentives.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0xba0bba40.
//
// Solidity: function setup() returns()
func (_BaasIncentives *BaasIncentivesTransactorSession) Setup() (*types.Transaction, error) {
	return _BaasIncentives.Contract.Setup(&_BaasIncentives.TransactOpts)
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
	Id      [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentiveProvided is a free log retrieval operation binding the contract event 0x8c5cf7bdeff0f0e688b7f7ff9a9d4be7c19cb375b45a8e03e2b1f130cc6f3d6a.
//
// Solidity: e IncentiveProvided(account indexed address, id indexed bytes32, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterIncentiveProvided(opts *bind.FilterOpts, account []common.Address, id [][32]byte) (*BaasIncentivesIncentiveProvidedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "IncentiveProvided", accountRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesIncentiveProvidedIterator{contract: _BaasIncentives.contract, event: "IncentiveProvided", logs: logs, sub: sub}, nil
}

// WatchIncentiveProvided is a free log subscription operation binding the contract event 0x8c5cf7bdeff0f0e688b7f7ff9a9d4be7c19cb375b45a8e03e2b1f130cc6f3d6a.
//
// Solidity: e IncentiveProvided(account indexed address, id indexed bytes32, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchIncentiveProvided(opts *bind.WatchOpts, sink chan<- *BaasIncentivesIncentiveProvided, account []common.Address, id [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "IncentiveProvided", accountRule, idRule)
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

// BaasIncentivesSetupCompletedIterator is returned from FilterSetupCompleted and is used to iterate over the raw logs and unpacked data for SetupCompleted events raised by the BaasIncentives contract.
type BaasIncentivesSetupCompletedIterator struct {
	Event *BaasIncentivesSetupCompleted // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesSetupCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesSetupCompleted)
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
		it.Event = new(BaasIncentivesSetupCompleted)
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
func (it *BaasIncentivesSetupCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesSetupCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesSetupCompleted represents a SetupCompleted event raised by the BaasIncentives contract.
type BaasIncentivesSetupCompleted struct {
	Supply *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetupCompleted is a free log retrieval operation binding the contract event 0xd406aba267fc52b718da0f82dc70773ae07bd549015dbd3fe6858094443b9532.
//
// Solidity: e SetupCompleted(supply uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterSetupCompleted(opts *bind.FilterOpts) (*BaasIncentivesSetupCompletedIterator, error) {

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesSetupCompletedIterator{contract: _BaasIncentives.contract, event: "SetupCompleted", logs: logs, sub: sub}, nil
}

// WatchSetupCompleted is a free log subscription operation binding the contract event 0xd406aba267fc52b718da0f82dc70773ae07bd549015dbd3fe6858094443b9532.
//
// Solidity: e SetupCompleted(supply uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchSetupCompleted(opts *bind.WatchOpts, sink chan<- *BaasIncentivesSetupCompleted) (event.Subscription, error) {

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesSetupCompleted)
				if err := _BaasIncentives.contract.UnpackLog(event, "SetupCompleted", log); err != nil {
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
