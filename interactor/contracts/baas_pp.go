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

// BaasPPABI is the input ABI used to generate the binding from.
const BaasPPABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"discountType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"Finalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"discountType\",\"type\":\"uint8\"}],\"name\":\"issue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"discountType\",\"type\":\"uint8\"}],\"name\":\"tokensIssued\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"discountType\",\"type\":\"uint256\"}],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasPPBin is the compiled bytecode used for deploying new contracts.
const BaasPPBin = `0x608060405234801561001057600080fd5b506040516020806112bb8339810180604052602081101561003057600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061116d8061014e6000396000f3fe6080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806326215f25146100b45780634bb278f3146101345780635c5e1d5214610163578063715018a6146101b55780638d4e4083146101cc5780638da5cb5b146101fb5780638f32d59b146102525780639d76ea5814610281578063b69ef8a8146102d8578063f2fde38b14610303578063ff2ad8e414610354575b600080fd5b3480156100c057600080fd5b5061011a600480360360608110156100d757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803560ff1690602001909291905050506103a3565b604051808215151515815260200191505060405180910390f35b34801561014057600080fd5b50610149610552565b604051808215151515815260200191505060405180910390f35b34801561016f57600080fd5b5061019f6004803603602081101561018657600080fd5b81019080803560ff1690602001909291905050506107c6565b6040518082815260200191505060405180910390f35b3480156101c157600080fd5b506101ca610830565b005b3480156101d857600080fd5b506101e1610902565b604051808215151515815260200191505060405180910390f35b34801561020757600080fd5b50610210610919565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025e57600080fd5b50610267610942565b604051808215151515815260200191505060405180910390f35b34801561028d57600080fd5b50610296610999565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102e457600080fd5b506102ed6109c3565b6040518082815260200191505060405180910390f35b34801561030f57600080fd5b506103526004803603602081101561032657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ac0565b005b34801561036057600080fd5b5061038d6004803603602081101561037757600080fd5b8101908080359060200190929190505050610adf565b6040518082815260200191505060405180910390f35b60006103ad610942565b15156103b857600080fd5b600460009054906101000a900460ff1615151561043d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f636f6e74726163742077617320616c72656164792066696e616c697a6564000081525060200191505060405180910390fd5b600160ff168260ff1614806104585750600260ff168260ff16145b15156104cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f62616420646973636f756e74207479706520706173736564000000000000000081525060200191505060405180910390fd5b600160ff168260ff1614156104ea576104e58484610b64565b6104f5565b6104f48484610db4565b5b8160ff168473ffffffffffffffffffffffffffffffffffffffff167f064e268fa97e769a5abed875cfd116cb13d34b361b6088e9213df29cc7ea85e4856040518082815260200191505060405180910390a3600190509392505050565b600061055c610942565b151561056757600080fd5b600460009054906101000a900460ff161515156105ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f636f6e74726163742077617320616c72656164792066696e616c697a6564000081525060200191505060405180910390fd5b60006105f66109c3565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634cca31b630836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156106bd57600080fd5b505af11580156106d1573d6000803e3d6000fd5b505050506040513d60208110156106e757600080fd5b8101908080519060200190929190505050151561076c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6275726e696e67206f6620746f6b656e73206661696c6564000000000000000081525060200191505060405180910390fd5b6001600460006101000a81548160ff0219169083151502179055507f839cf22e1ba87ce2f5b9bbf46cf0175a09eed52febdfaac8852478e68203c763816040518082815260200191505060405180910390a1600191505090565b60008060ff168260ff1614156107f4576107ed60035460025461100490919063ffffffff16565b905061082b565b600160ff168260ff16141561080d57600254905061082b565b600260ff168260ff16141561082657600354905061082b565b600090505b919050565b610838610942565b151561084357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600460009054906101000a900460ff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610a8057600080fd5b505afa158015610a94573d6000803e3d6000fd5b505050506040513d6020811015610aaa57600080fd5b8101908080519060200190929190505050905090565b610ac8610942565b1515610ad357600080fd5b610adc81611025565b50565b60008060ff16821415610b1c57610b156a0422ca8b0a00a4250000006a0c685fa11e01ec6f00000061100490919063ffffffff16565b9050610b5f565b600160ff16821415610b3b576a0422ca8b0a00a4250000009050610b5f565b600260ff16821415610b5a576a0c685fa11e01ec6f0000009050610b5f565b600090505b919050565b610b846002546a0422ca8b0a00a42500000061111f90919063ffffffff16565b8111151515610c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b8152602001807f616d6f756e742065786365656473206861726420636170206f6620646973636f81526020017f756e74656420746f6b656e00000000000000000000000000000000000000000081525060400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050506040513d6020811015610d1057600080fd5b81019080805190602001909291905050501515610d95576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b610daa8160025461100490919063ffffffff16565b6002819055505050565b610dd46003546a0c685fa11e01ec6f00000061111f90919063ffffffff16565b8111151515610e71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f616d6f756e742065786365656473206861726420636170206f66206e6f74206481526020017f6973636f756e74656420746f6b656e000000000000000000000000000000000081525060400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610f3657600080fd5b505af1158015610f4a573d6000803e3d6000fd5b505050506040513d6020811015610f6057600080fd5b81019080805190602001909291905050501515610fe5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f746f6b656e207472616e73666572206661696c6564000000000000000000000081525060200191505060405180910390fd5b610ffa8160035461100490919063ffffffff16565b6003819055505050565b600080828401905083811015151561101b57600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561106157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600082821115151561113057600080fd5b60008284039050809150509291505056fea165627a7a72305820c0db296bb431e8a2f3fea470eb7fb9b61e597359cf86a8355e0ab97ec0d45aeb0029`

