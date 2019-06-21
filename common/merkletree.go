package common

import "crypto/sha256"

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

func NewMerkleNode(left *MerkleNode, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		childData := append(left.Data, right.Data...)
		hash := sha256.Sum256(childData)
		node.Data = hash[:]
	}
	node.Left = left
	node.Right = right
	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	if len(data)%2 == 1 {
		data = append(data, data[len(data)-1])
	}
	for _, nodeData := range data {
		node := NewMerkleNode(nil, nil, nodeData)
		nodes = append(nodes, *node)
	}
	for j := 0; j < len(data)/2; j++ {
		var parentNode []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			newNode := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			parentNode = append(parentNode, *newNode)
		}
		nodes = parentNode
	}
	Tree := MerkleTree{&nodes[0]}
	return &Tree
}
