package blockchain_test

import (
	"os"
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
)

// TODO : Add more test case
func TestNewBlockChain(t *testing.T) {
	genesisConfFilePath := "../db/test/genesis.json"
	databasePath := "../db/test/block"

	defer os.RemoveAll(databasePath)

	_, err := blockchain.New(databasePath, genesisConfFilePath)
	if err != nil {
		t.Error(err)
	}
}

// TODO : Implements test case
func TestBlockchain_AddBlock(t *testing.T) {

}
