package ethereumv2

import (
	"context"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereumv2/types"
)

// Client is a lightweight Ethereum RPC.
type Client interface {
	// BlockNumber returns the most recent block number.
	BlockNumber(ctx context.Context) (uint64, error)
	// BlockByNumber returns the block with the given number. If full is true, the
	// returned type is types.BlockTxObjects that contains all transactions in the
	// block. Otherwise, the returned type is types.BlockTxHashes that contains only
	// the transaction hashes.
	BlockByNumber(ctx context.Context, number types.BlockNumber, full bool) (any, error)
	// GetTransactionCount returns the number of transactions sent from the given
	// address.
	GetTransactionCount(ctx context.Context, account types.Address, block types.BlockNumber) (uint64, error)
	// SendRawTransaction sends an encoded transaction to the network.
	SendRawTransaction(ctx context.Context, data types.Bytes) (*types.Hash, error)
	// GetStorageAt returns the value of key in the contract storage at the given
	// address.
	GetStorageAt(ctx context.Context, acc types.Address, key types.Hash, block types.BlockNumber) (*types.Hash, error)
	// FilterLogs returns logs that match the given query.
	FilterLogs(ctx context.Context, q types.FilterLogsQuery) ([]types.Log, error)
}
