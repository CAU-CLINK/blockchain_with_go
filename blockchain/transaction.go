package blockchain

import "github.com/CAU-CLINK/blockchain_with_go/script"

const subsidy = 10

//Version int32
//LockTime int32
type Transaction struct {
	TxIn  []TxInput
	TxOut []TxOutput
}

// TODO: Implements me with test case
func NewTransaction() *Transaction {
	return nil
}

// TODO: Implements me with test case
func (tx *Transaction) Hash() []byte {
	return nil
}

func NewCoinbase(to string) *Transaction {
	txin := TxInput{nil, -1, script.ScriptSig{}}
	txout := NewTxOutput(subsidy, to)
	tx := Transaction{[]TxInput{txin}, []TxOutput{txout}}

	return &tx
}

func (tx Transaction) IsCoinbase() bool {
	return len(tx.TxIn) == 1 && tx.TxIn[0].Vout == -1
}
