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

package rpcclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereumv2/types"
	"github.com/chronicleprotocol/oracle-suite/pkg/util/errutil"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type testableClient struct {
	*Client
	req *http.Request
	res *http.Response
}

func newTestableClient() *testableClient {
	t := &testableClient{}
	t.Client = New(errutil.Must(rpc.DialHTTPWithClient("http://localhost/", &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) (*http.Response, error) {
			t.req = req
			return t.res, nil
		}),
	})))
	return t
}

func TestClient_BlockNumber(t *testing.T) {
	t.Parallel()
	cli := newTestableClient()
	cli.res = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))),
	}
	blockNumber, err := cli.BlockNumber(context.Background())

	require.NoError(t, err)
	assert.Equal(t, uint64(1), blockNumber)
	assert.Equal(t, http.MethodPost, cli.req.Method)
	assert.Equal(t, "/", cli.req.URL.Path)
	assert.Equal(t, "application/json", cli.req.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"jsonrpc":"2.0","method":"eth_blockNumber","id":1}`, readAll(t, cli.req.Body))
}

func TestClient_GetTransactionCount(t *testing.T) {
	t.Parallel()
	cli := newTestableClient()
	cli.res = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))),
	}
	nonce, err := cli.GetTransactionCount(
		context.Background(),
		types.HexToAddress("0x00112233445566778899aabbccddeeff00112233"),
		types.StringToBlockNumber("latest"),
	)

	require.NoError(t, err)
	assert.Equal(t, uint64(1), nonce)
	assert.Equal(t, http.MethodPost, cli.req.Method)
	assert.Equal(t, "/", cli.req.URL.Path)
	assert.Equal(t, "application/json", cli.req.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["0x00112233445566778899aabbccddeeff00112233","latest"],"id":1}`, readAll(t, cli.req.Body))
}

func TestClient_SendRawTransaction(t *testing.T) {
	t.Parallel()
	cli := newTestableClient()
	cli.res = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))),
	}
	txHash, err := cli.SendRawTransaction(
		context.Background(),
		types.HexToBytes("0x00112233445566778899aabbccddeeff00112233"),
	)

	require.NoError(t, err)
	assert.Equal(t, types.HexToHash("0x1"), *txHash)
	assert.Equal(t, http.MethodPost, cli.req.Method)
	assert.Equal(t, "/", cli.req.URL.Path)
	assert.Equal(t, "application/json", cli.req.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0x00112233445566778899aabbccddeeff00112233"],"id":1}`, readAll(t, cli.req.Body))
}

func TestClient_GetStorageAt(t *testing.T) {
	t.Parallel()
	cli := newTestableClient()
	cli.res = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))),
	}
	storage, err := cli.GetStorageAt(
		context.Background(),
		types.HexToAddress("0x00112233445566778899aabbccddeeff00112233"),
		types.HexToHash("0x00000000000000000000000000112233445566778899aabbccddeeff00112233"),
		types.StringToBlockNumber("latest"),
	)

	require.NoError(t, err)
	assert.Equal(t, types.HexToHash("0x1"), *storage)
	assert.Equal(t, http.MethodPost, cli.req.Method)
	assert.Equal(t, "/", cli.req.URL.Path)
	assert.Equal(t, "application/json", cli.req.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"jsonrpc":"2.0","method":"eth_getStorageAt","params":["0x00112233445566778899aabbccddeeff00112233","0x00000000000000000000000000112233445566778899aabbccddeeff00112233","latest"],"id":1}`, readAll(t, cli.req.Body))
}

func TestClient_FilterLogs(t *testing.T) {
	t.Parallel()
	cli := newTestableClient()
	cli.res = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(`{"jsonrpc":"2.0","id":1,"result":[{"address":"0x00112233445566778899aabbccddeeff00112233","topics":["0x00112233445566778899aabbccddeeff00112233"],"data":"0x00112233445566778899aabbccddeeff00112233","blockNumber":"0x1","transactionHash":"0x00112233445566778899aabbccddeeff00112233","transactionIndex":"0x1","blockHash":"0x00112233445566778899aabbccddeeff00112233","logIndex":"0x1","removed":false}]}`))),
	}
	logs, err := cli.FilterLogs(
		context.Background(),
		types.FilterLogsQuery{
			FromBlock: ptr(types.Uint64ToBlockNumber(10000)),
			ToBlock:   ptr(types.StringToBlockNumber("latest")),
			Address: types.Addresses{
				types.HexToAddress("0x00112233445566778899aabbccddeeff00112233"),
			},
			Topics: []types.Hashes{
				{
					types.HexToHash("0x00112233445566778899aabbccddeeff00112233"),
				},
			},
		},
	)

	require.NoError(t, err)
	assert.Equal(t, []types.Log{
		{
			Address:     types.HexToAddress("0x00112233445566778899aabbccddeeff00112233"),
			Topics:      []types.Hash{types.HexToHash("0x00112233445566778899aabbccddeeff00112233")},
			Data:        types.HexToBytes("0x00112233445566778899aabbccddeeff00112233"),
			BlockNumber: types.Uint64ToNumber(1),
			TxHash:      types.HexToHash("0x00112233445566778899aabbccddeeff00112233"),
			TxIndex:     types.Uint64ToNumber(1),
			BlockHash:   types.HexToHash("0x00112233445566778899aabbccddeeff00112233"),
			LogIndex:    types.Uint64ToNumber(1),
			Removed:     false,
		},
	}, logs)
	assert.Equal(t, http.MethodPost, cli.req.Method)
	assert.Equal(t, "/", cli.req.URL.Path)
	assert.Equal(t, "application/json", cli.req.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"jsonrpc":"2.0","method":"eth_getLogs","params":[{"fromBlock":"0x2710","toBlock":"latest","address":["0x00112233445566778899aabbccddeeff00112233"],"topics":[["0x00000000000000000000000000112233445566778899aabbccddeeff00112233"]]}],"id":1}`, readAll(t, cli.req.Body))
}

func readAll(t *testing.T, r io.Reader) string {
	buf := bytes.NewBuffer(nil)
	_, err := buf.ReadFrom(r)
	require.NoError(t, err)
	return buf.String()
}

func ptr[T any](v T) *T {
	return &v
}
