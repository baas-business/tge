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
const BaasIncentivesABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncentiveProvided1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"vestingStages\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"vestingBlocks\",\"type\":\"uint256\"}],\"name\":\"IncentiveProvided2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stage\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"remainingToken\",\"type\":\"uint256\"}],\"name\":\"Forfeited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"SetupCompleted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"setup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amountPerStage\",\"type\":\"uint256\"},{\"name\":\"totalVestingStages\",\"type\":\"uint256\"},{\"name\":\"blockTimePerStage\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"forfeited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"stage\",\"type\":\"uint256\"}],\"name\":\"claimState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesLeft\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentivesProvided\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incentives\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getIncentive\",\"outputs\":[{\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"uint256\"},{\"name\":\"atBlock\",\"type\":\"uint256\"},{\"name\":\"currentStage\",\"type\":\"uint256\"},{\"name\":\"totalStages\",\"type\":\"uint256\"},{\"name\":\"totalBlocks\",\"type\":\"uint256\"},{\"name\":\"stagesClaimed\",\"type\":\"bool[]\"},{\"name\":\"stagesBlockTime\",\"type\":\"uint256[]\"},{\"name\":\"isValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BaasIncentivesBin is the compiled bytecode used for deploying new contracts.
const BaasIncentivesBin = `0x60806040526000600560006101000a81548160ff02191690831515021790555034801561002b57600080fd5b50604051602080611f2683398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611dcb8061015b6000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde03146100f65780630d5df7ba14610186578063392e53cd146101f25780634e71d92d146102215780635742a36e1461024c578063715018a61461027757806375be7f7f1461028e5780638da5cb5b146102f55780638f32d59b1461034c5780638f70c2d81461037b5780639d76ea58146104905780639de59fad146104e7578063ab77e2d314610542578063b69ef8a81461056d578063ba0bba4014610598578063c5a3487c146105af578063f2fde38b14610628575b600080fd5b34801561010257600080fd5b5061010b61066b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014b578082015181840152602081019050610130565b50505050905090810190601f1680156101785780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019257600080fd5b5061019b6106a8565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101de5780820151818401526020810190506101c3565b505050509050019250505060405180910390f35b3480156101fe57600080fd5b50610207610736565b604051808215151515815260200191505060405180910390f35b34801561022d57600080fd5b5061023661074d565b6040518082815260200191505060405180910390f35b34801561025857600080fd5b5061026161075d565b6040518082815260200191505060405180910390f35b34801561028357600080fd5b5061028c610767565b005b34801561029a57600080fd5b506102d9600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610839565b604051808260ff1660ff16815260200191505060405180910390f35b34801561030157600080fd5b5061030a610aaf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035857600080fd5b50610361610ad8565b604051808215151515815260200191505060405180910390f35b34801561038757600080fd5b506103bc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b2f565b604051808a8152602001898152602001888152602001878152602001868152602001858152602001806020018060200184151515158152602001838103835286818151815260200191508051906020019060200280838360005b83811015610431578082015181840152602081019050610416565b50505050905001838103825285818151815260200191508051906020019060200280838360005b83811015610473578082015181840152602081019050610458565b505050509050019b50505050505050505050505060405180910390f35b34801561049c57600080fd5b506104a5610d74565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f357600080fd5b50610528600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d9e565b604051808215151515815260200191505060405180910390f35b34801561054e57600080fd5b50610557610f03565b6040518082815260200191505060405180910390f35b34801561057957600080fd5b50610582610f0d565b6040518082815260200191505060405180910390f35b3480156105a457600080fd5b506105ad61100c565b005b3480156105bb57600080fd5b5061060e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050611094565b604051808215151515815260200191505060405180910390f35b34801561063457600080fd5b50610669600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115be565b005b60606040805190810160405280600a81526020017f494e43454e544956455300000000000000000000000000000000000000000000815250905090565b6060600280548060200260200160405190810160405280929190818152602001828054801561072c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116106e2575b5050505050905090565b6000600560009054906101000a900460ff16905090565b6000610758336115dd565b905090565b6000600454905090565b61076f610ad8565b151561077a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000610843611be4565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061016060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820180548060200260200160405190810160405280929190818152602001828054801561099557602002820191906000526020600020906000905b82829054906101000a900460ff1615158152602001906001019060208260000104928301926001038202915080841161095f5790505b50505050508152602001600882018054806020026020016040519081016040528092919081815260200182805480156109ed57602002820191906000526020600020905b8154815260200190600101908083116109d9575b5050505050815260200160098201548152602001600a820160009054906101000a900460ff16151515158152505090508061014001511515610a325760049150610aa8565b8260018260e001515103111515610a4c5760039150610aa8565b8060e0015183815181101515610a5e57fe5b9060200190602002015115610a765760029150610aa8565b4381610100015184815181101515610a8a57fe5b906020019060200201511015610aa35760019150610aa8565b600091505b5092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000806000806000806060806000610b45611be4565b600160008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061016060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201805480602002602001604051908101604052809291908181526020018280548015610c9757602002820191906000526020600020906000905b82829054906101000a900460ff16151581526020019060010190602082600001049283019260010382029150808411610c615790505b5050505050815260200160088201805480602002602001604051908101604052809291908181526020018280548015610cef57602002820191906000526020600020905b815481526020019060010190808311610cdb575b5050505050815260200160098201548152602001600a820160009054906101000a900460ff161515151581525050905080602001518160400151826060015183608001518460a001518560c001518660e00151876101000151886101400151829250819150995099509950995099509950995099509950509193959799909294969850565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080610da9610ad8565b1515610db457600080fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600a0160009054906101000a900460ff161515610e0f57600080fd5b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600a0160006101000a81548160ff021916908315150217905550610e73836115dd565b9050610e8a81600354611a6a90919063ffffffff16565b600381905550610ea581600454611a8b90919063ffffffff16565b6004819055508273ffffffffffffffffffffffffffffffffffffffff167f5b1c616c572bf4d007644e7a85cb76c1ad8ad649c2587712bc80495fe76343a6826040518082815260200191505060405180910390a26001915050919050565b6000600354905090565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610fcc57600080fd5b505af1158015610fe0573d6000803e3d6000fd5b505050506040513d6020811015610ff657600080fd5b8101908080519060200190929190505050905090565b600560009054906101000a900460ff1615151561102857600080fd5b611030610f0d565b60038190555060006004819055506001600560006101000a81548160ff0219169083151502179055507fd406aba267fc52b718da0f82dc70773ae07bd549015dbd3fe6858094443b95326003546040518082815260200191505060405180910390a1565b60008060006110a1610ad8565b15156110ac57600080fd5b600560009054906101000a900460ff1615156110c757600080fd5b6000841115156110d657600080fd5b6000851115156110e557600080fd5b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600a0160009054906101000a900460ff1615151561114157600080fd5b600354861115151561115257600080fd5b610160604051908101604052808873ffffffffffffffffffffffffffffffffffffffff168152602001878152602001600081526020014381526020016000815260200186815260200185815260200160006040519080825280602002602001820160405280156111d15781602001602082028038833980820191505090505b50815260200160006040519080825280602002602001820160405280156112075781602001602082028038833980820191505090505b508152602001600280549050815260200160011515815250600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070190805190602001906112fd929190611c57565b5061010082015181600801908051906020019061131b929190611cfd565b50610120820151816009015561014082015181600a0160006101000a81548160ff02191690831515021790555090505060028790806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600091505b848210156114d657600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206007016000908060018154018082558091505090600182039060005260206000209060209182820401919006909192909190916101000a81548160ff021916908315150217905550504361145a8386611aac90919063ffffffff16565b019050600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060080181908060018154018082558091505090600182039060005260206000200160009091929091909150555081806001019250506113b6565b6114eb86600354611a8b90919063ffffffff16565b60038190555061150686600454611a6a90919063ffffffff16565b6004819055508673ffffffffffffffffffffffffffffffffffffffff167f034f6128ac8e081f981ac7e4cfb7cdbee543c758847a73b9477271c1830ce786876040518082815260200191505060405180910390a28673ffffffffffffffffffffffffffffffffffffffff167f06528902871f4354a885a97043acf8286955068d9d7baf204f4ead3cafa2397b8686604051808381526020018281526020019250505060405180910390a2600192505050949350505050565b6115c6610ad8565b15156115d157600080fd5b6115da81611aea565b50565b60006115e7611be4565b600080600560009054906101000a900460ff1615151561160657600080fd5b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061016060405190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820180548060200260200160405190810160405280929190818152602001828054801561175857602002820191906000526020600020906000905b82829054906101000a900460ff161515815260200190600101906020826000010492830192600103820291508084116117225790505b50505050508152602001600882018054806020026020016040519081016040528092919081815260200182805480156117b057602002820191906000526020600020905b81548152602001906001019080831161179c575b5050505050815260200160098201548152602001600a820160009054906101000a900460ff161515151581525050925082610140015115156117f157600080fd5b60009150600090505b8260a00151811015611924578260e001518181518110151561181857fe5b90602001906020020151151561191757438361010001518281518110151561183c57fe5b906020019060200201511115156119115761186883602001518460400151611a6a90919063ffffffff16565b83604001818152505060018360e001518281518110151561188557fe5b90602001906020020190151590811515815250506118b0836020015183611a6a90919063ffffffff16565b91508473ffffffffffffffffffffffffffffffffffffffff167f987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a846020015183604051808381526020018281526020019250505060405180910390a2611916565b611924565b5b80806001019150506117fa565b600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156119e957600080fd5b505af11580156119fd573d6000803e3d6000fd5b505050506040513d6020811015611a1357600080fd5b81019080805190602001909291905050501515611a2f57600080fd5b611a608360400151611a528560a001518660200151611aac90919063ffffffff16565b611a8b90919063ffffffff16565b9350505050919050565b6000808284019050838110151515611a8157600080fd5b8091505092915050565b600080838311151515611a9d57600080fd5b82840390508091505092915050565b6000806000841415611ac15760009150611ae3565b8284029050828482811515611ad257fe5b04141515611adf57600080fd5b8091505b5092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611b2657600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61016060405190810160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016060815260200160608152602001600081526020016000151581525090565b82805482825590600052602060002090601f01602090048101928215611cec5791602002820160005b83821115611cbd57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302611c80565b8015611cea5782816101000a81549060ff0219169055600101602081600001049283019260010302611cbd565b505b509050611cf99190611d4a565b5090565b828054828255906000526020600020908101928215611d39579160200282015b82811115611d38578251825591602001919060010190611d1d565b5b509050611d469190611d7a565b5090565b611d7791905b80821115611d7357600081816101000a81549060ff021916905550600101611d50565b5090565b90565b611d9c91905b80821115611d98576000816000905550600101611d80565b5090565b905600a165627a7a72305820a2e24cffa14fc579fedae9f9d74b4209d65bf1e3d6730fab2cdd5737cc9749870029`

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

