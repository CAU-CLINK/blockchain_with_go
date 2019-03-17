package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func GenerateKey() (*PrivateKey, error) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{private}, nil
}

type PrivateKey struct {
	*ecdsa.PrivateKey
}

func (privateKey *PrivateKey) PubKey() *PublicKey {
	return &PublicKey{&privateKey.PublicKey}
}

type PublicKey struct {
	*ecdsa.PublicKey
}

func (pubkey PublicKey) Bytes() []byte {
	return append(pubkey.X.Bytes(), pubkey.Y.Bytes()...)
}
