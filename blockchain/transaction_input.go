package blockchain

import "github.com/CAU-CLINK/blockchain_with_go/script"

type TxInput struct {
	Txid      []byte
	Vout      int
	ScriptSig script.ScriptSig
}
