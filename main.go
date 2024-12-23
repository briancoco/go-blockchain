package main

import (
	"go-blockchain/blockchain"
	"go-blockchain/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db.Close()

	cli := cli.CLI{bc}
	cli.Run()
}
