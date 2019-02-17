package blockchain_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
	"github.com/stretchr/testify/assert"
)

const testAddress = "16KqkkwaqHd7bBZnadopWYGpfc3aWwFZEV"

// TODO: Implements test case
func TestTransaction_Hash(t *testing.T) {

}

func TestNewCoinbase(t *testing.T) {
	coinbaseTx := blockchain.NewCoinbase(testAddress)
	assert.NotNil(t, coinbaseTx)
	isCoinbase := coinbaseTx.IsCoinbase()
	assert.True(t, isCoinbase, 1)
}

// TODO: Implements test cas
func TestNewTransaction(t *testing.T) {

}
