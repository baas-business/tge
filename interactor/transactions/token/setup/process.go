package setup

import (
	"github.com/baas-business/tge-sol/interactor/contracts"
	"github.com/baas-business/tge-sol/interactor/tge"
	"github.com/ellsol/solidity-tools/utils/web3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func process(dappContext *web3.DAppContext) error {
	txOps := bind.NewKeyedTransactor(dappContext.Key.PrivateKey)
	txOps.Value = big.NewInt(0)

	contract, err := contracts.NewBaasToken(tge.TokenAddress(dappContext), dappContext.Client)

	if err != nil {
		return err
	}

	tx, err := contract.Setup(txOps,
		tge.EscrowAddress(dappContext),
		tge.PPAddress(dappContext),
		tge.FounderAddress(dappContext),
		tge.IncentivesAddress(dappContext))


	if err != nil {
		return err
	}

	return web3.ExecuteTransaction("Setup Token", dappContext.Client, tx)
}
