package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ellsol/etherlytics-api/config/dapp"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"log"
	"os"
)

type ContractConfig struct {
	Deployer          *common.Address `json:"deployer"`
	Block             int64           `json:"block"`
	TokenAddress      *common.Address `json:"token_address"`
	TokenHash         common.Hash     `json:"token_hash"`
	EscrowAddress     *common.Address `json:"escrow_address"`
	EscrowHash        common.Hash     `json:"escrow_hash"`
	FounderAddress    *common.Address `json:"founder_address"`
	FounderHash       common.Hash     `json:"founder_hash"`
	IncentivesAddress *common.Address `json:"incentives_address"`
	IncentivesHash    common.Hash     `json:"incentives_hash"`
	PPAddress         *common.Address `json:"pp_address"`
	PPHash            common.Hash     `json:"pp_hash"`
	ROIAddress        *common.Address `json:"roi_address"`
	ROIHash           common.Hash     `json:"roi_hash"`
}

func (cc *ContractConfig) ExportResult(version string) error {
	path := fmt.Sprintf("build/%v", version)


	if _, err := os.Stat(path); !os.IsNotExist(err) {
		version = version + "_alternative"
		path = fmt.Sprintf("build/%v", version)
	}


	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}

	err = cc.Write(fmt.Sprintf("%v/contract_config.json", path))

	if err != nil {
		return err
	}

	// write dapp.json for Etherlytics
	return cc.DAppConfig().Write(fmt.Sprintf("%v/dapp.json", path))
}

func (cc *ContractConfig) DAppConfig() *dapp.DAppRaw {
	log.Println("Writing dapp config")
	da := &dapp.DAppRaw{}

	da.ID = "TGE"
	da.Name = "TGE"
	da.CurrentBlock = cc.Block

	contracts := make([]dapp.ContractRaw, 0)

	contracts = append(contracts, createContractRaw("BaasToken", cc.TokenAddress.String(), cc.TokenHash.String(), 0))
	contracts = append(contracts, createContractRaw("BaasEscrow", cc.EscrowAddress.String(), cc.EscrowHash.String(), 0))
	contracts = append(contracts, createContractRaw("BaasFounder", cc.FounderAddress.String(), cc.FounderHash.String(), 0))
	contracts = append(contracts, createContractRaw("BaasIncentives", cc.IncentivesAddress.String(), cc.IncentivesHash.String(), 0))
	contracts = append(contracts, createContractRaw("BaasPP", cc.PPAddress.String(), cc.PPHash.String(), 0))
	contracts = append(contracts, createContractRaw("BaasROI", cc.ROIAddress.String(), cc.ROIHash.String(), 0))

	da.Contracts = contracts

	return da
}

func createContractRaw(name string, address string, hash string, typemask int) dapp.ContractRaw {
	return dapp.ContractRaw{
		Name:                   name,
		Address:                address,
		GenesisTransactionHash: hash,
		Symbol:                 name,
		AbiFilename:            "dapp/" + name + ".json",
		Type:                   typemask,
		DAppID:                 "TGE",
		SourceFilename:         "dapp/" + name + ".sol",
	}
}

func (cc *ContractConfig) Write(filename string) error {
	j, err := json.Marshal(cc)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, j, 0644)
	fmt.Printf("%+v", cc)
	return err
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

	tokenHash := common.HexToHash(raw.TokenHash)
	escrowHash := common.HexToHash(raw.EscrowHash)
	founderHash := common.HexToHash(raw.FounderHash)
	ppHash := common.HexToHash(raw.PPHash)
	incentivesHash := common.HexToHash(raw.IncentivesHash)
	roiHash := common.HexToHash(raw.ROIHash)

	return &ContractConfig{
		Deployer:          &deployer,
		Block:             raw.Block,
		TokenAddress:      &token,
		TokenHash:         tokenHash,
		EscrowAddress:     &escrow,
		EscrowHash:        escrowHash,
		FounderAddress:    &founder,
		FounderHash:       founderHash,
		IncentivesAddress: &incentives,
		IncentivesHash:    incentivesHash,
		PPAddress:         &pp,
		PPHash:            ppHash,
		ROIAddress:        &roi,
		ROIHash:           roiHash,
	}, nil
}

type ContractConfigRaw struct {
	Deployer          string `json:"deployer"`
	Block             int64  `json:"block"`
	TokenAddress      string `json:"token_address"`
	TokenHash         string `json:"token_hash"`
	EscrowAddress     string `json:"escrow_address"`
	EscrowHash        string `json:"escrow_hash"`
	FounderAddress    string `json:"founder_address"`
	FounderHash       string `json:"founder_hash"`
	IncentivesAddress string `json:"incentives_address"`
	IncentivesHash    string `json:"incentives_hash"`
	PPAddress         string `json:"pp_address"`
	PPHash            string `json:"pp_hash"`
	ROIAddress        string `json:"roi_address"`
	ROIHash           string `json:"roi_hash"`
}

func (cc *ContractConfig) String() string {
	bb := bytes.Buffer{}

	bb.WriteString("Deployer:   ")
	bb.WriteString(cc.Deployer.String())
	bb.WriteString("\n")
	bb.WriteString(string(cc.Block))
	bb.WriteString("\n")

	bb.WriteString("Token:      ")
	bb.WriteString(cc.TokenAddress.String())
	bb.WriteString("\n")
	bb.WriteString(cc.TokenHash.String())
	bb.WriteString("\n")

	bb.WriteString("Escrow:     ")
	bb.WriteString(cc.EscrowAddress.String())
	bb.WriteString("\n")
	bb.WriteString(cc.EscrowHash.String())
	bb.WriteString("\n")

	bb.WriteString("Founder:    ")
	bb.WriteString(cc.FounderAddress.String())
	bb.WriteString("\n")
	bb.WriteString(cc.FounderHash.String())
	bb.WriteString("\n")

	bb.WriteString("Incentives: ")
	bb.WriteString(cc.IncentivesAddress.String())
	bb.WriteString("\n")
	bb.WriteString(cc.IncentivesHash.String())
	bb.WriteString("\n")

	bb.WriteString("Private P.: ")
	bb.WriteString(cc.PPAddress.String())
	bb.WriteString("\n")
	bb.WriteString(cc.PPHash.String())
	bb.WriteString("\n")

	bb.WriteString("ROI:        ")
	bb.WriteString(cc.ROIAddress.String())
	bb.WriteString("\n")
	bb.WriteString(cc.ROIHash.String())
	bb.WriteString("\n")

	return bb.String()
}
