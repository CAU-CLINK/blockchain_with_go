package blockchain

import (
	"os"

	"github.com/CAU-CLINK/blockchain_with_go/util"
)

type Blockchain struct {
	db Database
}

func NewBlockChain(databasePath string, genesisConfFilePath string) (*Blockchain, error) {
	if !IsExists(genesisConfFilePath) {
		return CreateBlockChain(databasePath, genesisConfFilePath)
	}

	db, err := NewLevelDB(databasePath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return &Blockchain{
		db: db,
	}, nil
}

func CreateBlockChain(databasePath string, genesisConfFilePath string) (*Blockchain, error) {
	genesisBlock, err := CreateGenesisBlock(genesisConfFilePath)
	if err != nil {
		return nil, err
	}

	db, err := NewLevelDB(databasePath)
	if err != nil {
		return nil, err
	}

	genesisBlockHash := genesisBlock.Hash()
	serializedGenesisBlock, err := util.Serialize(genesisBlock)
	if err != nil {
		return nil, err
	}

	db.Put(genesisBlockHash, serializedGenesisBlock)
	db.Put([]byte("1"), genesisBlockHash)

	blockchain := &Blockchain{db: db}

	return blockchain, nil
}

// TODO : implements me with test case
func (b *Blockchain) AddBlock(block *Block) {

}

func IsExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
