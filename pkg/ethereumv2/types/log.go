package types

// Log represents a contract log event.
type Log struct {
	Address     Address `json:"address"`
	Topics      []Hash  `json:"topics"`
	Data        Bytes   `json:"data"`
	BlockHash   Hash    `json:"blockHash"`
	BlockNumber Number  `json:"blockNumber"`
	TxHash      Hash    `json:"transactionHash"`
	TxIndex     Number  `json:"transactionIndex"`
	LogIndex    Number  `json:"logIndex"`
	Removed     bool    `json:"removed"`
}

// FilterLogsQuery represents a query to filter logs.
type FilterLogsQuery struct {
	Address   Addresses    `json:"address"`
	FromBlock *BlockNumber `json:"fromBlock,omitempty"`
	ToBlock   *BlockNumber `json:"toBlock,omitempty"`
	Topics    []Hashes     `json:"topics"`
	BlockHash *Hash        `json:"blockhash,omitempty"`
}
