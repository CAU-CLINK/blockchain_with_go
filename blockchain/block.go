package blockchain

//uint to int
type BlockHeader struct {
	Version           int32 // original bitcoin spec, use little endian format
	PreviousBlockHash []byte // 32bytes
	MerkleRootHash    []byte // 32bytes
	Timestamp         int32 // original bitcoin spec
	Nonce             int32 // original bitcoin spec
	Bits			  int32 // original bitcoin spec
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