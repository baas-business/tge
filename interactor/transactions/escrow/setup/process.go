package setup

import (
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(dAppContext *web3.DAppContext, args *CommandArgs) error {
	txOps := bind.NewKeyedTransactor(dAppContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasEscrow(tge.EscrowAddress(dAppContext), dAppContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Setup(txOps, new(big.Int).SetInt64(args.Vesting))

	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("Setup Escrow", dAppContext.Client, tx)
}
