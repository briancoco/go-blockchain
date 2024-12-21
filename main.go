package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"strconv"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")

	//traverse and print blocks
	for _, block := range bc.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
