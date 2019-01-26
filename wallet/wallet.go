package wallet

type Wallet struct {
	PrivateKey *PrivateKey
	PublicKey  *PublicKey
}

func New() (*Wallet, error) {
	privateKey, err := GenerateKey()
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.PubKey()

	return &Wallet{privateKey, publicKey}, nil
}

// TODO: Implements me with test case
func (w Wallet) GetAddress() []byte {
	return nil
}
