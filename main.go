package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ea7c5898898f4dcfa59ebe635621fbf8")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	bChain := BlockChain{}
	bChain.ConnectServiceProvider("")
	bChain.GetBlockByNumber(12442098)
	bChain.GetTransactions()

	// account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	// balance, err := client.BalanceAt(context.Background(), account, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Balance:", balance) // 25893180161173005034

}
