package main

import (
	"fmt"
	"os"
)

func main() {

	env := os.Getenv("ETHER_PROVIVER")

	if len(env) == 0 {
		fmt.Println("Please provider an ETHER_PROVIVER.")
		return
	}

	fmt.Println(env)

	bChain := BlockChain{}
	bChain.ConnectServiceProvider(env)
	bChain.GetBlockByNumber(12442098)
	bChain.GetTransactions()

}
