package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type BlockHeader struct {
	Version           int32
	PreviousBlockHash []byte
	MerkleRootHash    []byte
	Timestamp         int64
	Nonce             int64
	Bits              int32
}

type Block struct {
	Header       *BlockHeader
	Transactions []*Transaction
}

// TODO : implement this (with pow) with test case
func CreateNewBlock(blockHeader *BlockHeader, transactions []*Transaction) *Block {
	block := &Block{blockHeader, transactions}
	return block
}

// TODO : test case
func CreateGenesisBlock(genesisConfFilePath string) (*Block, error) {
	jsonFile, err := os.Open(genesisConfFilePath)
	defer jsonFile.Close()

	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	genesisHeader := &BlockHeader{}

	err = json.Unmarshal(byteValue, genesisHeader)
	if err != nil {
		return nil, err
	}

	genesisHeader.Timestamp = time.Now().Unix()

	genesisBlock := &Block{}
	genesisBlock.Header = genesisHeader

	return genesisBlock, nil
}

func SetHeader(version int32, previousBlockHash []byte, merkleRootHash []byte) *BlockHeader {
	blockHeader := &BlockHeader{version, previousBlockHash,
		merkleRootHash, time.Now().Unix(), 0, 0}

	return blockHeader
}

// TODO : Implements me with test case
func (b *Block) Hash() []byte {
	return nil
}
