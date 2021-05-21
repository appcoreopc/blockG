package main

import (
	"context"
	"math/big"
	"testing"

	"github.com/smartystreets/assertions"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

/// Settng up Mock
type ClientBlockMock struct {
	mock.Mock
}

func (m *ClientBlockMock) BlockByNumber(context context.Context, blockNumber *big.Int) (*types.Block, error) {
	args := m.Called(context, blockNumber)
	return args.Get(0).(*types.Block), args.Error(1)
}

////////////////////

func TestWhenInvalidBlockNumberThenBlockDataReturns(t *testing.T) {

	blockChainData := &types.Block{}
	blockNumber := big.NewInt(100)

	mock := ClientBlockMock{}

	mock.On("BlockByNumber", context.Background(), blockNumber).Return(blockChainData, nil)

	bService := NewBlockChainService(&mock)
	bService.GetBlockByNumber(100)
	assertions.ShouldBeNil(bService.client)

}

func TestWhenValidBlockNumberThenBlockDataReturns(t *testing.T) {

	blockChainData := &types.Block{}
	blockNumber := big.NewInt(100)

	mock := ClientBlockMock{}
	mock.On("BlockByNumber", context.Background(), blockNumber).Return(blockChainData, nil)

	bService := NewBlockChainService(&mock)
	bService.GetBlockByNumber(100)
	assertions.ShouldNotBeNil(bService.client)

}
