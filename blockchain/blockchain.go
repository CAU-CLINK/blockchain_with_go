package blockchain

type Blockchain struct {
	db *Database
}

func NewBlockChain(db *Database) *Blockchain {
	return &Blockchain{
		db: db,
	}
}

// TODO : implement this with testcase
func (b *Blockchain) AddBlock(block *Block) {

}
