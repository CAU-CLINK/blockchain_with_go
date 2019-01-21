package consensus


//좀더 연구하자
//https://medium.com/@mycoralhealth/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f
//https://github.com/topics/proof-of-work?l=go
//https://github.com/bwesterb/go-pow 
//https://github.com/mycoralhealth/blockchain-tutorial/blob/master/proof-work/main.go 

import (
	"blockchain_with_go/blockchain"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var maxNonce = math.MaxInt64

const targetBits = 5

type ProofOfWork struct {
	Header *BlockHeader
	target *big.Int
}

func NewProofOfWork(b *BlockHeader) *ProofOfWork {
	target : big.NewInt(1)
	target.Lsh(target, int(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func IntToHex(n int64) []byte {
    return []byte(strconv.Formatint(n, 16))
}

//내맘대로 수정
func (pow *ProofOfWork) prepareData(Nonce int64) {
	data := bytes.Join(
		[][]bytes{
				pow.Header.PreviousBlockHash,
				pow.Header.MerkelRootHash,
				intToHex(pow.Header.Timestamp)
				intToHex(int64(targetBits))
				intToHex(int64(Nonce))
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	Nonce := 0
	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n", pow.Header)

	for Nonce < maxNonce {
		data := pow.prepareData(Nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
				break
		} else {
				Nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var isValid bool

	data := pow.prepareData(pow.Header.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1 {
		isValid = true
	} else {
		isValid = false
	}

	return isValid

}
