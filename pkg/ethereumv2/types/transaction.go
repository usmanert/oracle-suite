package types

// Transaction represents a transaction.
type Transaction struct {
	Hash             Hash    `json:"hash"`
	BlockHash        Hash    `json:"blockHash"`
	BlockNumber      Number  `json:"blockNumber"`
	TransactionIndex Number  `json:"transactionIndex"`
	From             Address `json:"from"`
	To               Address `json:"to"`
	Gas              Number  `json:"gas"`
	GasPrice         Number  `json:"gasPrice"`
	Input            Bytes   `json:"input"`
	Nonce            Number  `json:"nonce"`
	Value            Number  `json:"value"`
	V                Number  `json:"v"`
	R                Number  `json:"r"`
	S                Number  `json:"s"`
}

// TransactionReceiptType represents transaction receipt.
type TransactionReceiptType struct {
	TransactionHash   Hash     `json:"transactionHash"`
	TransactionIndex  Number   `json:"transactionIndex"`
	BlockHash         Hash     `json:"blockHash"`
	BlockNumber       Number   `json:"blockNumber"`
	From              Address  `json:"from"`
	To                Address  `json:"to"`
	CumulativeGasUsed Number   `json:"cumulativeGasUsed"`
	GasUsed           Number   `json:"gasUsed"`
	ContractAddress   *Address `json:"contractAddress"`
	Logs              []Log    `json:"logs"`
	LogsBloom         Bytes    `json:"logsBloom"`
	Root              *Hash    `json:"root"`
	Status            *Number  `json:"status"`
}
