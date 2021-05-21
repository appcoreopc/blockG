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

// Create instance of ethClient
func NewBlockServiceProvider(serviceProviderUrl string) *ethclient.Client {

	client, err := ethclient.Dial(serviceProviderUrl)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// create a separate stuct / type implementation for the eth client
type BlockClientImpl interface {
	//Transactions()
	BlockByNumber(context context.Context, blockNumber *big.Int) (*types.Block, error)
}

func NewBlockChainService(ethClient BlockClientImpl) BlockChainService {
	return BlockChainService{
		client: ethClient,
		quit:   make(chan int),
	}
}

////////////////////////////////////////////////////

type BlockChainService struct {
	client       BlockClientImpl
	currentBlock *types.Block
	quit         chan int
}

func (b *BlockChainService) GetBlockByNumber(blockNo int64) {

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

func (b *BlockChainService) GetTransactions() {
	var wg sync.WaitGroup
	wg.Add(1)

	txCh := b.processTransactions(&wg)
	b.storeTransaction(&wg, txCh)
	wg.Wait()
}

func (b *BlockChainService) storeTransaction(wg *sync.WaitGroup, txIn <-chan *types.Transaction) {

	go func() {
		defer wg.Done()
		for tx := range txIn {
			fmt.Println(tx.Hash().Hex())
		}
	}()
}

func (b *BlockChainService) processTransactions(wg *sync.WaitGroup) <-chan *types.Transaction {

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
