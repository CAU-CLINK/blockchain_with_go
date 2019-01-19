package blockchain

type BlockHeader struct {
	Version           int32
	PreviousBlockHash []byte
	MerkleRootHash    []byte
	Timestamp         uint64
	Nonce             uint64
}

type Block struct {
	Header       BlockHeader
	Transactions []*Transaction
}

func CreateNewBlock(transactions []*Transaction) *Block {
	return nil
}

func CreateGenesisBlock(genesisconfFilePath string) (*Block, error) {
	return nil, nil
}