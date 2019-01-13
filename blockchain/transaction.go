package blockchain

import "blockchain_with_go/script"

type Transaction struct {
	TxIn  []TxInput
	TxOut []TxOutput
}

type TxInput struct {
	Txid      []byte
	ScriptSig script.ScriptSig
}

type TxOutput struct {
	Value        uint
	ScriptPubKey script.ScriptPubKey
}
