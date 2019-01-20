package blockchain

import (
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
)

type Database interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Close()
	Tip() ([]byte, error)
}

type LevelDB struct {
	sync.RWMutex
	db *leveldb.DB
}

func NewLevelDB(dbPath string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}

	return &LevelDB{
		db: db,
	}, nil
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	db.Lock()
	defer db.Unlock()

	data, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *LevelDB) Put(key []byte, value []byte) error {
	db.Lock()
	defer db.Unlock()

	err := db.db.Put(key, value, nil)
	if err != nil {
		return err
	}
	return nil
}

func (db *LevelDB) Close() {
	db.db.Close()
}

func (db *LevelDB) Tip() ([]byte, error) {
	tip, err := db.Get([]byte("1"))
	if err != nil {
		return nil, err
	}
	return tip, nil
}
