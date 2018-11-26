package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/baas/tge-sol/deployer/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"time"
)

func DeployTGEContracts(txOps *bind.TransactOpts, client *ethclient.Client) (*ContractConfig, error) {
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

	return deploy("Token", txOps, client, tx, &addr)
}

func deployBaasFounder(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasFounder(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Founder", txOps, client, tx, &addr)
}

func deployBaasEscrow(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasEscrow(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Escrow", txOps, client, tx, &addr)
}

func deployBaasIncentives(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasIncentive(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Incentives", txOps, client, tx, &addr)
}

func DeployBaasPP(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasPP(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

	return deploy("Private Placement", txOps, client, tx, &addr)
}

func deployBaasROI(txOps *bind.TransactOpts, client *ethclient.Client, cc *ContractConfig) (*common.Address, error) {
	addr, tx, _, err := contracts.DeployBaasROI(txOps, client, *cc.TokenAddress)

	if err != nil {
		return nil, err
	}

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

type ContractConfig struct {
	Deployer          *common.Address `json:"deployer"`
	TokenAddress      *common.Address `json:"token_address"`
	EscrowAddress     *common.Address `json:"escrow_address"`
	FounderAddress    *common.Address `json:"founder_address"`
	IncentivesAddress *common.Address `json:"incentives_address"`
	PPAddress         *common.Address `json:"pp_address"`
	ROIAddress        *common.Address `json:"roi_address"`
}

func (cc *ContractConfig) Write(filename string) error {
	j, err := json.Marshal(cc)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, j, 0644)
	fmt.Printf("%+v", cc)
	return nil
}

func LoadContractConfig(filename string) (*ContractConfig, error) {
	f, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	raw := &ContractConfigRaw{}
	err = json.Unmarshal(f, raw)
	if err != nil {
		return nil, err
	}

	deployer := common.HexToAddress(raw.Deployer)
	token := common.HexToAddress(raw.TokenAddress)
	escrow := common.HexToAddress(raw.EscrowAddress)
	founder := common.HexToAddress(raw.FounderAddress)
	pp := common.HexToAddress(raw.PPAddress)
	incentives := common.HexToAddress(raw.IncentivesAddress)
	roi := common.HexToAddress(raw.ROIAddress)

	return &ContractConfig{
		Deployer:          &deployer,
		TokenAddress:      &token,
		EscrowAddress:     &escrow,
		FounderAddress:    &founder,
		IncentivesAddress: &incentives,
		PPAddress:         &pp,
		ROIAddress:        &roi,
	}, nil
}

type ContractConfigRaw struct {
	Deployer          string `json:"deployer"`
	TokenAddress      string `json:"token_address"`
	EscrowAddress     string `json:"escrow_address"`
	FounderAddress    string `json:"founder_address"`
	IncentivesAddress string `json:"incentives_address"`
	PPAddress         string `json:"pp_address"`
	ROIAddress        string `json:"roi_address"`
}

func (cc *ContractConfig) String() string{
	bb := bytes.Buffer{}

	bb.WriteString("Deployer:   ")
	bb.WriteString(cc.Deployer.String())
	bb.WriteString("\n")

	bb.WriteString("Token:      ")
	bb.WriteString(cc.TokenAddress.String())
	bb.WriteString("\n")

	bb.WriteString("Escrow:     ")
	bb.WriteString(cc.EscrowAddress.String())
	bb.WriteString("\n")

	bb.WriteString("Founder:    ")
	bb.WriteString(cc.FounderAddress.String())
	bb.WriteString("\n")

	bb.WriteString("Incentives: ")
	bb.WriteString(cc.IncentivesAddress.String())
	bb.WriteString("\n")

	bb.WriteString("Private P.: ")
	bb.WriteString(cc.PPAddress.String())
	bb.WriteString("\n")

	bb.WriteString("ROI:        ")
	bb.WriteString(cc.ROIAddress.String())
	bb.WriteString("\n")

	return bb.String()
}
