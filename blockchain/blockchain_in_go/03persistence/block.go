package main

import (
	"bytes"
	"encoding/gob"
	"log"
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


// Serialize serializes a Block struct.
func (b *Block) Serialize() []byte {
	var results bytes.Buffer
	if err := gob.NewEncoder(&results).Encode(b); err != nil {
		log.Panic(err)
	}
	return results.Bytes()
}

// DeserializeBlock deserializes a gob struct into a Block.
func DeserializeBlock(d []byte) *Block {
	var block Block
	if err := gob.NewDecoder(bytes.NewReader(d)).Decode(&block); err != nil {
		log.Panic(err)
	}
	return &block
}
