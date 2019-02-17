package blockchain_test

import (
	"encoding/hex"
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
	"github.com/CAU-CLINK/blockchain_with_go/script"
	"github.com/stretchr/testify/assert"
)

func TestNewTxOutput(t *testing.T) {
	txOutput := blockchain.NewTxOutput(0, testAddress)
	assert.Equal(t, uint(0), txOutput.Value)
	assert.Len(t, txOutput.ScriptPubKey.PubkeyHash, 20)

	pubkeyHashStr := bytesToHexString(txOutput.ScriptPubKey.PubkeyHash)
	assert.Equal(t, "3a68d07bd0f5359b1b67f1154f8e5c51bf0c7a13", pubkeyHashStr)
}

func TestTxOutput_Lock(t *testing.T) {
	txOutput := blockchain.TxOutput{0, script.ScriptPubKey{}}
	txOutput.Lock(testAddress)

	assert.Len(t, txOutput.ScriptPubKey.PubkeyHash, 20)
	pubkeyHashStr := bytesToHexString(txOutput.ScriptPubKey.PubkeyHash)
	assert.Equal(t, "3a68d07bd0f5359b1b67f1154f8e5c51bf0c7a13", pubkeyHashStr)
}

func bytesToHexString(bytes []byte) string {
	str := hex.EncodeToString(bytes)
	return str
}
