package main

import (
	"time"
)

// Block keeps the Block's headers and a Data section.
type Block struct {
	// Metadata (Headers).
	Timestamp     int64 // Time of block creation.
	PrevBlockHash []byte
	Hash          []byte
	Nonce		  int

	// Body.
	Data []byte
}


// NewBlock will create a block and return a reference to it.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
		Data:          []byte(data),
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock will create a genesis block.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
