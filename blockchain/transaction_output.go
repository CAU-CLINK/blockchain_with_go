package blockchain

import "github.com/CAU-CLINK/blockchain_with_go/script"

type TxOutput struct {
	Value        uint
	ScriptPubKey script.ScriptPubKey
}
