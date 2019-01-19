package blockchain

import (
	"sync"
)

type Blockchain struct {
	m *sync.RWMutex
	db *Database
}

func NewBlockChain(db *Database) *Blockchain {
	return &Blockchain{
		db:db,
	}
}

func (b *Blockchain) AddBlock(block *Block){

}