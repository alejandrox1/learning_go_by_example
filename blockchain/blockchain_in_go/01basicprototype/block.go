package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps the Block's headers and a Data section.
type Block struct {
	// Metadata (Headers).
	Timestamp     int64 // Time of block creation.
	PrevBlockHash []byte
	Hash          []byte

	// Body.
	Data []byte
}

// SetHash calculates and sets the block's hash.
// Take the block fields, concatenate them, calculate the SHA256 on the
// concatenated combination.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock will create a block and return a reference to it.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
	}
	block.SetHash()
	return block
}

// NewGenesisBlock will create a genesis block.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
