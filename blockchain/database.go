package blockchain

import "github.com/syndtr/goleveldb/leveldb"

type Database interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Tip() ([]byte, error)
}

type levelDB struct {
	db *leveldb.DB
}

func(db *levelDB) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func(db *levelDB) Put(key []byte, value []byte) error {
	return nil
}

func(db *levelDB) Tip() ([]byte, error) {
	return nil, nil
}
