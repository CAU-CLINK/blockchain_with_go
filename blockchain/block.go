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
	transactions Transaction
}