// DeployBaasPP deploys a new Ethereum contract, binding an instance of BaasPP to it.
func DeployBaasPP(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *BaasPP, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasPPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasPPBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasPP{BaasPPCaller: BaasPPCaller{contract: contract}, BaasPPTransactor: BaasPPTransactor{contract: contract}, BaasPPFilterer: BaasPPFilterer{contract: contract}}, nil
}

// BaasPP is an auto generated Go binding around an Ethereum contract.
type BaasPP struct {
	BaasPPCaller     // Read-only binding to the contract
	BaasPPTransactor // Write-only binding to the contract
	BaasPPFilterer   // Log filterer for contract events
}

// BaasPPCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasPPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasPPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasPPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasPPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasPPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasPPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasPPSession struct {
	Contract     *BaasPP           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasPPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasPPCallerSession struct {
	Contract *BaasPPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BaasPPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasPPTransactorSession struct {
	Contract     *BaasPPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasPPRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasPPRaw struct {
	Contract *BaasPP // Generic contract binding to access the raw methods on
}

// BaasPPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasPPCallerRaw struct {
	Contract *BaasPPCaller // Generic read-only contract binding to access the raw methods on
}

// BaasPPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasPPTransactorRaw struct {
	Contract *BaasPPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasPP creates a new instance of BaasPP, bound to a specific deployed contract.
func NewBaasPP(address common.Address, backend bind.ContractBackend) (*BaasPP, error) {
	contract, err := bindBaasPP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasPP{BaasPPCaller: BaasPPCaller{contract: contract}, BaasPPTransactor: BaasPPTransactor{contract: contract}, BaasPPFilterer: BaasPPFilterer{contract: contract}}, nil
}

// NewBaasPPCaller creates a new read-only instance of BaasPP, bound to a specific deployed contract.
func NewBaasPPCaller(address common.Address, caller bind.ContractCaller) (*BaasPPCaller, error) {
	contract, err := bindBaasPP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasPPCaller{contract: contract}, nil
}

// NewBaasPPTransactor creates a new write-only instance of BaasPP, bound to a specific deployed contract.
func NewBaasPPTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasPPTransactor, error) {
	contract, err := bindBaasPP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasPPTransactor{contract: contract}, nil
}

// NewBaasPPFilterer creates a new log filterer instance of BaasPP, bound to a specific deployed contract.
func NewBaasPPFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasPPFilterer, error) {
	contract, err := bindBaasPP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasPPFilterer{contract: contract}, nil
}

