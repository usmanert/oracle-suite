//  Copyright (C) 2020 Maker Ecosystem Growth Holdings, INC.
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero General Public License as
//  published by the Free Software Foundation, either version 3 of the
//  License, or (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU Affero General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
