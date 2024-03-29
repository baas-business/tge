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

// BaasTokenABI is the input ABI used to generate the binding from.
const BaasTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SetupCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"escrowAddress\",\"type\":\"address\"},{\"name\":\"ppAddress\",\"type\":\"address\"},{\"name\":\"founderAddress\",\"type\":\"address\"},{\"name\":\"incentivesAddress\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"potAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnTokensFromPot\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPotAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pots\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isTokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenHolderCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenHolderSnapShot\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"circulatingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"potSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BaasTokenBin is the compiled bytecode used for deploying new contracts.
const BaasTokenBin = `0x60806040526000600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3612d61806101d96000396000f3fe60806040526004361061015f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde0314610164578063095ea7b3146101f457806318160ddd1461026757806323b872dd14610292578063313ce56714610325578063392e53cd14610350578063395093511461037f5780634cca31b6146103f257806354c35a3c146104655780635c975abb1461052e57806365731fe91461055d5780636f8220ec146105c657806370a08231146105f1578063715018a6146106565780637af70c1f1461066d5780638da5cb5b146106d95780638f32d59b146107305780639358928b1461075f57806395d89b411461078a578063a457c2d71461081a578063a654cfab1461088d578063a9059cbb1461097d578063bedb86fb146109f0578063d13e584614610a2d578063dd62ed3e14610a58578063e09ca55714610add578063f2fde38b14610b46575b600080fd5b34801561017057600080fd5b50610179610b97565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101b957808201518184015260208101905061019e565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020057600080fd5b5061024d6004803603604081101561021757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bd0565b604051808215151515815260200191505060405180910390f35b34801561027357600080fd5b5061027c610cfd565b6040518082815260200191505060405180910390f35b34801561029e57600080fd5b5061030b600480360360608110156102b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d07565b604051808215151515815260200191505060405180910390f35b34801561033157600080fd5b5061033a610ec2565b6040518082815260200191505060405180910390f35b34801561035c57600080fd5b50610365610ec7565b604051808215151515815260200191505060405180910390f35b34801561038b57600080fd5b506103d8600480360360408110156103a257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ede565b604051808215151515815260200191505060405180910390f35b3480156103fe57600080fd5b5061044b6004803603604081101561041557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611115565b604051808215151515815260200191505060405180910390f35b34801561047157600080fd5b506105146004803603608081101561048857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113cb565b604051808215151515815260200191505060405180910390f35b34801561053a57600080fd5b506105436116df565b604051808215151515815260200191505060405180910390f35b34801561056957600080fd5b506105ac6004803603602081101561058057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116f6565b604051808215151515815260200191505060405180910390f35b3480156105d257600080fd5b506105db6117c0565b6040518082815260200191505060405180910390f35b3480156105fd57600080fd5b506106406004803603602081101561061457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506118de565b6040518082815260200191505060405180910390f35b34801561066257600080fd5b5061066b611926565b005b34801561067957600080fd5b506106826119fa565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156106c55780820151818401526020810190506106aa565b505050509050019250505060405180910390f35b3480156106e557600080fd5b506106ee611a88565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561073c57600080fd5b50610745611ab2565b604051808215151515815260200191505060405180910390f35b34801561076b57600080fd5b50610774611b0a565b6040518082815260200191505060405180910390f35b34801561079657600080fd5b5061079f611b32565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107df5780820151818401526020810190506107c4565b50505050905090810190601f16801561080c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561082657600080fd5b506108736004803603604081101561083d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611b6b565b604051808215151515815260200191505060405180910390f35b34801561089957600080fd5b506108a2611da2565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390f35b34801561098957600080fd5b506109d6600480360360408110156109a057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611e42565b604051808215151515815260200191505060405180910390f35b3480156109fc57600080fd5b50610a2b60048036036020811015610a1357600080fd5b81019080803515159060200190929190505050611ffb565b005b348015610a3957600080fd5b50610a42612117565b6040518082815260200191505060405180910390f35b348015610a6457600080fd5b50610ac760048036036040811015610a7b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612124565b6040518082815260200191505060405180910390f35b348015610ae957600080fd5b50610b2c60048036036020811015610b0057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506121ab565b604051808215151515815260200191505060405180910390f35b348015610b5257600080fd5b50610b9560048036036020811015610b6957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061230d565b005b6040805190810160405280600a81526020017f4261615320546f6b656e0000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610c0d57600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600254905090565b6000600760159054906101000a900460ff168015610d325750600760149054906101000a900460ff16155b1515610da6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f656974686572206e6f7420696e697469616c697a6564206f722070617573656481525060200191505060405180910390fd5b610daf836121ab565b151515610e24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f63616e6e6f742073656e6420746f6b656e7320746f20706f747300000000000081525060200191505060405180910390fd5b610e2f84848461232c565b1515610ea3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6661696c656420746f207472616e7366657220746f6b656e730000000000000081525060200191505060405180910390fd5b610eac84612453565b50610eb683612453565b50600190509392505050565b601281565b6000600760159054906101000a900460ff16905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610f1b57600080fd5b610faa82600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546125cc90919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b6000600760159054906101000a900460ff16151561119b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6e6565647320746f20626520696e697469616c697a656400000000000000000081525060200191505060405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611264576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001807f6f6e6c7920706f7473207468656d73656c7665732063616e206275726e20746f81526020017f6b656e730000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b61126d836121ab565b1515611307576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001807f6f6e6c7920746f6b656e732066726f6d20706f742063616e206265206275726e81526020017f656400000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b611310836118de565b82111515156113ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f616d6f756e74206d75737420626520736d616c6c6572207468656e20746f6b6581526020017f6e732068656c642062792074686520706f74000000000000000000000000000081525060400191505060405180910390fd5b6113b783836125ed565b6113c083612453565b506001905092915050565b60006113d5611ab2565b15156113e057600080fd5b600760159054906101000a900460ff161515156113fc57600080fd5b84600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611537600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a31a17e847807b1bc000000612741565b61156e600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a108b2a2c28029094000000612741565b6115a5600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a084595161401484a000000612741565b6115dc600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a084595161401484a000000612741565b611607600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612453565b50611633600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612453565b5061165f600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612453565b5061168b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612453565b506001600760156101000a81548160ff0219169083151502179055507f8601fb6b43cf244c5828073ea9e16745bc09e755520a6a80cecae662b31db52860405160405180910390a160019050949350505050565b6000600760149054906101000a900460ff16905090565b600080600980549050141561170e57600090506117bb565b8173ffffffffffffffffffffffffffffffffffffffff166009600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481548110151561177657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505b919050565b6000600760159054906101000a900460ff1615156117e157600090506118db565b600061180e600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166118de565b9050600061183d600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166118de565b9050600061186c600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166118de565b9050600061189b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166118de565b90506118d4826118c6856118b885896125cc90919063ffffffff16565b6125cc90919063ffffffff16565b6125cc90919063ffffffff16565b9450505050505b90565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61192e611ab2565b151561193957600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60606009805480602002602001604051908101604052809291908181526020018280548015611a7e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611a34575b5050505050905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000611b2d611b176117c0565b611b1f610cfd565b61289590919063ffffffff16565b905090565b6040805190810160405280600481526020017f426161530000000000000000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515611ba857600080fd5b611c3782600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461289590919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b600080600080600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16935093509350935090919293565b6000600760159054906101000a900460ff168015611e6d5750600760149054906101000a900460ff16155b1515611ee1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f656974686572206e6f7420696e697469616c697a6564206f722070617573656481525060200191505060405180910390fd5b611eea836121ab565b151515611f5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f63616e6e6f742073656e6420746f6b656e7320746f20706f747300000000000081525060200191505060405180910390fd5b611f6983836128b7565b1515611fdd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f6661696c656420746f207472616e7366657220746f6b656e730000000000000081525060200191505060405180910390fd5b611fe633612453565b50611ff083612453565b506001905092915050565b612003611ab2565b151561200e57600080fd5b801515600760149054906101000a900460ff161515141515156120bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f63616e6e6f742070617573652f756e706175736520696620616c72656164792081526020017f7061757365642f756e706175736564000000000000000000000000000000000081525060400191505060405180910390fd5b80600760146101000a81548160ff0219169083151502179055507f0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd281604051808215151515815260200191505060405180910390a150565b6000600980549050905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806122565750600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b806122ae5750600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b806123065750600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b9050919050565b612315611ab2565b151561232057600080fd5b612329816128ce565b50565b60006123bd82600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461289590919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506124488484846129ca565b600190509392505050565b60008061245f836118de565b14156124835761246e826116f6565b1561247e5761247c82612b96565b505b6125c3565b61248c826116f6565b15156125c25781600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160098390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055505b5b60019050919050565b60008082840190508381101515156125e357600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561262957600080fd5b61263e8160025461289590919063ffffffff16565b600281905550612695816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461289590919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561277d57600080fd5b612792816002546125cc90919063ffffffff16565b6002819055506127e9816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546125cc90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b60008282111515156128a657600080fd5b600082840390508091505092915050565b60006128c43384846129ca565b6001905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561290a57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515612a0657600080fd5b612a57816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461289590919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612aea816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546125cc90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905060006009600160098054905003815481101515612bf757fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600983815481101515612c3457fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506009805480919060019003612cd89190612ce4565b50600192505050919050565b815481835581811115612d0b57818360005260206000209182019101612d0a9190612d10565b5b505050565b612d3291905b80821115612d2e576000816000905550600101612d16565b5090565b9056fea165627a7a7230582096c8f2b5bf1e09304991467d93170a4ed9108ec752942866f178aaf0f7388b520029`

