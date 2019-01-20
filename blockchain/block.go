package blockchain

import "time"

type BlockHeader struct {
	Version           int32
	PreviousBlockHash []byte
	MerkleRootHash    []byte
	Timestamp         int64
	Nonce             int64
}

type Block struct {
	Header       *BlockHeader
	Transactions []*Transaction
}

// TODO : implement this (with pow) with testcase
func CreateNewBlock(blockHeader *BlockHeader, transactions []*Transaction) *Block {
	block := &Block{blockHeader, transactions}
	return block
}

// TODO : implement this with testcase
func CreateGenesisBlock(genesisconfFilePath string) (*Block, error) {
	return nil, nil
}

func SetHeader(version int32, previousBlockHash []byte, merkleRootHash []byte) *BlockHeader {
	blockHeader := &BlockHeader{version, previousBlockHash,
		merkleRootHash, time.Now().Unix(), 0}

	return blockHeader
}
