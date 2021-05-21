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

	client := NewBlockServiceProvider(env)

	bChain := NewBlockChainService(client)

	//bChain.SetupServiceProvider(env)
	bChain.GetBlockByNumber(12442098)
	//bChain.GetTransactions()

}