// ClaimState is a free data retrieval call binding the contract method 0x75be7f7f.
//
// Solidity: function claimState(account address, stage uint256) constant returns(uint8)
func (_BaasIncentives *BaasIncentivesCaller) ClaimState(opts *bind.CallOpts, account common.Address, stage *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BaasIncentives.contract.Call(opts, out, "claimState", account, stage)
	return *ret0, err
}

// ClaimState is a free data retrieval call binding the contract method 0x75be7f7f.
//
// Solidity: function claimState(account address, stage uint256) constant returns(uint8)
func (_BaasIncentives *BaasIncentivesSession) ClaimState(account common.Address, stage *big.Int) (uint8, error) {
	return _BaasIncentives.Contract.ClaimState(&_BaasIncentives.CallOpts, account, stage)
}

// ClaimState is a free data retrieval call binding the contract method 0x75be7f7f.
//
// Solidity: function claimState(account address, stage uint256) constant returns(uint8)
func (_BaasIncentives *BaasIncentivesCallerSession) ClaimState(account common.Address, stage *big.Int) (uint8, error) {
	return _BaasIncentives.Contract.ClaimState(&_BaasIncentives.CallOpts, account, stage)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, currentStage uint256, totalStages uint256, totalBlocks uint256, stagesClaimed bool[], stagesBlockTime uint256[], isValue bool)
