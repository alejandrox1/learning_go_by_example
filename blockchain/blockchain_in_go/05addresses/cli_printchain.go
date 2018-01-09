package main

import (
	"fmt"
	"strconv"
)



func (cli *CLI) printChain() {
	bc := NewBlockchain("")
	defer bc.db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
