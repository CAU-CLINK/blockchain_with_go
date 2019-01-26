package wallet_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/wallet"
	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	privateKey, err := wallet.GenerateKey()
	assert.NoError(t, err)
	assert.NotNil(t, privateKey)
}

func TestPrivateKey_PubKey(t *testing.T) {
	privateKey, err := wallet.GenerateKey()
	assert.NoError(t, err)
	assert.NotNil(t, privateKey)

	publickKey := privateKey.Public()
	assert.NotNil(t, publickKey)
}
