package blockchain

type UTXOs map[string]UTXO

type UTXO struct {
	height     int
	value      int
	typ        string // transaction type
	pubkeyHash []byte // 20 bytes
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
