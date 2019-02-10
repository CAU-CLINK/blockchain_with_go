package blockchain_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
)

func TestCreateGenesisBlock(t *testing.T) {
	genesisConfigFilePath := "../db/genesis.json"

	genesisBlock, err := blockchain.CreateGenesisBlock(genesisConfigFilePath)
	if err != nil {
		t.Error(err)
	}

	if genesisBlock == nil {
		t.Error("Fail to create genesis block")
	}
}

// TODO : Implements test case
func TestSetHeader(t *testing.T) {

}

// TODO : Implements test case
func TestHash(t *testing.T) {

}