func (_BaasIncentives *BaasIncentivesCaller) GetIncentive(opts *bind.CallOpts, account common.Address) (struct {
	InitialAmount   *big.Int
	Withdrawn       *big.Int
	AtBlock         *big.Int
	CurrentStage    *big.Int
	TotalStages     *big.Int
	TotalBlocks     *big.Int
	StagesClaimed   []bool
	StagesBlockTime []*big.Int
	IsValue         bool
}, error) {
	ret := new(struct {
		InitialAmount   *big.Int
		Withdrawn       *big.Int
		AtBlock         *big.Int
		CurrentStage    *big.Int
		TotalStages     *big.Int
		TotalBlocks     *big.Int
		StagesClaimed   []bool
		StagesBlockTime []*big.Int
		IsValue         bool
	})
	out := ret
	err := _BaasIncentives.contract.Call(opts, out, "getIncentive", account)
	return *ret, err
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, currentStage uint256, totalStages uint256, totalBlocks uint256, stagesClaimed bool[], stagesBlockTime uint256[], isValue bool)
func (_BaasIncentives *BaasIncentivesSession) GetIncentive(account common.Address) (struct {
	InitialAmount   *big.Int
	Withdrawn       *big.Int
	AtBlock         *big.Int
	CurrentStage    *big.Int
	TotalStages     *big.Int
	TotalBlocks     *big.Int
	StagesClaimed   []bool
	StagesBlockTime []*big.Int
	IsValue         bool
}, error) {
	return _BaasIncentives.Contract.GetIncentive(&_BaasIncentives.CallOpts, account)
}

