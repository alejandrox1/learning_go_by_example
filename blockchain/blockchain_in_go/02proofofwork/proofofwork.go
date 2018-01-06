package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// Difficulty of mining.
const targetBits = 24

var maxNonce = math.MaxInt64


// ProofOfWork represents a proof-of-work.
type ProofOfWork struct {
	block  *Block
	target *big.Int
}


// NewProofOfWork builds and returns a ProofOfWork.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// Lsh sets target = target << 256-targetBits and returns target.
	target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{b, target}
}

// prepData prepares data to hash it.
func (pow *ProofOfWork) prepData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
}

// Run performs a proof-of-work.
func (pow *ProofOfWork) Run() (int, []byte) {
	// Int representation of the hash.
	var hashInt big.Int
	var hash    [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n")

	return nonce, hash[:]
}


// Validate validates a block's POW.
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
