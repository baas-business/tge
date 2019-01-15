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
const BaasEscrowABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x715018a6\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8da5cb5b\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8f32d59b\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf2fde38b\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CapitalRaised\",\"type\":\"event\",\"signature\":\"0xbc65650fcb1dd5a9e3c92e066533c3425c66a51f1f60fbfe11ce352452c9b08d\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vestingStartBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingEndBlock\",\"type\":\"uint256\"}],\"name\":\"SetupCompleted\",\"type\":\"event\",\"signature\":\"0x05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd099\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenPrice\",\"type\":\"uint256\"}],\"name\":\"TokenDelivered\",\"type\":\"event\",\"signature\":\"0x80438a1ee44a1e699790add4b8fddef1ff8532d2fdb33aa686c90b79bb268478\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"signature\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\"},{\"constant\":false,\"inputs\":[{\"name\":\"vestingPeriod\",\"type\":\"uint256\"}],\"name\":\"setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x4313b9e5\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"raiseCapital\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xfbd207ed\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"conversionRate\",\"type\":\"uint256\"},{\"name\":\"capitalRaiseId\",\"type\":\"uint256\"}],\"name\":\"provideToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x0685df17\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xb69ef8a8\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x9d76ea58\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x392e53cd\"},{\"constant\":true,\"inputs\":[{\"name\":\"capitalRaiseId\",\"type\":\"uint256\"}],\"name\":\"capitalRaise\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x4f4a7a79\"},{\"constant\":true,\"inputs\":[{\"name\":\"capitalRaiseId\",\"type\":\"uint256\"}],\"name\":\"whiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x5c475d42\"},{\"constant\":true,\"inputs\":[{\"name\":\"capitalRaiseId\",\"type\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"whiteListedAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x0ed868ac\"}]"

