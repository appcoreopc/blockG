package main

import (
	"fmt"
	"os"
)

func main() {

	env := os.Getenv("ETHER_PROVIVER")
	fmt.Println(env)

	bChain := BlockChain{}
	bChain.ConnectServiceProvider(env)
	bChain.GetBlockByNumber(12442098)
	bChain.GetTransactions()

}
