package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereumv2/types"
)

type Client struct {
	mock.Mock
}

func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	args := c.Called(ctx)
	return args.Get(0).(uint64), args.Error(1)
}

func (c *Client) BlockByNumber(ctx context.Context, number types.BlockNumber, full bool) (any, error) {
	args := c.Called(ctx, number, full)
	return args.Get(0), args.Error(1)
}

func (c *Client) GetTransactionCount(ctx context.Context, acc types.Address, block types.BlockNumber) (uint64, error) {
	args := c.Called(ctx, acc, block)
	return args.Get(0).(uint64), args.Error(1)
}

func (c *Client) SendRawTransaction(ctx context.Context, data types.Bytes) (*types.Hash, error) {
	args := c.Called(ctx, data)
	return args.Get(0).(*types.Hash), args.Error(1)
}

func (c *Client) GetStorageAt(ctx context.Context, acc types.Address, key types.Hash, block types.BlockNumber) (*types.Hash, error) {
	args := c.Called(ctx, acc, key, block)
	return args.Get(0).(*types.Hash), args.Error(1)
}

func (c *Client) FilterLogs(ctx context.Context, q types.FilterLogsQuery) ([]types.Log, error) {
	args := c.Called(ctx, q)
	return args.Get(0).([]types.Log), args.Error(1)
}
