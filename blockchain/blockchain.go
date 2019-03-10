package blockchain

import (
	"os"

	"github.com/CAU-CLINK/blockchain_with_go/common"
)

type Blockchain struct {
	db Database
}

func New(databasePath string, genesisConfFilePath string) (*Blockchain, error) {
	if !IsExists(genesisConfFilePath) {
		return CreateBlockChain(databasePath, genesisConfFilePath)
	}

	db, err := NewLevelDB(databasePath)
	if err != nil {
		return nil, err
	}

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
	serializedGenesisBlock, err := common.Serialize(genesisBlock)
	if err != nil {
		return nil, err
	}

	db.Put(genesisBlockHash, serializedGenesisBlock)
	db.Put([]byte("1"), genesisBlockHash)

	blockchain := &Blockchain{db: db}

	return blockchain, nil
}

// TODO : implements me with test case
func (b *Blockchain) AddBlock(block *Block) error {
	serializedBlock, err := common.Serialize(block)
	if err != nil {
		return err
	}

	err = b.db.Put(block.Hash(), serializedBlock)
	if err != nil {
		return err
	}

	lastHash, err := b.db.Tip()
	if err != nil {
		return err
	}

	lastSerializedBlock, err := b.db.Get(lastHash)
	if err != nil {
		return err
	}

	// need to validate of new block with lastBlock!
	// ex> height, hash, etc...
	var lastBlock *Block = &Block{}
	err = common.Deserialize(lastSerializedBlock, lastBlock)
	if err != nil {
		return err
	}

	err = b.db.Put([]byte("1"), block.Hash())

	return nil
}

// TODO : implements me with test case
func (b *Blockchain) IsValid(block *Block) bool {
	return false
}

func IsExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
