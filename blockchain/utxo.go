package blockchain

import (
	"encoding/gob"
	"encoding/hex"
	"log"
	"strconv"

	"github.com/CAU-CLINK/blockchain_with_go/common"

	"bytes"

	"github.com/pkg/errors"
)

type UTXOKey string

func NewUTXOKey(txID []byte, vout int) (UTXOKey, error) {
	if len(txID) != 32 {
		return "", errors.New("Invalid txid")
	}

	var key []byte
	key = append(key, []byte{0x63}...)
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

func (key UTXOKey) Bytes() []byte {
	bytes, err := hex.DecodeString(string(key))
	if err != nil {
		log.Panic(err)
	}
	return bytes
}

type UTXOs map[UTXOKey]*UTXO

/*
origianl
key : txid
value : height + value + type + pubkeyhash or pubkey
*/
type UTXO struct {
	TxOutput TxOutput
	Height   uint
	Typ      string // transaction type
}

// TODO: Implements me w/ test case. it must contain height ( can get in coinbase)
func NewUTXO(output TxOutput, height uint) UTXO {
	return UTXO{output, height, ""}
}

// TODO: Implements me w/ test case
func (utxo UTXO) Bytes() []byte {
	return nil
}

func (utxo UTXO) Value() uint {
	return utxo.TxOutput.Value
}

func (utxo UTXO) PubkeyHash() []byte {
	return utxo.TxOutput.ScriptPubKey.PubkeyHash
}

func (utxo UTXO) txOut() TxOutput {
	txOutput := TxOutput{utxo.Value(), utxo.TxOutput.ScriptPubKey}
	return txOutput
}

func (tx *UTXO) Serialize() ([]byte, error) {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		return nil, err
	}

	return encoded.Bytes(), nil
}

func DeserializeUtxo(data []byte) (*UTXO, error) {
	var utxo UTXO

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&utxo)
	if err != nil {
		return nil, err
	}

	return &utxo, nil
}

type UTXOSet struct {
	db Database
}

func NewUTXOSet(chainstatePath string) (*UTXOSet, error) {
	db, err := NewLevelDB(chainstatePath)
	if err != nil {
		return nil, err
	}

	return &UTXOSet{db: db}, nil
}

// TODO : Implements me w/ test case
func (u UTXOSet) FindUTXOList(pubkeyHash []byte) UTXOs {
	return UTXOs{}
}

// TODO : Implements me w/ test case
func (u *UTXOSet) FindUTXOs(pubkeyHash []byte, amount uint) (UTXOs, error) {
	var utxos UTXOs = make(map[UTXOKey]*UTXO)
	var acc uint = 0

	db := u.db

	iter := db.Iterator()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		utxoKey := hex.EncodeToString(key)

		var utxo *UTXO = &UTXO{}

		err := common.Deserialize(value, utxo)
		if err != nil {
			return nil, err
		}

		if bytes.Equal(utxo.PubkeyHash(), pubkeyHash) {
			acc += utxo.Value()
			utxos[UTXOKey(utxoKey)] = utxo
		}

		if acc >= amount {
			return utxos, nil
		}
	}

	return utxos, errors.New("Not enought utxo!")
}

// Delete transaction's input (spended)
// Add transaction's output (utxo)
// TODO : Refactoring me w/ test case
func (u *UTXOSet) Update(block *Block) {
	db := u.db

	for _, tx := range block.Transactions {
		if !tx.IsCoinbase() {
			for _, vin := range tx.TxIn {
				deletedKey, err := NewUTXOKey(vin.Txid, vin.Vout)
				if err != nil {
					log.Println(err)
					continue
				}
				db.Delete(deletedKey.Bytes())
			}
		}

		for vout, out := range tx.TxOut {
			txID, err := tx.Hash()
			if err != nil {
				log.Println(err)
				continue
			}
			updatedKey, err := NewUTXOKey(txID, vout)

			utxo := NewUTXO(out, 0)
			serializedUTXO, err := common.Serialize(&utxo)
			if err != nil {
				log.Println(err)
				continue
			}
			db.Put(updatedKey.Bytes(), serializedUTXO)
		}
	}
}
