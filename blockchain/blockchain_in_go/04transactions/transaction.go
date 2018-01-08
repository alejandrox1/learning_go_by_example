package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsidy = 10

// Transaction represents a bitcoin transaction.
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// SetID sets the ID for a transaction.
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var has [32]byte

	if err := gob.NewEncoder(&encoded).Encode(tx); err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}



// TXInput represents a transaction input.
type TXInput struct {
	TXid      []byte
	Vout      int
	ScriptSig string
}



// TXOutput represents a transaction output.
type TXOutput struct {
	Value        int
	ScriptPubKey string
}


// CanUnlockOutputWith checks whether the addresses initiated the transaction.
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// CanBeUnlockedWith checks if the output can be unlocked witht he provided
// data.
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

// NewCoinbaseTX creates a new coinbase transaction.
func NewCoinbaseTX(to string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin},[]TXOutput{txout}}
	tx.SetID()
	return &tx
}