// GetIncentive is a free data retrieval call binding the contract method 0x8f70c2d8.
//
// Solidity: function getIncentive(account address) constant returns(initialAmount uint256, withdrawn uint256, atBlock uint256, currentStage uint256, totalStages uint256, totalBlocks uint256, stagesClaimed bool[], stagesBlockTime uint256[], isValue bool)
func (_BaasIncentives *BaasIncentivesCallerSession) GetIncentive(account common.Address) (struct {
	InitialAmount   *big.Int
	Withdrawn       *big.Int
	AtBlock         *big.Int
	CurrentStage    *big.Int
	TotalStages     *big.Int
	TotalBlocks     *big.Int
	StagesClaimed   []bool
	StagesBlockTime []*big.Int
	IsValue         bool
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

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_BaasIncentives *BaasIncentivesTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_BaasIncentives *BaasIncentivesSession) Claim() (*types.Transaction, error) {
	return _BaasIncentives.Contract.Claim(&_BaasIncentives.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_BaasIncentives *BaasIncentivesTransactorSession) Claim() (*types.Transaction, error) {
	return _BaasIncentives.Contract.Claim(&_BaasIncentives.TransactOpts)
}

// Forfeited is a paid mutator transaction binding the contract method 0x9de59fad.
//
// Solidity: function forfeited(account address) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Forfeited(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "forfeited", account)
}

// Forfeited is a paid mutator transaction binding the contract method 0x9de59fad.
//
// Solidity: function forfeited(account address) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Forfeited(account common.Address) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Forfeited(&_BaasIncentives.TransactOpts, account)
}

// Forfeited is a paid mutator transaction binding the contract method 0x9de59fad.
//
// Solidity: function forfeited(account address) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Forfeited(account common.Address) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Forfeited(&_BaasIncentives.TransactOpts, account)
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

