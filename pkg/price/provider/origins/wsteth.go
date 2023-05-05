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

package origins

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/defiweb/go-eth/abi"
	"github.com/defiweb/go-eth/types"

	"github.com/chronicleprotocol/oracle-suite/pkg/ethereum"
)

//go:embed wsteth_abi.json
var wrappedStakedETHABI []byte

type WrappedStakedETH struct {
	ethClient ethereum.Client //nolint:staticcheck // deprecated ethereum.Client
	addrs     ContractAddresses
	abi       *abi.Contract
	blocks    []int64
}

//nolint:staticcheck // deprecated ethereum.Client
func NewWrappedStakedETH(cli ethereum.Client, addrs ContractAddresses, blocks []int64) (*WrappedStakedETH, error) {
	a, err := abi.ParseJSON(wrappedStakedETHABI)
	if err != nil {
		return nil, err
	}
	return &WrappedStakedETH{
		ethClient: cli,
		addrs:     addrs,
		abi:       a,
		blocks:    blocks,
	}, nil
}

func (s WrappedStakedETH) PullPrices(pairs []Pair) []FetchResult {
	return callSinglePairOrigin(&s, pairs)
}

func (s WrappedStakedETH) callOne(pair Pair) (*Price, error) {
	contract, inverted, err := s.addrs.AddressByPair(pair)
	if err != nil {
		return nil, err
	}

	var callData []byte
	if !inverted {
		callData, err = s.abi.Methods["stEthPerToken"].EncodeArgs()
	} else {
		callData, err = s.abi.Methods["tokensPerStEth"].EncodeArgs()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get contract args for pair: %s", pair.String())
	}

	resp, err := s.ethClient.CallBlocks(context.Background(), types.Call{To: &contract, Input: callData}, s.blocks)
	if err != nil {
		return nil, err
	}

	price, err := reduceEtherAverageFloat(resp)
	if err != nil {
		return nil, err
	}
	priceFloat, _ := price.Float64()
	return &Price{
		Pair:      pair,
		Price:     priceFloat,
		Timestamp: time.Now(),
	}, nil
}
