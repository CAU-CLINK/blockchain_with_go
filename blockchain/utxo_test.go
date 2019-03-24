package blockchain_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/wallet"

	"github.com/CAU-CLINK/blockchain_with_go/common"

	"github.com/btcsuite/btcutil/base58"

	"os"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
	"github.com/stretchr/testify/assert"
)

var utxoSetPath = "../db/test/chainstate/"

func testUTXOset(address string) error {
	chainstate, err := blockchain.NewLevelDB(utxoSetPath)
	if err != nil {
		return err
	}

	testDBs := []struct {
		transaction *blockchain.Transaction
		vout        int
	}{
		{
			transaction: blockchain.NewCoinbase(address),
			vout:        0,
		},
		{
			transaction: blockchain.NewCoinbase(address),
			vout:        1,
		},
		{
			transaction: blockchain.NewCoinbase(address),
			vout:        2,
		},
		{
			transaction: blockchain.NewCoinbase(address),
			vout:        3,
		},
		{
			transaction: blockchain.NewCoinbase(address),
			vout:        4,
		},
	}

	for _, db := range testDBs {
		txID, err := db.transaction.Hash()
		if err != nil {
			return err
		}
		utxoKey, err := blockchain.NewUTXOKey(txID, db.vout)
		if err != nil {
			return err
		}

		utxo := blockchain.NewUTXO(db.transaction.TxOut[0], 0)

		serialiezedUTXO, err := common.Serialize(&utxo)
		if err != nil {
			return err
		}

		chainstate.Put(utxoKey.Bytes(), serialiezedUTXO)
	}

	chainstate.Close()

	return nil
}

func TestUTXOKey(t *testing.T) {
	tests := []struct {
		txID         []byte
		vout         int
		expectedKey  string
		expectedID   []byte
		expectedVout int
	}{
		{
			txID: []byte{0xe9, 0xef, 0x55, 0xe2, 0xb2, 0x98, 0xd4, 0x35, 0x4f, 0x6f, 0x4a, 0x0f, 0x8e, 0x79, 0x64, 0x8a, 0x71,
				0xe6, 0x24, 0x01, 0xb6, 0x06, 0xd8, 0x6b, 0x3d, 0xed, 0xbb, 0xdf, 0xd0, 0xb3, 0x52, 0xc9},
			vout:        0,
			expectedKey: "63e9ef55e2b298d4354f6f4a0f8e79648a71e62401b606d86b3dedbbdfd0b352c900",
			expectedID: []byte{0xe9, 0xef, 0x55, 0xe2, 0xb2, 0x98, 0xd4, 0x35, 0x4f, 0x6f, 0x4a, 0x0f, 0x8e, 0x79, 0x64, 0x8a, 0x71,
				0xe6, 0x24, 0x01, 0xb6, 0x06, 0xd8, 0x6b, 0x3d, 0xed, 0xbb, 0xdf, 0xd0, 0xb3, 0x52, 0xc9},
			expectedVout: 0,
		},
		{
			txID: []byte{0x6d, 0xed, 0x89, 0xcb, 0x82, 0x59, 0xe6, 0xd9, 0x6e, 0x6c, 0x8f, 0x1e, 0xe6, 0xae, 0xf5, 0x70, 0x17,
				0x1b, 0xcc, 0xe3, 0x28, 0xfc, 0x1d, 0xd4, 0x98, 0x72, 0x54, 0x55, 0x96, 0x58, 0x28, 0xbc},
			vout:        1,
			expectedKey: "636ded89cb8259e6d96e6c8f1ee6aef570171bcce328fc1dd498725455965828bc01",
			expectedID: []byte{0x6d, 0xed, 0x89, 0xcb, 0x82, 0x59, 0xe6, 0xd9, 0x6e, 0x6c, 0x8f, 0x1e, 0xe6, 0xae, 0xf5, 0x70, 0x17,
				0x1b, 0xcc, 0xe3, 0x28, 0xfc, 0x1d, 0xd4, 0x98, 0x72, 0x54, 0x55, 0x96, 0x58, 0x28, 0xbc},
			expectedVout: 1,
		},
		{
			txID: []byte{0x44, 0x5f, 0x3a, 0x75, 0x37, 0x73, 0xa7, 0x29, 0x74, 0x39, 0x77, 0xbe, 0xd1, 0x34, 0xda, 0x6c, 0x58,
				0x90, 0xe3, 0x34, 0x7a, 0xc0, 0x5d, 0x67, 0xf5, 0xbe, 0x2c, 0xf9, 0x6f, 0x32, 0xd0, 0x1c},
			vout:        2,
			expectedKey: "63445f3a753773a729743977bed134da6c5890e3347ac05d67f5be2cf96f32d01c02",
			expectedID: []byte{0x44, 0x5f, 0x3a, 0x75, 0x37, 0x73, 0xa7, 0x29, 0x74, 0x39, 0x77, 0xbe, 0xd1, 0x34, 0xda, 0x6c, 0x58,
				0x90, 0xe3, 0x34, 0x7a, 0xc0, 0x5d, 0x67, 0xf5, 0xbe, 0x2c, 0xf9, 0x6f, 0x32, 0xd0, 0x1c},
			expectedVout: 2,
		},
	}

	for _, test := range tests {
		key, err := blockchain.NewUTXOKey(test.txID, test.vout)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedKey, string(key))
		assert.Equal(t, test.expectedID, key.TxID())
		assert.Equal(t, test.expectedVout, key.Vout())
	}

}

