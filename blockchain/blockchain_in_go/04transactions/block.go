package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

// Block keeps the Block's headers and a Data section.
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce		  int
}


// NewBlock will create a block and return a reference to it.
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock will create a genesis block.
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}


// Serialize serializes a Block struct.
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	if err := gob.NewEncoder(&result).Encode(b); err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// DeserializeBlock deserializes a gob struct into a Block.
func DeserializeBlock(d []byte) *Block {
	var block Block
	if err := gob.NewDecoder(bytes.NewReader(d)).Decode(&block); err != nil {
		log.Panic(err)
	}
	return &block
}
