package blockchain

import (
	"encoding/hex"
	"log"
	"strconv"

	"github.com/CAU-CLINK/blockchain_with_go/script"
	"github.com/pkg/errors"
)

type UTXOKey string

func NewUTXOKey(txid string, vout int) (UTXOKey, error) {
	if len(txid) != 64 {
		return "", errors.New("Invalid txid")
	}

	var key []byte

	key = append(key, []byte{0x63}...)

	txID, err := hex.DecodeString(txid)
	if err != nil {
		return "", err
	}

	key = append(key, txID...)
	key = append(key, byte(vout))

	strKey := hex.EncodeToString(key)

	return UTXOKey(strKey), nil
}

func (key UTXOKey) TxID() []byte {
	bytes, err := hex.DecodeString(string(key))
	if err != nil {
		log.Panic(err)
	}

	return bytes[1:33]
}

func (key UTXOKey) Vout() int {
	bytes, err := hex.DecodeString(string(key))
	if err != nil {
		log.Panic(err)
	}

	strVout := hex.EncodeToString(bytes[33:])
	vout, err := strconv.Atoi(strVout)

	return vout
}

type UTXOs map[UTXOKey][]UTXO

/*
origianl
key : txid
value : height + value + type + pubkeyhash or pubkey
*/
type UTXO struct {
	height     uint
	value      uint
	typ        string // transaction type
	pubkeyHash []byte // 20 bytes
}

func (utxo UTXO) txOut() TxOutput {
	txOutput := TxOutput{utxo.value, script.ScriptPubKey{utxo.pubkeyHash}}
	return txOutput
}

type UTXOSet struct {
	db Database
}

func NewUTXOSet() (*UTXOSet, error) {
	db, err := NewLevelDB("../db/chainstate/")
	if err != nil {
		return nil, err
	}

	return &UTXOSet{db: db}, nil
}

func (u UTXOSet) FindUTXOList(pubkeyHash []byte) UTXOs {

	return UTXOs{}
}

// TODO : Implements me w/ test case
func (u UTXOSet) FindUTXOs(pubkeyHash []byte, amount int) UTXOs {
	return UTXOs{}
}

// Delete transaction's input (spended)
// Add transaction's output (utxo)
// TODO : Implements me w/ test case
func (u UTXOSet) Update(block *Block) {

}
