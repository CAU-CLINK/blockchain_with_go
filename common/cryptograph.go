package common

import (
	"log"

	"github.com/btcsuite/btcutil/base58"
	"github.com/minio/sha256-simd"
	"golang.org/x/crypto/ripemd160"
)

const version = 0x00
const checksumLength = 4

func Base58CheckEncode(payload []byte) string {
	versionedPayload := append([]byte{version}, payload...)
	checksum := Checksum(versionedPayload)

	fullyPayload := append(versionedPayload, checksum...)
	encodedPayload := base58.Encode(fullyPayload)

	return encodedPayload
}

func PubkeyHash(pubKey []byte) []byte {
	pubkeySHA256 := sha256.Sum256(pubKey)

	RIPEMD160 := ripemd160.New()
	_, err := RIPEMD160.Write(pubkeySHA256[:])
	if err != nil {
		log.Panic(err)
	}
	pubkeyRIPEMD160 := RIPEMD160.Sum(nil)

	return pubkeyRIPEMD160
}

func Checksum(hash160 []byte) []byte {
	hashed := doubleHash(hash160)
	return hashed[:checksumLength]
}

func doubleHash(payload []byte) []byte {
	fisrt := sha256.Sum256(payload)
	second := sha256.Sum256(fisrt[:])
	return second[:]
}