// Reward is a paid mutator transaction binding the contract method 0xc5a3487c.
//
// Solidity: function reward(account address, amountPerStage uint256, totalVestingStages uint256, blockTimePerStage uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactor) Reward(opts *bind.TransactOpts, account common.Address, amountPerStage *big.Int, totalVestingStages *big.Int, blockTimePerStage *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.contract.Transact(opts, "reward", account, amountPerStage, totalVestingStages, blockTimePerStage)
}

// Reward is a paid mutator transaction binding the contract method 0xc5a3487c.
//
// Solidity: function reward(account address, amountPerStage uint256, totalVestingStages uint256, blockTimePerStage uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesSession) Reward(account common.Address, amountPerStage *big.Int, totalVestingStages *big.Int, blockTimePerStage *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Reward(&_BaasIncentives.TransactOpts, account, amountPerStage, totalVestingStages, blockTimePerStage)
}

// Reward is a paid mutator transaction binding the contract method 0xc5a3487c.
//
// Solidity: function reward(account address, amountPerStage uint256, totalVestingStages uint256, blockTimePerStage uint256) returns(bool)
func (_BaasIncentives *BaasIncentivesTransactorSession) Reward(account common.Address, amountPerStage *big.Int, totalVestingStages *big.Int, blockTimePerStage *big.Int) (*types.Transaction, error) {
	return _BaasIncentives.Contract.Reward(&_BaasIncentives.TransactOpts, account, amountPerStage, totalVestingStages, blockTimePerStage)
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

// BaasIncentivesClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the BaasIncentives contract.
type BaasIncentivesClaimedIterator struct {
	Event *BaasIncentivesClaimed // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesClaimed)
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
		it.Event = new(BaasIncentivesClaimed)
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
func (it *BaasIncentivesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesClaimed represents a Claimed event raised by the BaasIncentives contract.
type BaasIncentivesClaimed struct {
	Account common.Address
	Amount  *big.Int
	Stage   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: e Claimed(account indexed address, amount uint256, stage uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterClaimed(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "Claimed", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesClaimedIterator{contract: _BaasIncentives.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: e Claimed(account indexed address, amount uint256, stage uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BaasIncentivesClaimed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "Claimed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesClaimed)
				if err := _BaasIncentives.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// BaasIncentivesForfeitedIterator is returned from FilterForfeited and is used to iterate over the raw logs and unpacked data for Forfeited events raised by the BaasIncentives contract.
type BaasIncentivesForfeitedIterator struct {
	Event *BaasIncentivesForfeited // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesForfeitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesForfeited)
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
		it.Event = new(BaasIncentivesForfeited)
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
func (it *BaasIncentivesForfeitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesForfeitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesForfeited represents a Forfeited event raised by the BaasIncentives contract.
type BaasIncentivesForfeited struct {
	Account        common.Address
	RemainingToken *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterForfeited is a free log retrieval operation binding the contract event 0x5b1c616c572bf4d007644e7a85cb76c1ad8ad649c2587712bc80495fe76343a6.
//
// Solidity: e Forfeited(account indexed address, remainingToken uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterForfeited(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesForfeitedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "Forfeited", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesForfeitedIterator{contract: _BaasIncentives.contract, event: "Forfeited", logs: logs, sub: sub}, nil
}

// WatchForfeited is a free log subscription operation binding the contract event 0x5b1c616c572bf4d007644e7a85cb76c1ad8ad649c2587712bc80495fe76343a6.
//
// Solidity: e Forfeited(account indexed address, remainingToken uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchForfeited(opts *bind.WatchOpts, sink chan<- *BaasIncentivesForfeited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "Forfeited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesForfeited)
				if err := _BaasIncentives.contract.UnpackLog(event, "Forfeited", log); err != nil {
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

// BaasIncentivesIncentiveProvided1Iterator is returned from FilterIncentiveProvided1 and is used to iterate over the raw logs and unpacked data for IncentiveProvided1 events raised by the BaasIncentives contract.
type BaasIncentivesIncentiveProvided1Iterator struct {
	Event *BaasIncentivesIncentiveProvided1 // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesIncentiveProvided1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesIncentiveProvided1)
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
		it.Event = new(BaasIncentivesIncentiveProvided1)
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
func (it *BaasIncentivesIncentiveProvided1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesIncentiveProvided1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesIncentiveProvided1 represents a IncentiveProvided1 event raised by the BaasIncentives contract.
type BaasIncentivesIncentiveProvided1 struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIncentiveProvided1 is a free log retrieval operation binding the contract event 0x034f6128ac8e081f981ac7e4cfb7cdbee543c758847a73b9477271c1830ce786.
//
// Solidity: e IncentiveProvided1(account indexed address, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterIncentiveProvided1(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesIncentiveProvided1Iterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "IncentiveProvided1", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesIncentiveProvided1Iterator{contract: _BaasIncentives.contract, event: "IncentiveProvided1", logs: logs, sub: sub}, nil
}

// WatchIncentiveProvided1 is a free log subscription operation binding the contract event 0x034f6128ac8e081f981ac7e4cfb7cdbee543c758847a73b9477271c1830ce786.
//
// Solidity: e IncentiveProvided1(account indexed address, amount uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchIncentiveProvided1(opts *bind.WatchOpts, sink chan<- *BaasIncentivesIncentiveProvided1, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "IncentiveProvided1", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesIncentiveProvided1)
				if err := _BaasIncentives.contract.UnpackLog(event, "IncentiveProvided1", log); err != nil {
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

// BaasIncentivesIncentiveProvided2Iterator is returned from FilterIncentiveProvided2 and is used to iterate over the raw logs and unpacked data for IncentiveProvided2 events raised by the BaasIncentives contract.
type BaasIncentivesIncentiveProvided2Iterator struct {
	Event *BaasIncentivesIncentiveProvided2 // Event containing the contract specifics and raw log

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
func (it *BaasIncentivesIncentiveProvided2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaasIncentivesIncentiveProvided2)
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
		it.Event = new(BaasIncentivesIncentiveProvided2)
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
func (it *BaasIncentivesIncentiveProvided2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaasIncentivesIncentiveProvided2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaasIncentivesIncentiveProvided2 represents a IncentiveProvided2 event raised by the BaasIncentives contract.
type BaasIncentivesIncentiveProvided2 struct {
	Account       common.Address
	VestingStages *big.Int
	VestingBlocks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterIncentiveProvided2 is a free log retrieval operation binding the contract event 0x06528902871f4354a885a97043acf8286955068d9d7baf204f4ead3cafa2397b.
//
// Solidity: e IncentiveProvided2(account indexed address, vestingStages uint256, vestingBlocks uint256)
func (_BaasIncentives *BaasIncentivesFilterer) FilterIncentiveProvided2(opts *bind.FilterOpts, account []common.Address) (*BaasIncentivesIncentiveProvided2Iterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.FilterLogs(opts, "IncentiveProvided2", accountRule)
	if err != nil {
		return nil, err
	}
	return &BaasIncentivesIncentiveProvided2Iterator{contract: _BaasIncentives.contract, event: "IncentiveProvided2", logs: logs, sub: sub}, nil
}

// WatchIncentiveProvided2 is a free log subscription operation binding the contract event 0x06528902871f4354a885a97043acf8286955068d9d7baf204f4ead3cafa2397b.
//
// Solidity: e IncentiveProvided2(account indexed address, vestingStages uint256, vestingBlocks uint256)
func (_BaasIncentives *BaasIncentivesFilterer) WatchIncentiveProvided2(opts *bind.WatchOpts, sink chan<- *BaasIncentivesIncentiveProvided2, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BaasIncentives.contract.WatchLogs(opts, "IncentiveProvided2", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaasIncentivesIncentiveProvided2)
				if err := _BaasIncentives.contract.UnpackLog(event, "IncentiveProvided2", log); err != nil {
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
