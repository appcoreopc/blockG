package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockChain struct {
	client       ethclient.Client
	currentBlock *types.Block
	quit         chan int
}

func (b *BlockChain) ConnectServiceProvider(serviceProviderUrl string) *BlockChain {

	client, err := ethclient.Dial(serviceProviderUrl)
	if err != nil {
		log.Fatal(err)
	}

	b.client = *client
	b.quit = make(chan int)

	return &BlockChain{}
}

func (b *BlockChain) GetBlockByNumber(blockNo int64) {

	blockNumber := big.NewInt(blockNo)
	block, err := b.client.BlockByNumber(context.Background(), blockNumber)
	b.currentBlock = block

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("parent go")
	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Difficulty().Uint64())
	fmt.Println(block.Hash().Hex())
	fmt.Println(len(block.Transactions()))
}

func (b *BlockChain) GetTransactions() {

	var wg sync.WaitGroup

	wg.Add(1)

	txCh := b.processTransactions(&wg)
	b.storeTransaction(&wg, txCh)

	wg.Wait()
}

func (b *BlockChain) storeTransaction(wg *sync.WaitGroup, txIn <-chan *types.Transaction) {

	go func() {
		defer wg.Done()
		for tx := range txIn {
			fmt.Println(tx.Hash().Hex())
		}
	}()
}

func (b *BlockChain) processTransactions(wg *sync.WaitGroup) <-chan *types.Transaction {

	transactionChannel := make(chan *types.Transaction)
	trans := b.currentBlock.Transactions()
	fmt.Println("total trans", len(trans))

	go func() {

		for _, tx := range trans {
			transactionChannel <- tx
		}
		close(transactionChannel)
	}()

	return transactionChannel
}