// BaasEscrowBin is the compiled bytecode used for deploying new contracts.
const BaasEscrowBin = `0x60806040526000600860006101000a81548160ff02191690831515021790555034801561002b57600080fd5b506040516020806116dd8339810180604052602081101561004b57600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611574806101696000396000f3fe6080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630685df17146100ca5780630ed868ac14610151578063392e53cd146101c05780634313b9e5146101ef5780634f4a7a79146102425780635c475d42146102b1578063715018a6146103415780638da5cb5b146103585780638f32d59b146103af5780639d76ea58146103de578063b69ef8a814610435578063f2fde38b14610460578063fbd207ed146104b1575b600080fd5b3480156100d657600080fd5b50610137600480360360808110156100ed57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019092919050505061050e565b604051808215151515815260200191505060405180910390f35b34801561015d57600080fd5b506101aa6004803603604081101561017457600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108ba565b6040518082815260200191505060405180910390f35b3480156101cc57600080fd5b506101d5610915565b604051808215151515815260200191505060405180910390f35b3480156101fb57600080fd5b506102286004803603602081101561021257600080fd5b810190808035906020019092919050505061092c565b604051808215151515815260200191505060405180910390f35b34801561024e57600080fd5b5061027b6004803603602081101561026557600080fd5b81019080803590602001909291905050506109d0565b60405180868152602001858152602001848152602001838152602001821515151581526020019550505050505060405180910390f35b3480156102bd57600080fd5b506102ea600480360360208110156102d457600080fd5b8101908080359060200190929190505050610a71565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561032d578082015181840152602081019050610312565b505050509050019250505060405180910390f35b34801561034d57600080fd5b50610356610b12565b005b34801561036457600080fd5b5061036d610be4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103bb57600080fd5b506103c4610c0d565b604051808215151515815260200191505060405180910390f35b3480156103ea57600080fd5b506103f3610c64565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561044157600080fd5b5061044a610c8e565b6040518082815260200191505060405180910390f35b34801561046c57600080fd5b506104af6004803603602081101561048357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d8b565b005b3480156104bd57600080fd5b506104f4600480360360408110156104d457600080fd5b810190808035906020019092919080359060200190929190505050610daa565b604051808215151515815260200191505060405180910390f35b6000610518610c0d565b151561052357600080fd5b61052b611516565b6005600084815260200190815260200160002060a06040519081016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050905080608001511515610608576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f77726f6e67206361706974616c2072616973652069640000000000000000000081525060200191505060405180910390fd5b610623816040015182602001516113d990919063ffffffff16565b851015156106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f616d6f756e742073686f756c64206265206c657373207468656e20776861742081526020017f6973206c6566740000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb87876040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561078457600080fd5b505af1158015610798573d6000803e3d6000fd5b505050506040513d60208110156107ae57600080fd5b81019080805190602001909291905050501515610833576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f7472616e73666572206661696c6564000000000000000000000000000000000081525060200191505060405180910390fd5b8460056000858152602001908152602001600020600201600082825401925050819055508573ffffffffffffffffffffffffffffffffffffffff167f80438a1ee44a1e699790add4b8fddef1ff8532d2fdb33aa686c90b79bb2684788686604051808381526020018281526020019250505060405180910390a26001915050949350505050565b60006006600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600860009054906101000a900460ff16905090565b6000610936610c0d565b151561094157600080fd5b43600281905550816003819055506001600860006101000a81548160ff0219169083151502179055507f05e6f0d652e49955317590bd949b5d7600bc573d84320a97c0fec4dc5cdbd0996002546003546109a86003546002546113fb90919063ffffffff16565b60405180848152602001838152602001828152602001935050505060405180910390a1919050565b60008060008060006109e0611516565b6005600088815260200190815260200160002060a06040519081016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff161515151581525050905080600001518160200151826040015183606001518460800151955095509550955095505091939590929450565b606060076000838152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610b0657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610abc575b50505050509050919050565b610b1a610c0d565b1515610b2557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610d4b57600080fd5b505afa158015610d5f573d6000803e3d6000fd5b505050506040513d6020811015610d7557600080fd5b8101908080519060200190929190505050905090565b610d93610c0d565b1515610d9e57600080fd5b610da78161141c565b50565b6000610db4610c0d565b1515610dbf57600080fd5b610ddf6004546a31a17e847807b1bc0000006113d990919063ffffffff16565b8311151515610ded57600080fd5b6005600083815260200190815260200160002060040160009054906101000a900460ff16151515610e1d57600080fd5b60a0604051908101604052808381526020018481526020016000815260200143815260200160011515815250600560008481526020019081526020016000206000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a81548160ff021916908315150217905550905050600080600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a654cfab6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160806040518083038186803b158015610f3157600080fd5b505afa158015610f45573d6000803e3d6000fd5b505050506040513d6080811015610f5b57600080fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050809450819550829650839750505050506060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637af70c1f6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561102257600080fd5b505af1158015611036573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561106057600080fd5b81019080805164010000000081111561107857600080fd5b8281019050602081018481111561108e57600080fd5b81518560208202830111640100000000821117156110ab57600080fd5b5050929190505050905060008151905060008090505b8181101561139057600083828151811015156110d957fe5b9060200190602002015190508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16148061114a57508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b8061118057508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b806111b657508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156111c057611382565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231826040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561127b57600080fd5b505afa15801561128f573d6000803e3d6000fd5b505050506040513d60208110156112a557600080fd5b8101908080519060200190929190505050600660008c815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600760008b81526020019081526020016000208190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b5080806001019150506110c1565b50877fbc65650fcb1dd5a9e3c92e066533c3425c66a51f1f60fbfe11ce352452c9b08d8a6040518082815260200191505060405180910390a26001965050505050505092915050565b60008282111515156113ea57600080fd5b600082840390508091505092915050565b600080828401905083811015151561141257600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561145857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60a06040519081016040528060008152602001600081526020016000815260200160008152602001600015158152509056fea165627a7a723058200dd52d6b758ea7812d4866206fb43e670ee69005559706d9138297c29bd2d9300029`

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

// CapitalRaise is a free data retrieval call binding the contract method 0x4f4a7a79.
//
// Solidity: function capitalRaise(capitalRaiseId uint256) constant returns(id uint256, amount uint256, provided uint256, atBlock uint256, isValue bool)
func (_BaasEscrow *BaasEscrowCaller) CapitalRaise(opts *bind.CallOpts, capitalRaiseId *big.Int) (struct {
	Id       *big.Int
	Amount   *big.Int
	Provided *big.Int
	AtBlock  *big.Int
	IsValue  bool
}, error) {
	ret := new(struct {
		Id       *big.Int
		Amount   *big.Int
		Provided *big.Int
		AtBlock  *big.Int
		IsValue  bool
	})
	out := ret
	err := _BaasEscrow.contract.Call(opts, out, "capitalRaise", capitalRaiseId)
	return *ret, err
}

// CapitalRaise is a free data retrieval call binding the contract method 0x4f4a7a79.
//
// Solidity: function capitalRaise(capitalRaiseId uint256) constant returns(id uint256, amount uint256, provided uint256, atBlock uint256, isValue bool)
func (_BaasEscrow *BaasEscrowSession) CapitalRaise(capitalRaiseId *big.Int) (struct {
	Id       *big.Int
	Amount   *big.Int
	Provided *big.Int
	AtBlock  *big.Int
	IsValue  bool
}, error) {
	return _BaasEscrow.Contract.CapitalRaise(&_BaasEscrow.CallOpts, capitalRaiseId)
}

