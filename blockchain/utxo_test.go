package blockchain_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/blockchain"
	"github.com/stretchr/testify/assert"
)

func TestUTXOKey(t *testing.T) {
	tests := []struct {
		txID         string
		vout         int
		expectedKey  string
		expectedID   []byte
		expectedVout int
	}{
		{
			txID:        "e9ef55e2b298d4354f6f4a0f8e79648a71e62401b606d86b3dedbbdfd0b352c9",
			vout:        0,
			expectedKey: "63e9ef55e2b298d4354f6f4a0f8e79648a71e62401b606d86b3dedbbdfd0b352c900",
			expectedID: []byte{0xe9, 0xef, 0x55, 0xe2, 0xb2, 0x98, 0xd4, 0x35, 0x4f, 0x6f, 0x4a, 0x0f, 0x8e, 0x79, 0x64, 0x8a, 0x71,
				0xe6, 0x24, 0x01, 0xb6, 0x06, 0xd8, 0x6b, 0x3d, 0xed, 0xbb, 0xdf, 0xd0, 0xb3, 0x52, 0xc9},
			expectedVout: 0,
		},
		{
			txID:        "6ded89cb8259e6d96e6c8f1ee6aef570171bcce328fc1dd498725455965828bc",
			vout:        1,
			expectedKey: "636ded89cb8259e6d96e6c8f1ee6aef570171bcce328fc1dd498725455965828bc01",
			expectedID: []byte{0x6d, 0xed, 0x89, 0xcb, 0x82, 0x59, 0xe6, 0xd9, 0x6e, 0x6c, 0x8f, 0x1e, 0xe6, 0xae, 0xf5, 0x70, 0x17,
				0x1b, 0xcc, 0xe3, 0x28, 0xfc, 0x1d, 0xd4, 0x98, 0x72, 0x54, 0x55, 0x96, 0x58, 0x28, 0xbc},
			expectedVout: 1,
		},
		{
			txID:        "445f3a753773a729743977bed134da6c5890e3347ac05d67f5be2cf96f32d01c",
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
	utxoSet, err := blockchain.NewUTXOSet()
	assert.NoError(t, err)
	assert.NotNil(t, utxoSet)
}

func TestUTXOSet_FindUTXOList(t *testing.T) {

}

func TestUTXOSet_FindUTXOs(t *testing.T) {

}

func TestUTXOSet_Update(t *testing.T) {

}
