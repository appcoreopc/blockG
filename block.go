package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockChain struct {
	client       ethclient.Client
	currentBlock *types.Block
}

func (b *BlockChain) ConnectServiceProvider(serviceProviderUrl string) *BlockChain {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ea7c5898898f4dcfa59ebe635621fbf8")
	if err != nil {
		log.Fatal(err)
	}

	b.client = *client
	return &BlockChain{}
}

func (b *BlockChain) GetBlockByNumber(blockNo int64) {

	blockNumber := big.NewInt(blockNo)
	block, err := b.client.BlockByNumber(context.Background(), blockNumber)
	b.currentBlock = block

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Difficulty().Uint64())
	fmt.Println(block.Hash().Hex())
	fmt.Println(len(block.Transactions()))
}

func (b *BlockChain) GetTransactions() {

	for _, tx := range b.currentBlock.Transactions() {
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())
	}
}
