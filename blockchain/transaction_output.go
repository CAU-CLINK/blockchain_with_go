package blockchain

import (
	"github.com/CAU-CLINK/blockchain_with_go/script"
	"github.com/btcsuite/btcutil/base58"
)

type TxOutput struct {
	Value        uint
	ScriptPubKey script.ScriptPubKey
}

func NewTxOutput(value uint, address string) TxOutput {
	txOutput := TxOutput{value, script.ScriptPubKey{}}
	txOutput.Lock(address)

	return txOutput
}

func (txOutput *TxOutput) Lock(address string) {
	pubkeyHash := base58.Decode(address)
	pubkeyHash = pubkeyHash[1 : len(pubkeyHash)-4] // 0 : version, last 4 : checksum
	txOutput.ScriptPubKey.PubkeyHash = pubkeyHash
}
