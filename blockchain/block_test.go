package blockchain_test

import (
	"bytes"
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
	tests := []struct {
		Version           int32
		PreviousBlockHash []byte
		MerkleRootHash    []byte
		expectedHeader    *blockchain.BlockHeader
	}{
		{
			32,
			[]byte{0x86},
			[]byte{0x35},
			blockchain.SetHeader(32, []byte{0x86}, []byte{0x35}),
		},
		{
			22,
			[]byte{0x43},
			[]byte{0x43},
			blockchain.SetHeader(22, []byte{0x43}, []byte{0x43}),
		},
	}
	for _, test := range tests {
		if test.expectedHeader.Version != test.Version {
			t.Error("Error at test Version!")
		}
		if !bytes.Equal(test.expectedHeader.MerkleRootHash, test.MerkleRootHash) {
			t.Error("Error at test MerkleRootHash")
		}
		if !bytes.Equal(test.expectedHeader.PreviousBlockHash, test.PreviousBlockHash) {
			t.Error("Error at test PreviousBlockHash")
		}
	}
}

// TODO : Implements test case
func TestHash(t *testing.T) {
	tests := []struct {
		blockheader *blockchain.BlockHeader
	}{
		{
			blockchain.SetHeader(32, nil, nil),
		},
		{
			blockchain.SetHeader(32, []byte{0x22}, []byte{0x33}),
		},
	}
	for _, test := range tests {
		block := blockchain.Block{test.blockheader, nil}
		if len(block.Hash()) != 32 {
			t.Error("Error at test Hash")
		}
	}
}