// bindBaasPP binds a generic wrapper to an already deployed contract.
func bindBaasPP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasPPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasPP *BaasPPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasPP.Contract.BaasPPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasPP *BaasPPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.Contract.BaasPPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasPP *BaasPPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasPP.Contract.BaasPPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasPP *BaasPPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasPP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasPP *BaasPPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasPP *BaasPPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasPP.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasPP *BaasPPCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasPP *BaasPPSession) Balance() (*big.Int, error) {
	return _BaasPP.Contract.Balance(&_BaasPP.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_BaasPP *BaasPPCallerSession) Balance() (*big.Int, error) {
	return _BaasPP.Contract.Balance(&_BaasPP.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0xff2ad8e4.
//
// Solidity: function cap(discountType uint256) constant returns(uint256)
func (_BaasPP *BaasPPCaller) Cap(opts *bind.CallOpts, discountType *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "cap", discountType)
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0xff2ad8e4.
//
// Solidity: function cap(discountType uint256) constant returns(uint256)
func (_BaasPP *BaasPPSession) Cap(discountType *big.Int) (*big.Int, error) {
	return _BaasPP.Contract.Cap(&_BaasPP.CallOpts, discountType)
}

// Cap is a free data retrieval call binding the contract method 0xff2ad8e4.
//
// Solidity: function cap(discountType uint256) constant returns(uint256)
func (_BaasPP *BaasPPCallerSession) Cap(discountType *big.Int) (*big.Int, error) {
	return _BaasPP.Contract.Cap(&_BaasPP.CallOpts, discountType)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_BaasPP *BaasPPCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_BaasPP *BaasPPSession) IsFinalized() (bool, error) {
	return _BaasPP.Contract.IsFinalized(&_BaasPP.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() constant returns(bool)
func (_BaasPP *BaasPPCallerSession) IsFinalized() (bool, error) {
	return _BaasPP.Contract.IsFinalized(&_BaasPP.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasPP *BaasPPCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasPP *BaasPPSession) IsOwner() (bool, error) {
	return _BaasPP.Contract.IsOwner(&_BaasPP.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasPP *BaasPPCallerSession) IsOwner() (bool, error) {
	return _BaasPP.Contract.IsOwner(&_BaasPP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasPP *BaasPPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasPP *BaasPPSession) Owner() (common.Address, error) {
	return _BaasPP.Contract.Owner(&_BaasPP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasPP *BaasPPCallerSession) Owner() (common.Address, error) {
	return _BaasPP.Contract.Owner(&_BaasPP.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasPP *BaasPPCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasPP *BaasPPSession) TokenAddress() (common.Address, error) {
	return _BaasPP.Contract.TokenAddress(&_BaasPP.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_BaasPP *BaasPPCallerSession) TokenAddress() (common.Address, error) {
	return _BaasPP.Contract.TokenAddress(&_BaasPP.CallOpts)
}

// TokensIssued is a free data retrieval call binding the contract method 0x5c5e1d52.
//
// Solidity: function tokensIssued(discountType uint8) constant returns(uint256)
func (_BaasPP *BaasPPCaller) TokensIssued(opts *bind.CallOpts, discountType uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasPP.contract.Call(opts, out, "tokensIssued", discountType)
	return *ret0, err
}

// TokensIssued is a free data retrieval call binding the contract method 0x5c5e1d52.
//
// Solidity: function tokensIssued(discountType uint8) constant returns(uint256)
func (_BaasPP *BaasPPSession) TokensIssued(discountType uint8) (*big.Int, error) {
	return _BaasPP.Contract.TokensIssued(&_BaasPP.CallOpts, discountType)
}

// TokensIssued is a free data retrieval call binding the contract method 0x5c5e1d52.
//
// Solidity: function tokensIssued(discountType uint8) constant returns(uint256)
func (_BaasPP *BaasPPCallerSession) TokensIssued(discountType uint8) (*big.Int, error) {
	return _BaasPP.Contract.TokensIssued(&_BaasPP.CallOpts, discountType)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_BaasPP *BaasPPTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_BaasPP *BaasPPSession) Finalize() (*types.Transaction, error) {
	return _BaasPP.Contract.Finalize(&_BaasPP.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns(bool)
func (_BaasPP *BaasPPTransactorSession) Finalize() (*types.Transaction, error) {
	return _BaasPP.Contract.Finalize(&_BaasPP.TransactOpts)
}

// Issue is a paid mutator transaction binding the contract method 0x26215f25.
//
// Solidity: function issue(account address, amount uint256, discountType uint8) returns(bool)
func (_BaasPP *BaasPPTransactor) Issue(opts *bind.TransactOpts, account common.Address, amount *big.Int, discountType uint8) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "issue", account, amount, discountType)
}

// Issue is a paid mutator transaction binding the contract method 0x26215f25.
//
// Solidity: function issue(account address, amount uint256, discountType uint8) returns(bool)
func (_BaasPP *BaasPPSession) Issue(account common.Address, amount *big.Int, discountType uint8) (*types.Transaction, error) {
	return _BaasPP.Contract.Issue(&_BaasPP.TransactOpts, account, amount, discountType)
}

// Issue is a paid mutator transaction binding the contract method 0x26215f25.
//
// Solidity: function issue(account address, amount uint256, discountType uint8) returns(bool)
func (_BaasPP *BaasPPTransactorSession) Issue(account common.Address, amount *big.Int, discountType uint8) (*types.Transaction, error) {
	return _BaasPP.Contract.Issue(&_BaasPP.TransactOpts, account, amount, discountType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasPP *BaasPPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasPP *BaasPPSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasPP.Contract.RenounceOwnership(&_BaasPP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasPP *BaasPPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasPP.Contract.RenounceOwnership(&_BaasPP.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasPP *BaasPPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasPP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasPP *BaasPPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasPP.Contract.TransferOwnership(&_BaasPP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasPP *BaasPPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasPP.Contract.TransferOwnership(&_BaasPP.TransactOpts, newOwner)
}

// BaasPPFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the BaasPP contract.
type BaasPPFinalizedIterator struct {
	Event *BaasPPFinalized // Event containing the contract specifics and raw log

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
func (it *BaasPPFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasPPFinalized)
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
		it.Event = new(BaasPPFinalized)
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
func (it *BaasPPFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasPPFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasPPFinalized represents a Finalized event raised by the BaasPP contract.
type BaasPPFinalized struct {
	BurnedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x839cf22e1ba87ce2f5b9bbf46cf0175a09eed52febdfaac8852478e68203c763.
//
// Solidity: e Finalized(burnedAmount uint256)
func (_BaasPP *BaasPPFilterer) FilterFinalized(opts *bind.FilterOpts) (*BaasPPFinalizedIterator, error) {

	logs, sub, err := _BaasPP.contract.FilterLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return &BaasPPFinalizedIterator{contract: _BaasPP.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x839cf22e1ba87ce2f5b9bbf46cf0175a09eed52febdfaac8852478e68203c763.
//
// Solidity: e Finalized(burnedAmount uint256)
func (_BaasPP *BaasPPFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *BaasPPFinalized) (event.Subscription, error) {

	logs, sub, err := _BaasPP.contract.WatchLogs(opts, "Finalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasPPFinalized)
				if err := _BaasPP.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// BaasPPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasPP contract.
type BaasPPOwnershipTransferredIterator struct {
	Event *BaasPPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasPPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasPPOwnershipTransferred)
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
		it.Event = new(BaasPPOwnershipTransferred)
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
func (it *BaasPPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasPPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasPPOwnershipTransferred represents a OwnershipTransferred event raised by the BaasPP contract.
type BaasPPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasPP *BaasPPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasPPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasPP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasPPOwnershipTransferredIterator{contract: _BaasPP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasPP *BaasPPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasPPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasPP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasPPOwnershipTransferred)
				if err := _BaasPP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaasPPTokensIssuedIterator is returned from FilterTokensIssued and is used to iterate over the raw logs and unpacked data for TokensIssued events raised by the BaasPP contract.
type BaasPPTokensIssuedIterator struct {
	Event *BaasPPTokensIssued // Event containing the contract specifics and raw log

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
func (it *BaasPPTokensIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasPPTokensIssued)
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
		it.Event = new(BaasPPTokensIssued)
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
func (it *BaasPPTokensIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasPPTokensIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasPPTokensIssued represents a TokensIssued event raised by the BaasPP contract.
type BaasPPTokensIssued struct {
	Account      common.Address
	DiscountType uint8
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensIssued is a free log retrieval operation binding the contract event 0x064e268fa97e769a5abed875cfd116cb13d34b361b6088e9213df29cc7ea85e4.
//
// Solidity: e TokensIssued(account indexed address, discountType indexed uint8, amount uint256)
func (_BaasPP *BaasPPFilterer) FilterTokensIssued(opts *bind.FilterOpts, account []common.Address, discountType []uint8) (*BaasPPTokensIssuedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var discountTypeRule []interface{}
	for _, discountTypeItem := range discountType {
		discountTypeRule = append(discountTypeRule, discountTypeItem)
	}

	logs, sub, err := _BaasPP.contract.FilterLogs(opts, "TokensIssued", accountRule, discountTypeRule)
	if err != nil {
		return nil, err
	}
	return &BaasPPTokensIssuedIterator{contract: _BaasPP.contract, event: "TokensIssued", logs: logs, sub: sub}, nil
}

// WatchTokensIssued is a free log subscription operation binding the contract event 0x064e268fa97e769a5abed875cfd116cb13d34b361b6088e9213df29cc7ea85e4.
//
// Solidity: e TokensIssued(account indexed address, discountType indexed uint8, amount uint256)
func (_BaasPP *BaasPPFilterer) WatchTokensIssued(opts *bind.WatchOpts, sink chan<- *BaasPPTokensIssued, account []common.Address, discountType []uint8) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var discountTypeRule []interface{}
	for _, discountTypeItem := range discountType {
		discountTypeRule = append(discountTypeRule, discountTypeItem)
	}

	logs, sub, err := _BaasPP.contract.WatchLogs(opts, "TokensIssued", accountRule, discountTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasPPTokensIssued)
				if err := _BaasPP.contract.UnpackLog(event, "TokensIssued", log); err != nil {
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
