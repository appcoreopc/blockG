package main

import (
	"context"
	"math/big"
	"testing"

	"github.com/smartystreets/assertions"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

/// Mock
type ClientBlockMock struct {
	mock.Mock
}

func (m *ClientBlockMock) BlockByNumber(context context.Context, blockNumber *big.Int) (*types.Block, error) {
	args := m.Called(context, blockNumber)
	return args.Get(0).(*types.Block), args.Error(1)
}

func WhenValidBlockNumberThenBlockDataReturns(t *testing.T) {

	mock := ClientBlockMock{}
	mock.On("BlockByNumber", context.Background(), 100).Return(nil, nil)

}

func WhenInvBlockNumberThenBlockDataReturns(t *testing.T) {

	mock := ClientBlockMock{}
	mock.On("BlockByNumber", context.Background(), 100).Return(nil, nil)

	bService := NewBlockChainService(&mock)
	bService.GetBlockByNumber(100)
	assertions.ShouldBeNil(bService.client)

}
