package deployer

import (
	"fmt"
	"github.com/baas/tge-sol/interactor/contracts"
	"github.com/baas/tge-sol/interactor/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"math/big"
	"time"
)

func process(httpPath string, keystoreUTCPath string, passwordFile string, version string) error {
	log.Println("Deploying contracts to ", httpPath)
	client, err := utils.RPCClient(httpPath)

	if err != nil {
		return fmt.Errorf("dial client %v", err.Error())
	}

	privKey, err := utils.PrivateKeyFromWalletAndPasswordFile(keystoreUTCPath, passwordFile)

	if err != nil {
		return fmt.Errorf("read private key %v", err.Error())
	}

	txOps := bind.NewKeyedTransactor(privKey.PrivateKey)
	txOps.Value = big.NewInt(0)

	// DEPLOY
	cc, err := deployTGEContracts(txOps, client)

	if err != nil {
		return fmt.Errorf("deploy contracts %v", err.Error())
	}

	// export configuration
	return cc.ExportResult( version)
}


func deployTGEContracts(txOps *bind.TransactOpts, client *ethclient.Client) (*ContractConfig, error) {
	var err error

	cc := &ContractConfig{
		Deployer: &txOps.From,
	}

	cc.TokenAddress, err = deployBaasToken(txOps, client, cc)
	if err != nil {
		return nil, err
	}

	cc.FounderAddress, err = deployBaasFounder(txOps, client, cc)
	if err != nil {
		return nil, err
	}

	cc.IncentivesAddress, err = deployBaasIncentives(txOps, client, cc)
	if err != nil {
		return nil, err
	}

	cc.PPAddress, err = DeployBaasPP(txOps, client, cc)
	if err != nil {
		return nil, err
	}

	cc.EscrowAddress, err = deployBaasEscrow(txOps, client, cc)
	if err != nil {
		return nil, err
	}

	cc.ROIAddress, err = deployBaasROI(txOps, client, cc)
	if err != nil {
		log.Fatal(err)
	}

	return cc, nil
}

func deployBaasToken(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasToken(txOps, client)

	if err != nil {
		return nil, err
	}

	cc.TokenHash = tx.Hash()
	return deploy("Token", txOps, client, tx, &addr)
}

func deployBaasFounder(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasFounder(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}
	cc.FounderHash = tx.Hash()

	return deploy("Founder", txOps, client, tx, &addr)
}

func deployBaasEscrow(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasEscrow(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	cc.EscrowHash = tx.Hash()
	return deploy("Escrow", txOps, client, tx, &addr)
}

func deployBaasIncentives(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasIncentive(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	cc.IncentivesHash = tx.Hash()
	return deploy("Incentives", txOps, client, tx, &addr)
}

func DeployBaasPP(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasPP(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	cc.PPHash = tx.Hash()
	return deploy("Private Placement", txOps, client, tx, &addr)
}

func deployBaasROI(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasROI(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	cc.ROIHash = tx.Hash()
	return deploy("ROI", txOps, client, tx, &addr)
}

func deploy(label string, txOps *bind.TransactOpts, client *ethclient.Client, tx *types.Transaction, addr *common.Address) (*common.Address, error) {
	var err error
	log.Println("Deploying", label, addr.String(), tx.Hash().String())
	for pending := true; pending; _, pending, err = client.TransactionByHash(context.Background(), tx.Hash()) {
		if err != nil {
			return nil, err
		}
		time.Sleep(2 * time.Second)
	}

	log.Println(label, "Deployed, Address: ", addr.String())

	return addr, nil
}
