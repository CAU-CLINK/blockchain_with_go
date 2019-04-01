package common

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMerkleNode(t *testing.T) {
	data := [][]byte{
		[]byte("block"),
		[]byte("c-link"),
		[]byte("jang"),
	}

	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	n7 := NewMerkleNode(n5, n6, nil)
	assert.Equal(
		t,
		"363392912e1025f5f98ce0a11b671d5ffbdb2d61b727a853974d60e265a486ad",
		hex.EncodeToString(n5.Data),
		"Level 1 hash 1 is correct",
	)
	assert.Equal(
		t,
		"6978e0721e17413cc4ee204f340893710540b3b93c0d0da380f983b11ac5e420",
		hex.EncodeToString(n6.Data),
		"Level 1 hash 2 is correct",
	)
	assert.Equal(
		t,
		"b854fb4f828c57df3eac8f7835847b2c00b814ecff10fb5fcc26e3973431fb50",
		hex.EncodeToString(n7.Data),
		"Root hash is correct",
	)
}
func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("block"),
		[]byte("c-link"),
		[]byte("jang"),
	}

	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	n7 := NewMerkleNode(n5, n6, nil)

	rootNode := hex.EncodeToString(n7.Data)
	mTree := NewMerkleTree(data)

	assert.Equal(
		t,
		rootNode,
		hex.EncodeToString(mTree.RootNode.Data),
		"Merkle tree root hash is correct")
}