// CapitalRaise is a free data retrieval call binding the contract method 0x4f4a7a79.
//
// Solidity: function capitalRaise(capitalRaiseId uint256) constant returns(id uint256, amount uint256, provided uint256, atBlock uint256, isValue bool)
func (_BaasEscrow *BaasEscrowCallerSession) CapitalRaise(capitalRaiseId *big.Int) (struct {
	Id       *big.Int
	Amount   *big.Int
	Provided *big.Int
	AtBlock  *big.Int
	IsValue  bool
}, error) {
	return _BaasEscrow.Contract.CapitalRaise(&_BaasEscrow.CallOpts, capitalRaiseId)
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

// WhiteList is a free data retrieval call binding the contract method 0x5c475d42.
//
// Solidity: function whiteList(capitalRaiseId uint256) constant returns(address[])
func (_BaasEscrow *BaasEscrowCaller) WhiteList(opts *bind.CallOpts, capitalRaiseId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "whiteList", capitalRaiseId)
	return *ret0, err
}

// WhiteList is a free data retrieval call binding the contract method 0x5c475d42.
//
// Solidity: function whiteList(capitalRaiseId uint256) constant returns(address[])
func (_BaasEscrow *BaasEscrowSession) WhiteList(capitalRaiseId *big.Int) ([]common.Address, error) {
	return _BaasEscrow.Contract.WhiteList(&_BaasEscrow.CallOpts, capitalRaiseId)
}

// WhiteList is a free data retrieval call binding the contract method 0x5c475d42.
//
// Solidity: function whiteList(capitalRaiseId uint256) constant returns(address[])
func (_BaasEscrow *BaasEscrowCallerSession) WhiteList(capitalRaiseId *big.Int) ([]common.Address, error) {
	return _BaasEscrow.Contract.WhiteList(&_BaasEscrow.CallOpts, capitalRaiseId)
}

// WhiteListedAccount is a free data retrieval call binding the contract method 0x0ed868ac.
//
// Solidity: function whiteListedAccount(capitalRaiseId uint256, account address) constant returns(uint256)
func (_BaasEscrow *BaasEscrowCaller) WhiteListedAccount(opts *bind.CallOpts, capitalRaiseId *big.Int, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasEscrow.contract.Call(opts, out, "whiteListedAccount", capitalRaiseId, account)
	return *ret0, err
}

// WhiteListedAccount is a free data retrieval call binding the contract method 0x0ed868ac.
//
// Solidity: function whiteListedAccount(capitalRaiseId uint256, account address) constant returns(uint256)
func (_BaasEscrow *BaasEscrowSession) WhiteListedAccount(capitalRaiseId *big.Int, account common.Address) (*big.Int, error) {
	return _BaasEscrow.Contract.WhiteListedAccount(&_BaasEscrow.CallOpts, capitalRaiseId, account)
}

// WhiteListedAccount is a free data retrieval call binding the contract method 0x0ed868ac.
//
// Solidity: function whiteListedAccount(capitalRaiseId uint256, account address) constant returns(uint256)
func (_BaasEscrow *BaasEscrowCallerSession) WhiteListedAccount(capitalRaiseId *big.Int, account common.Address) (*big.Int, error) {
	return _BaasEscrow.Contract.WhiteListedAccount(&_BaasEscrow.CallOpts, capitalRaiseId, account)
}

// ProvideToken is a paid mutator transaction binding the contract method 0x0685df17.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint256, capitalRaiseId uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactor) ProvideToken(opts *bind.TransactOpts, account common.Address, amount *big.Int, conversionRate *big.Int, capitalRaiseId *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.contract.Transact(opts, "provideToken", account, amount, conversionRate, capitalRaiseId)
}

// ProvideToken is a paid mutator transaction binding the contract method 0x0685df17.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint256, capitalRaiseId uint256) returns(bool)
func (_BaasEscrow *BaasEscrowSession) ProvideToken(account common.Address, amount *big.Int, conversionRate *big.Int, capitalRaiseId *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.ProvideToken(&_BaasEscrow.TransactOpts, account, amount, conversionRate, capitalRaiseId)
}

// ProvideToken is a paid mutator transaction binding the contract method 0x0685df17.
//
// Solidity: function provideToken(account address, amount uint256, conversionRate uint256, capitalRaiseId uint256) returns(bool)
func (_BaasEscrow *BaasEscrowTransactorSession) ProvideToken(account common.Address, amount *big.Int, conversionRate *big.Int, capitalRaiseId *big.Int) (*types.Transaction, error) {
	return _BaasEscrow.Contract.ProvideToken(&_BaasEscrow.TransactOpts, account, amount, conversionRate, capitalRaiseId)
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
