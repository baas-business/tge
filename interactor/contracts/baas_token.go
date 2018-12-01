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
const BaasTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SetupCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"escrowAddress\",\"type\":\"address\"},{\"name\":\"ppAddress\",\"type\":\"address\"},{\"name\":\"founderAddress\",\"type\":\"address\"},{\"name\":\"incentivesAddress\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"potAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnTokensFromPot\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pots\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenHolderAddress\",\"type\":\"address\"}],\"name\":\"isTokenHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenHolderCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenHolderSnapShot\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BaasTokenBin is the compiled bytecode used for deploying new contracts.
const BaasTokenBin = `0x60806040526000600660006101000a81548160ff0219169083151502179055506000600660016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3612476806101f46000396000f300608060405260043610610128576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461012d578063095ea7b3146101bd57806318160ddd1461022257806323b872dd1461024d578063392e53cd146102d257806339509351146103015780634cca31b61461036657806354c35a3c146103cb5780635c975abb1461048657806365731fe9146104b557806370a0823114610510578063715018a6146105675780637af70c1f1461057e5780638da5cb5b146105ea5780638f32d59b14610641578063a457c2d714610670578063a654cfab146106d5578063a9059cbb146107c5578063bedb86fb1461082a578063d13e584614610859578063dd62ed3e14610884578063f2fde38b146108fb575b600080fd5b34801561013957600080fd5b5061014261093e565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610182578082015181840152602081019050610167565b50505050905090810190601f1680156101af5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c957600080fd5b50610208600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061097b565b604051808215151515815260200191505060405180910390f35b34801561022e57600080fd5b50610237610aa8565b6040518082815260200191505060405180910390f35b34801561025957600080fd5b506102b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ab2565b604051808215151515815260200191505060405180910390f35b3480156102de57600080fd5b506102e7610b1d565b604051808215151515815260200191505060405180910390f35b34801561030d57600080fd5b5061034c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b34565b604051808215151515815260200191505060405180910390f35b34801561037257600080fd5b506103b1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d6b565b604051808215151515815260200191505060405180910390f35b3480156103d757600080fd5b5061046c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f8e565b604051808215151515815260200191505060405180910390f35b34801561049257600080fd5b5061049b6112a2565b604051808215151515815260200191505060405180910390f35b3480156104c157600080fd5b506104f6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112b9565b604051808215151515815260200191505060405180910390f35b34801561051c57600080fd5b50610551600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611383565b6040518082815260200191505060405180910390f35b34801561057357600080fd5b5061057c6113cb565b005b34801561058a57600080fd5b5061059361149f565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156105d65780820151818401526020810190506105bb565b505050509050019250505060405180910390f35b3480156105f657600080fd5b506105ff61152d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561064d57600080fd5b50610656611557565b604051808215151515815260200191505060405180910390f35b34801561067c57600080fd5b506106bb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506115af565b604051808215151515815260200191505060405180910390f35b3480156106e157600080fd5b506106ea6117e6565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200194505050505060405180910390f35b3480156107d157600080fd5b50610810600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611886565b604051808215151515815260200191505060405180910390f35b34801561083657600080fd5b506108576004803603810190808035151590602001909291905050506118ee565b005b34801561086557600080fd5b5061086e61197b565b6040518082815260200191505060405180910390f35b34801561089057600080fd5b506108e5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611988565b6040518082815260200191505060405180910390f35b34801561090757600080fd5b5061093c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a0f565b005b60606040805190810160405280600581526020017f544f4b454e000000000000000000000000000000000000000000000000000000815250905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156109b857600080fd5b81600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600254905090565b600080600660009054906101000a900460ff168015610ade5750600360149054906101000a900460ff16155b1515610ae957600080fd5b610af4858585611a2e565b90508015610b1157610b0585611b55565b50610b0f84611b55565b505b60019150509392505050565b6000600660009054906101000a900460ff16905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610b7157600080fd5b610c0082600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cce90919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b6000600660009054906101000a900460ff161515610d8857600080fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610df45750610dc561152d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610dff57600080fd5b600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161480610ea85750600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80610f005750600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b80610f585750600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1515610f6357600080fd5b610f6c83611383565b8211151515610f7a57600080fd5b610f848383611cef565b6001905092915050565b6000610f98611557565b1515610fa357600080fd5b600660009054906101000a900460ff16151515610fbf57600080fd5b84600660016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110fa600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a31a17e847807b1bc000000611e43565b611131600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a108b2a2c28029094000000611e43565b611168600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a084595161401484a000000611e43565b61119f600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166a084595161401484a000000611e43565b6111ca600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611b55565b506111f6600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611b55565b50611222600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611b55565b5061124e600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611b55565b506001600660006101000a81548160ff0219169083151502179055507f8601fb6b43cf244c5828073ea9e16745bc09e755520a6a80cecae662b31db52860405160405180910390a160019050949350505050565b6000600360149054906101000a900460ff16905090565b60008060058054905014156112d1576000905061137e565b8173ffffffffffffffffffffffffffffffffffffffff166005600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015481548110151561133957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505b919050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6113d3611557565b15156113de57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6060600580548060200260200160405190810160405280929190818152602001828054801561152357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116114d9575b5050505050905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156115ec57600080fd5b61167b82600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f9790919063ffffffff16565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b600080600080600660019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16935093509350935090919293565b600080600660009054906101000a900460ff1680156118b25750600360149054906101000a900460ff16155b15156118bd57600080fd5b6118c78484611fb8565b905080156118e4576118d833611b55565b506118e284611b55565b505b8091505092915050565b6118f6611557565b151561190157600080fd5b801515600360149054906101000a900460ff1615151415151561192357600080fd5b80600360146101000a81548160ff0219169083151502179055507f0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd281604051808215151515815260200191505060405180910390a150565b6000600580549050905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b611a17611557565b1515611a2257600080fd5b611a2b81611fcf565b50565b6000611abf82600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f9790919063ffffffff16565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611b4a8484846120cb565b600190509392505050565b600080611b6183611383565b1415611b8557611b70826112b9565b15611b8057611b7e82612297565b505b611cc5565b611b8e826112b9565b1515611cc45781600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160058390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055505b5b60019050919050565b6000808284019050838110151515611ce557600080fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611d2b57600080fd5b611d4081600254611f9790919063ffffffff16565b600281905550611d97816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f9790919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611e7f57600080fd5b611e9481600254611cce90919063ffffffff16565b600281905550611eeb816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cce90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600080838311151515611fa957600080fd5b82840390508091505092915050565b6000611fc53384846120cb565b6001905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561200b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561210757600080fd5b612158816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611f9790919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506121eb816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cce90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008060006122a5846112b9565b15156122b057600080fd5b600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549150600560016005805490500381548110151561230c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060058381548110151561234957fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060058054809190600190036123ed91906123f9565b50600192505050919050565b8154818355818111156124205781836000526020600020918201910161241f9190612425565b5b505050565b61244791905b8082111561244357600081600090555060010161242b565b5090565b905600a165627a7a72305820d0b248f210ce72a418ec439ec4290c2879e1e22094f789f58d3922e10ef321b40029`

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

// IsTokenHolder is a free data retrieval call binding the contract method 0x65731fe9.
//
// Solidity: function isTokenHolder(tokenHolderAddress address) constant returns(bool)
func (_BaasToken *BaasTokenCaller) IsTokenHolder(opts *bind.CallOpts, tokenHolderAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BaasToken.contract.Call(opts, out, "isTokenHolder", tokenHolderAddress)
	return *ret0, err
}

// IsTokenHolder is a free data retrieval call binding the contract method 0x65731fe9.
//
// Solidity: function isTokenHolder(tokenHolderAddress address) constant returns(bool)
func (_BaasToken *BaasTokenSession) IsTokenHolder(tokenHolderAddress common.Address) (bool, error) {
	return _BaasToken.Contract.IsTokenHolder(&_BaasToken.CallOpts, tokenHolderAddress)
}

// IsTokenHolder is a free data retrieval call binding the contract method 0x65731fe9.
//
// Solidity: function isTokenHolder(tokenHolderAddress address) constant returns(bool)
func (_BaasToken *BaasTokenCallerSession) IsTokenHolder(tokenHolderAddress common.Address) (bool, error) {
	return _BaasToken.Contract.IsTokenHolder(&_BaasToken.CallOpts, tokenHolderAddress)
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
