package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
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

// IsCoinbase checks whether the transaction is a coinbase.
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].TXid) == 0 && tx.Vin[0].Vout == -1
}

// SetID sets the ID for a transaction.
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	if err := gob.NewEncoder(&encoded).Encode(tx); err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
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

// NewUTXOTransaction creates a new transaction.
func NewUTXOTransaction(from string, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: Not enough funds.")
	}

	// BUild a list of inputs.
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs.
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		// Change.
		outputs = append(outputs, TXOutput{acc - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}