func TestNewUTXOSet(t *testing.T) {
	utxoSetPath := "../db/chainstate"

	utxoSet, err := blockchain.NewUTXOSet(utxoSetPath)
	defer os.RemoveAll(utxoSetPath)

	assert.NoError(t, err)
	assert.NotNil(t, utxoSet)
}

func TestUTXOSet_FindUTXOList(t *testing.T) {

}

func TestUTXOSet_FindUTXOs(t *testing.T) {
	err := testUTXOset(testAddress)
	if err != nil {
		t.Error(err)
	}

	defer os.RemoveAll(utxoSetPath)

	utxoSet, err := blockchain.NewUTXOSet(utxoSetPath)

	pubkeyHash := base58.Decode(testAddress)
	pubkeyHash = pubkeyHash[1 : len(pubkeyHash)-4]

	utxos, err := utxoSet.FindUTXOs(pubkeyHash, 50)
	if err != nil {
		t.Error(err)
	}

	if len(utxos) != 5 {
		t.Error("Invalid utxo!")
	}
}

// Scenario
// wallet1 has utxo 10, 10, 10, 10, 10
// wallet1 sends 10 utxo to testAddress
// tx output is 10 value to testAddress and 40 value to wallet1 (change)
// utxoset has 2 utxos : wallet1 has 40 utxo, testAddress has 10 utxo
func TestUTXOSet_Update(t *testing.T) {
	wallet1, err := wallet.New()
	if err != nil {
		t.Error(err)
	}

	err = testUTXOset(wallet1.Address())
	if err != nil {
		t.Error(err)
	}

	defer os.RemoveAll(utxoSetPath)

	utxoSet, err := blockchain.NewUTXOSet(utxoSetPath)
	if err != nil {
		t.Error(err)
	}

	utxos, err := utxoSet.FindUTXOs(wallet1.PubKeyHash(), 50)
	if err != nil {
		t.Error(err)
	}

	tx, _ := blockchain.NewTransaction(wallet1.PublicKey.Bytes(), testAddress, 10, utxos)

	var txs []*blockchain.Transaction
	txs = append(txs, tx)

	block := blockchain.NewBlock(nil, txs)

	utxoSet.Update(block)

	testAddrPubkeyHash := base58.Decode(testAddress)
	testAddrPubkeyHash = testAddrPubkeyHash[1 : len(testAddrPubkeyHash)-4]

	wallet1UTXOs, err := utxoSet.FindUTXOs(wallet1.PubKeyHash(), 40)
	testAddrUTXOs, err := utxoSet.FindUTXOs(testAddrPubkeyHash, 10)

	if len(wallet1UTXOs) != 1 {
		t.Error("Invalid utxos!")
	}

	if len(testAddrUTXOs) != 1 {
		t.Error("Invalid utxos!")
	}
}
