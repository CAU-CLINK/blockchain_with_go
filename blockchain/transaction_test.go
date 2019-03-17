package blockchain_test

import (
	"os"
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/wallet"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
	"github.com/stretchr/testify/assert"
)

const testAddress = "16KqkkwaqHd7bBZnadopWYGpfc3aWwFZEV"

func testWalletDatas() ([]byte, blockchain.UTXOs) {
	wal, _ := wallet.New()
	pubKey := wal.PublicKey.Bytes()

	_ = testUTXOset(wal.Address())

	utxoSet, _ := blockchain.NewUTXOSet(utxoSetPath)
	defer os.RemoveAll(utxoSetPath)

	utxos, _ := utxoSet.FindUTXOs(wal.PubKeyHash(), 50)

	return pubKey, utxos
}

// TODO: Implements test case
func TestTransaction_Hash(t *testing.T) {

}

func TestNewCoinbase(t *testing.T) {
	coinbaseTx := blockchain.NewCoinbase(testAddress)
	assert.NotNil(t, coinbaseTx)
	isCoinbase := coinbaseTx.IsCoinbase()
	assert.True(t, isCoinbase, 1)
}

// transaction which send to testAddress
func TestNewTransaction_normal(t *testing.T) {
	pubKey, utxos := testWalletDatas()
	transaction, err := blockchain.NewTransaction(pubKey, testAddress, 50, utxos)
	if err != nil {
		t.Error(err)
	}

	if len(transaction.TxOut) != 1 {
		t.Errorf("Invalid txOutput counts! - expected : %d, got : %d", 1, len(transaction.TxOut))
	}
}

// transaction which send to testAddress and has exceed amount
func TestNewTransaction_exceed(t *testing.T) {
	pubKey, utxos := testWalletDatas()
	_, err := blockchain.NewTransaction(pubKey, testAddress, 60, utxos)
	if err != blockchain.ErrAmountExceed {
		t.Error("Invalid transaction result")
	}
}

// transaction which send to testAddress and has change
func TestNewTransaction_hasChange(t *testing.T) {
	pubKey, utxos := testWalletDatas()
	transaction, err := blockchain.NewTransaction(pubKey, testAddress, 40, utxos)
	if err != nil {
		t.Error(err)
	}

	if len(transaction.TxOut) != 2 {
		t.Errorf("Invalid txOutput counts! - expected : %d, got : %d", 2, len(transaction.TxOut))
	}
}
