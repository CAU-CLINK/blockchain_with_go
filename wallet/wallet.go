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

// Create New Transaction here
// Get UTXO from UTXOSet and sign by privatekey (scriptSig)
// After validation, send transaction to miner
// TODO: Implements me with test case
func (w Wallet) Send(from string, to string, amount int) {

}