// DeployBaasToken deploys a new Ethereum contract, binding an instance of BaasToken to it.
func DeployBaasToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaasToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BaasTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaasToken{BaasTokenCaller: BaasTokenCaller{contract: contract}, BaasTokenTransactor: BaasTokenTransactor{contract: contract}, BaasTokenFilterer: BaasTokenFilterer{contract: contract}}, nil
}

// BaasToken is an auto generated Go binding around an Ethereum contract.
type BaasToken struct {
	BaasTokenCaller     // Read-only binding to the contract
	BaasTokenTransactor // Write-only binding to the contract
	BaasTokenFilterer   // Log filterer for contract events
}

// BaasTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaasTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaasTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaasTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaasTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaasTokenSession struct {
	Contract     *BaasToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaasTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaasTokenCallerSession struct {
	Contract *BaasTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BaasTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaasTokenTransactorSession struct {
	Contract     *BaasTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BaasTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaasTokenRaw struct {
	Contract *BaasToken // Generic contract binding to access the raw methods on
}

// BaasTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaasTokenCallerRaw struct {
	Contract *BaasTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BaasTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaasTokenTransactorRaw struct {
	Contract *BaasTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaasToken creates a new instance of BaasToken, bound to a specific deployed contract.
func NewBaasToken(address common.Address, backend bind.ContractBackend) (*BaasToken, error) {
	contract, err := bindBaasToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaasToken{BaasTokenCaller: BaasTokenCaller{contract: contract}, BaasTokenTransactor: BaasTokenTransactor{contract: contract}, BaasTokenFilterer: BaasTokenFilterer{contract: contract}}, nil
}

// NewBaasTokenCaller creates a new read-only instance of BaasToken, bound to a specific deployed contract.
func NewBaasTokenCaller(address common.Address, caller bind.ContractCaller) (*BaasTokenCaller, error) {
	contract, err := bindBaasToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaasTokenCaller{contract: contract}, nil
}

// NewBaasTokenTransactor creates a new write-only instance of BaasToken, bound to a specific deployed contract.
func NewBaasTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BaasTokenTransactor, error) {
	contract, err := bindBaasToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaasTokenTransactor{contract: contract}, nil
}

// NewBaasTokenFilterer creates a new log filterer instance of BaasToken, bound to a specific deployed contract.
func NewBaasTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BaasTokenFilterer, error) {
	contract, err := bindBaasToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaasTokenFilterer{contract: contract}, nil
}

// bindBaasToken binds a generic wrapper to an already deployed contract.
func bindBaasToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaasTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasToken *BaasTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasToken.Contract.BaasTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasToken *BaasTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasToken.Contract.BaasTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasToken *BaasTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasToken.Contract.BaasTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaasToken *BaasTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BaasToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaasToken *BaasTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaasToken *BaasTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaasToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_BaasToken *BaasTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_BaasToken *BaasTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaasToken.Contract.Allowance(&_BaasToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BaasToken.Contract.Allowance(&_BaasToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_BaasToken *BaasTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_BaasToken *BaasTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BaasToken.Contract.BalanceOf(&_BaasToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BaasToken.Contract.BalanceOf(&_BaasToken.CallOpts, owner)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasToken *BaasTokenCaller) CirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "circulatingSupply")
	return *ret0, err
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasToken *BaasTokenSession) CirculatingSupply() (*big.Int, error) {
	return _BaasToken.Contract.CirculatingSupply(&_BaasToken.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0x9358928b.
//
// Solidity: function circulatingSupply() constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) CirculatingSupply() (*big.Int, error) {
	return _BaasToken.Contract.CirculatingSupply(&_BaasToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BaasToken *BaasTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BaasToken *BaasTokenSession) Decimals() (*big.Int, error) {
	return _BaasToken.Contract.Decimals(&_BaasToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) Decimals() (*big.Int, error) {
	return _BaasToken.Contract.Decimals(&_BaasToken.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasToken *BaasTokenCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "isInitialized")
	return *ret0, err
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasToken *BaasTokenSession) IsInitialized() (bool, error) {
	return _BaasToken.Contract.IsInitialized(&_BaasToken.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() constant returns(bool)
func (_BaasToken *BaasTokenCallerSession) IsInitialized() (bool, error) {
	return _BaasToken.Contract.IsInitialized(&_BaasToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasToken *BaasTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasToken *BaasTokenSession) IsOwner() (bool, error) {
	return _BaasToken.Contract.IsOwner(&_BaasToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BaasToken *BaasTokenCallerSession) IsOwner() (bool, error) {
	return _BaasToken.Contract.IsOwner(&_BaasToken.CallOpts)
}

// IsPotAddress is a free data retrieval call binding the contract method 0xe09ca557.
//
// Solidity: function isPotAddress(account address) constant returns(bool)
func (_BaasToken *BaasTokenCaller) IsPotAddress(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "isPotAddress", account)
	return *ret0, err
}

// IsPotAddress is a free data retrieval call binding the contract method 0xe09ca557.
//
// Solidity: function isPotAddress(account address) constant returns(bool)
func (_BaasToken *BaasTokenSession) IsPotAddress(account common.Address) (bool, error) {
	return _BaasToken.Contract.IsPotAddress(&_BaasToken.CallOpts, account)
}

// IsPotAddress is a free data retrieval call binding the contract method 0xe09ca557.
//
// Solidity: function isPotAddress(account address) constant returns(bool)
func (_BaasToken *BaasTokenCallerSession) IsPotAddress(account common.Address) (bool, error) {
	return _BaasToken.Contract.IsPotAddress(&_BaasToken.CallOpts, account)
}

// IsTokenHolder is a free data retrieval call binding the contract method 0x65731fe9.
//
// Solidity: function isTokenHolder(account address) constant returns(bool)
func (_BaasToken *BaasTokenCaller) IsTokenHolder(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "isTokenHolder", account)
	return *ret0, err
}

// IsTokenHolder is a free data retrieval call binding the contract method 0x65731fe9.
//
// Solidity: function isTokenHolder(account address) constant returns(bool)
func (_BaasToken *BaasTokenSession) IsTokenHolder(account common.Address) (bool, error) {
	return _BaasToken.Contract.IsTokenHolder(&_BaasToken.CallOpts, account)
}

// IsTokenHolder is a free data retrieval call binding the contract method 0x65731fe9.
//
// Solidity: function isTokenHolder(account address) constant returns(bool)
func (_BaasToken *BaasTokenCallerSession) IsTokenHolder(account common.Address) (bool, error) {
	return _BaasToken.Contract.IsTokenHolder(&_BaasToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasToken *BaasTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasToken *BaasTokenSession) Name() (string, error) {
	return _BaasToken.Contract.Name(&_BaasToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BaasToken *BaasTokenCallerSession) Name() (string, error) {
	return _BaasToken.Contract.Name(&_BaasToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasToken *BaasTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasToken *BaasTokenSession) Owner() (common.Address, error) {
	return _BaasToken.Contract.Owner(&_BaasToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BaasToken *BaasTokenCallerSession) Owner() (common.Address, error) {
	return _BaasToken.Contract.Owner(&_BaasToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BaasToken *BaasTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BaasToken *BaasTokenSession) Paused() (bool, error) {
	return _BaasToken.Contract.Paused(&_BaasToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_BaasToken *BaasTokenCallerSession) Paused() (bool, error) {
	return _BaasToken.Contract.Paused(&_BaasToken.CallOpts)
}

// PotSupply is a free data retrieval call binding the contract method 0x6f8220ec.
//
// Solidity: function potSupply() constant returns(uint256)
func (_BaasToken *BaasTokenCaller) PotSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "potSupply")
	return *ret0, err
}

// PotSupply is a free data retrieval call binding the contract method 0x6f8220ec.
//
// Solidity: function potSupply() constant returns(uint256)
func (_BaasToken *BaasTokenSession) PotSupply() (*big.Int, error) {
	return _BaasToken.Contract.PotSupply(&_BaasToken.CallOpts)
}

// PotSupply is a free data retrieval call binding the contract method 0x6f8220ec.
//
// Solidity: function potSupply() constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) PotSupply() (*big.Int, error) {
	return _BaasToken.Contract.PotSupply(&_BaasToken.CallOpts)
}

// Pots is a free data retrieval call binding the contract method 0xa654cfab.
//
// Solidity: function pots() constant returns(address, address, address, address)
func (_BaasToken *BaasTokenCaller) Pots(opts *bind.CallOpts) (common.Address, common.Address, common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(common.Address)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _BaasToken.contract.Call(opts, out, "pots")
	return *ret0, *ret1, *ret2, *ret3, err
}

// Pots is a free data retrieval call binding the contract method 0xa654cfab.
//
// Solidity: function pots() constant returns(address, address, address, address)
func (_BaasToken *BaasTokenSession) Pots() (common.Address, common.Address, common.Address, common.Address, error) {
	return _BaasToken.Contract.Pots(&_BaasToken.CallOpts)
}

// Pots is a free data retrieval call binding the contract method 0xa654cfab.
//
// Solidity: function pots() constant returns(address, address, address, address)
func (_BaasToken *BaasTokenCallerSession) Pots() (common.Address, common.Address, common.Address, common.Address, error) {
	return _BaasToken.Contract.Pots(&_BaasToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BaasToken *BaasTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BaasToken *BaasTokenSession) Symbol() (string, error) {
	return _BaasToken.Contract.Symbol(&_BaasToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BaasToken *BaasTokenCallerSession) Symbol() (string, error) {
	return _BaasToken.Contract.Symbol(&_BaasToken.CallOpts)
}

// TokenHolderCount is a free data retrieval call binding the contract method 0xd13e5846.
//
// Solidity: function tokenHolderCount() constant returns(uint256)
func (_BaasToken *BaasTokenCaller) TokenHolderCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "tokenHolderCount")
	return *ret0, err
}

// TokenHolderCount is a free data retrieval call binding the contract method 0xd13e5846.
//
// Solidity: function tokenHolderCount() constant returns(uint256)
func (_BaasToken *BaasTokenSession) TokenHolderCount() (*big.Int, error) {
	return _BaasToken.Contract.TokenHolderCount(&_BaasToken.CallOpts)
}

// TokenHolderCount is a free data retrieval call binding the contract method 0xd13e5846.
//
// Solidity: function tokenHolderCount() constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) TokenHolderCount() (*big.Int, error) {
	return _BaasToken.Contract.TokenHolderCount(&_BaasToken.CallOpts)
}

// TokenHolderSnapShot is a free data retrieval call binding the contract method 0x7af70c1f.
//
// Solidity: function tokenHolderSnapShot() constant returns(address[])
func (_BaasToken *BaasTokenCaller) TokenHolderSnapShot(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "tokenHolderSnapShot")
	return *ret0, err
}

// TokenHolderSnapShot is a free data retrieval call binding the contract method 0x7af70c1f.
//
// Solidity: function tokenHolderSnapShot() constant returns(address[])
func (_BaasToken *BaasTokenSession) TokenHolderSnapShot() ([]common.Address, error) {
	return _BaasToken.Contract.TokenHolderSnapShot(&_BaasToken.CallOpts)
}

// TokenHolderSnapShot is a free data retrieval call binding the contract method 0x7af70c1f.
//
// Solidity: function tokenHolderSnapShot() constant returns(address[])
func (_BaasToken *BaasTokenCallerSession) TokenHolderSnapShot() ([]common.Address, error) {
	return _BaasToken.Contract.TokenHolderSnapShot(&_BaasToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BaasToken *BaasTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BaasToken *BaasTokenSession) TotalSupply() (*big.Int, error) {
	return _BaasToken.Contract.TotalSupply(&_BaasToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BaasToken *BaasTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BaasToken.Contract.TotalSupply(&_BaasToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_BaasToken *BaasTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_BaasToken *BaasTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.Approve(&_BaasToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.Approve(&_BaasToken.TransactOpts, spender, value)
}

// BurnTokensFromPot is a paid mutator transaction binding the contract method 0x4cca31b6.
//
// Solidity: function burnTokensFromPot(potAddress address, amount uint256) returns(bool)
func (_BaasToken *BaasTokenTransactor) BurnTokensFromPot(opts *bind.TransactOpts, potAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "burnTokensFromPot", potAddress, amount)
}

// BurnTokensFromPot is a paid mutator transaction binding the contract method 0x4cca31b6.
//
// Solidity: function burnTokensFromPot(potAddress address, amount uint256) returns(bool)
func (_BaasToken *BaasTokenSession) BurnTokensFromPot(potAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.BurnTokensFromPot(&_BaasToken.TransactOpts, potAddress, amount)
}

// BurnTokensFromPot is a paid mutator transaction binding the contract method 0x4cca31b6.
//
// Solidity: function burnTokensFromPot(potAddress address, amount uint256) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) BurnTokensFromPot(potAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.BurnTokensFromPot(&_BaasToken.TransactOpts, potAddress, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_BaasToken *BaasTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_BaasToken *BaasTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.DecreaseAllowance(&_BaasToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.DecreaseAllowance(&_BaasToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_BaasToken *BaasTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_BaasToken *BaasTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.IncreaseAllowance(&_BaasToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.IncreaseAllowance(&_BaasToken.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasToken *BaasTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasToken *BaasTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasToken.Contract.RenounceOwnership(&_BaasToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BaasToken *BaasTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BaasToken.Contract.RenounceOwnership(&_BaasToken.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(pause bool) returns()
func (_BaasToken *BaasTokenTransactor) SetPause(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "setPause", pause)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(pause bool) returns()
func (_BaasToken *BaasTokenSession) SetPause(pause bool) (*types.Transaction, error) {
	return _BaasToken.Contract.SetPause(&_BaasToken.TransactOpts, pause)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(pause bool) returns()
func (_BaasToken *BaasTokenTransactorSession) SetPause(pause bool) (*types.Transaction, error) {
	return _BaasToken.Contract.SetPause(&_BaasToken.TransactOpts, pause)
}

// Setup is a paid mutator transaction binding the contract method 0x54c35a3c.
//
// Solidity: function setup(escrowAddress address, ppAddress address, founderAddress address, incentivesAddress address) returns(bool)
func (_BaasToken *BaasTokenTransactor) Setup(opts *bind.TransactOpts, escrowAddress common.Address, ppAddress common.Address, founderAddress common.Address, incentivesAddress common.Address) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "setup", escrowAddress, ppAddress, founderAddress, incentivesAddress)
}

// Setup is a paid mutator transaction binding the contract method 0x54c35a3c.
//
// Solidity: function setup(escrowAddress address, ppAddress address, founderAddress address, incentivesAddress address) returns(bool)
func (_BaasToken *BaasTokenSession) Setup(escrowAddress common.Address, ppAddress common.Address, founderAddress common.Address, incentivesAddress common.Address) (*types.Transaction, error) {
	return _BaasToken.Contract.Setup(&_BaasToken.TransactOpts, escrowAddress, ppAddress, founderAddress, incentivesAddress)
}

// Setup is a paid mutator transaction binding the contract method 0x54c35a3c.
//
// Solidity: function setup(escrowAddress address, ppAddress address, founderAddress address, incentivesAddress address) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) Setup(escrowAddress common.Address, ppAddress common.Address, founderAddress common.Address, incentivesAddress common.Address) (*types.Transaction, error) {
	return _BaasToken.Contract.Setup(&_BaasToken.TransactOpts, escrowAddress, ppAddress, founderAddress, incentivesAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_BaasToken *BaasTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_BaasToken *BaasTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.Transfer(&_BaasToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.Transfer(&_BaasToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_BaasToken *BaasTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_BaasToken *BaasTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.TransferFrom(&_BaasToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_BaasToken *BaasTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BaasToken.Contract.TransferFrom(&_BaasToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasToken *BaasTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BaasToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasToken *BaasTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasToken.Contract.TransferOwnership(&_BaasToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BaasToken *BaasTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BaasToken.Contract.TransferOwnership(&_BaasToken.TransactOpts, newOwner)
}

// BaasTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BaasToken contract.
type BaasTokenApprovalIterator struct {
	Event *BaasTokenApproval // Event containing the contract specifics and raw log

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
func (it *BaasTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasTokenApproval)
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
		it.Event = new(BaasTokenApproval)
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
func (it *BaasTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasTokenApproval represents a Approval event raised by the BaasToken contract.
type BaasTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BaasToken *BaasTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BaasTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaasToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BaasTokenApprovalIterator{contract: _BaasToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BaasToken *BaasTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BaasTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BaasToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasTokenApproval)
				if err := _BaasToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// BaasTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BaasToken contract.
type BaasTokenOwnershipTransferredIterator struct {
	Event *BaasTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BaasTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasTokenOwnershipTransferred)
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
		it.Event = new(BaasTokenOwnershipTransferred)
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
func (it *BaasTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BaasToken contract.
type BaasTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasToken *BaasTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BaasTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BaasTokenOwnershipTransferredIterator{contract: _BaasToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BaasToken *BaasTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BaasTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BaasToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasTokenOwnershipTransferred)
				if err := _BaasToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BaasTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BaasToken contract.
type BaasTokenPausedIterator struct {
	Event *BaasTokenPaused // Event containing the contract specifics and raw log

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
func (it *BaasTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasTokenPaused)
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
		it.Event = new(BaasTokenPaused)
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
func (it *BaasTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasTokenPaused represents a Paused event raised by the BaasToken contract.
type BaasTokenPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: e Paused(paused bool)
func (_BaasToken *BaasTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*BaasTokenPausedIterator, error) {

	logs, sub, err := _BaasToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BaasTokenPausedIterator{contract: _BaasToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: e Paused(paused bool)
func (_BaasToken *BaasTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BaasTokenPaused) (event.Subscription, error) {

	logs, sub, err := _BaasToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasTokenPaused)
				if err := _BaasToken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// BaasTokenSetupCompletedIterator is returned from FilterSetupCompleted and is used to iterate over the raw logs and unpacked data for SetupCompleted events raised by the BaasToken contract.
type BaasTokenSetupCompletedIterator struct {
	Event *BaasTokenSetupCompleted // Event containing the contract specifics and raw log

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
func (it *BaasTokenSetupCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasTokenSetupCompleted)
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
		it.Event = new(BaasTokenSetupCompleted)
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
func (it *BaasTokenSetupCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasTokenSetupCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasTokenSetupCompleted represents a SetupCompleted event raised by the BaasToken contract.
type BaasTokenSetupCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetupCompleted is a free log retrieval operation binding the contract event 0x8601fb6b43cf244c5828073ea9e16745bc09e755520a6a80cecae662b31db528.
//
// Solidity: e SetupCompleted()
func (_BaasToken *BaasTokenFilterer) FilterSetupCompleted(opts *bind.FilterOpts) (*BaasTokenSetupCompletedIterator, error) {

	logs, sub, err := _BaasToken.contract.FilterLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return &BaasTokenSetupCompletedIterator{contract: _BaasToken.contract, event: "SetupCompleted", logs: logs, sub: sub}, nil
}

// WatchSetupCompleted is a free log subscription operation binding the contract event 0x8601fb6b43cf244c5828073ea9e16745bc09e755520a6a80cecae662b31db528.
//
// Solidity: e SetupCompleted()
func (_BaasToken *BaasTokenFilterer) WatchSetupCompleted(opts *bind.WatchOpts, sink chan<- *BaasTokenSetupCompleted) (event.Subscription, error) {

	logs, sub, err := _BaasToken.contract.WatchLogs(opts, "SetupCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasTokenSetupCompleted)
				if err := _BaasToken.contract.UnpackLog(event, "SetupCompleted", log); err != nil {
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

// BaasTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BaasToken contract.
type BaasTokenTransferIterator struct {
	Event *BaasTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BaasTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasTokenTransfer)
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
		it.Event = new(BaasTokenTransfer)
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
func (it *BaasTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasTokenTransfer represents a Transfer event raised by the BaasToken contract.
type BaasTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BaasToken *BaasTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaasTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaasToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaasTokenTransferIterator{contract: _BaasToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BaasToken *BaasTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BaasTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaasToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasTokenTransfer)
				if err := _BaasToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
