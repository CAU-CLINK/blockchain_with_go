package blockchain_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
)

func TestLevelDB_New(t *testing.T) {
	filePath := "../db/test"

	levelDB, err := blockchain.NewLevelDB(filePath)
	defer os.RemoveAll(filePath)

	if err != nil {
		t.Error("Fail to open the testDB")
	}

	levelDB.Close()
}

func TestLevelDB_Put(t *testing.T) {
	filePath := "../db/test"

	levelDB, err := blockchain.NewLevelDB(filePath)
	defer os.RemoveAll(filePath)

	if err != nil {
		t.Error("Fail to open the testDB")
	}

	key := []byte("key")
	value := []byte("value")

	err = levelDB.Put(key, value)
	if err != nil {
		t.Error("Fail to put the DB")
	}

	levelDB.Close()
}

func TestLevelDB_Get(t *testing.T) {
	filePath := "../db/test"

	levelDB, err := blockchain.NewLevelDB(filePath)
	defer os.RemoveAll(filePath)

	if err != nil {
		t.Error("Fail to open the testDB")
	}

	key := []byte("key")
	value := []byte("value")

	err = levelDB.Put(key, value)
	if err != nil {
		t.Error("Fail to put the data")
	}

	data, err := levelDB.Get(key)
	if err != nil {
		t.Error("Fail to get the data")
	}

	if !bytes.Equal(value, data) {
		t.Error("Get invalid data")
	}

	levelDB.Close()
}

func TestLevelDB_Tip(t *testing.T) {
	filePath := "../db/test"

	levelDB, err := blockchain.NewLevelDB(filePath)
	defer os.RemoveAll(filePath)

	if err != nil {
		t.Error("Fail to open the testDB")
	}

	key := []byte("1")
	value := []byte("ThisIsLastHash")

	err = levelDB.Put(key, value)
	if err != nil {
		t.Error("Fail to put the data")
	}

	data, err := levelDB.Tip()
	if err != nil {
		t.Error("Fail to get the data")
	}

	if !bytes.Equal(value, data) {
		t.Error("Get invalid data")
	}

	levelDB.Close()
}
