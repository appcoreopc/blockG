package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockChain struct {
	client ethclient.Client
}

func New(serviceProviderUrl string) *ethclient.Client {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ea7c5898898f4dcfa59ebe635621fbf8")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (b *BlockChain) GetBlockByNumber(blockNo int64) {

	blockNumber := big.NewInt(blockNo)
	block, err := b.client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144
}
