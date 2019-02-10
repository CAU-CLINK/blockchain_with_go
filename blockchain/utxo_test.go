package blockchain_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
	"github.com/stretchr/testify/assert"
)

func TestNewUTXOSet(t *testing.T) {
	utxoSet, err := blockchain.NewUTXOSet()
	assert.NoError(t, err)
	assert.NotNil(t, utxoSet)
}

func TestUTXOSet_FindUTXOList(t *testing.T) {

}

func TestUTXOSet_FindUTXOs(t *testing.T) {

}

func TestUTXOSet_Update(t *testing.T) {

}
