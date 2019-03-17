package blockchain

import (
	"errors"
	"fmt"

	"github.com/CAU-CLINK/blockchain_with_go/common"
	"github.com/CAU-CLINK/blockchain_with_go/script"
	"github.com/minio/sha256-simd"
)

const subsidy = 10

var ErrAmountExceed = errors.New("Not enough amount")

//Version int32
//LockTime int32
type Transaction struct {
	TxIn  []TxInput
	TxOut []TxOutput
}

// TODO: Implements me with test case
func NewTransaction(pubKey []byte, to string, amount uint, utxos UTXOs) (*Transaction, error) {
	var inputs []TxInput
	var outputs []TxOutput
	var acc uint = 0

	pubKeyHash := common.PubkeyHash(pubKey)

	for key, utxo := range utxos {
		acc += utxo.Value()
		/*
			Need to sign here!
		*/
		input := TxInput{key.TxID(), key.Vout(), script.ScriptSig{nil, pubKeyHash}}
		inputs = append(inputs, input)
	}

	if acc < amount {
		return nil, ErrAmountExceed
	}

	address := common.Base58CheckEncode(pubKeyHash)
	from := fmt.Sprintf("%s", address)

	outputs = append(outputs, NewTxOutput(amount, to))
	if acc > amount {
		outputs = append(outputs, NewTxOutput(acc-amount, from)) // a change
	}

	tx := Transaction{inputs, outputs}

	return &tx, nil
}

// TODO: Implements me with test case
func (tx *Transaction) Hash() ([]byte, error) {
	var hash [32]byte

	serializedTx, err := common.Serialize(&tx)
	if err != nil {
		return nil, err
	}

	hash = sha256.Sum256(serializedTx)

	return hash[:], nil
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
