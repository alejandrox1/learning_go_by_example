package main

// The blockchain keeps a sequence of blocks.
type Blockchain struct {
	// Ordered, back-linked list.
	blocks []*Block
}

// AddBlock will save the provided data into a Block.
// Blocks are stored in order and each block is linked tothe previous one.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain will create a new Blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
