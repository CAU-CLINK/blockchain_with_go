package wallet_test

import (
	"testing"

	"github.com/CAU-CLINK/blockchain_with_go/wallet"
	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {
	wallet, err := wallet.New()
	assert.NoError(t, err)
	assert.NotNil(t, wallet)
}

func TestWallet_GetAddress(t *testing.T) {

}
