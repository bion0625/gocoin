package main

import (
	"fmt"

	"github.com/gocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	fmt.Println(chain)
}
