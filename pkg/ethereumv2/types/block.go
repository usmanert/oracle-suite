package types

// Block represents an Ethereum block.
type Block struct {
	Number           Number  `json:"number"`
	Hash             Hash    `json:"hash"`
	ParentHash       Hash    `json:"parentHash"`
	Nonce            Nonce   `json:"nonce"`
	Sha3Uncles       Hash    `json:"sha3Uncles"`
	LogsBloom        Bloom   `json:"logsBloom"`
	TransactionsRoot Hash    `json:"transactionsRoot"`
	StateRoot        Hash    `json:"stateRoot"`
	ReceiptsRoot     Hash    `json:"receiptsRoot"`
	Miner            Address `json:"miner"`
	MixHash          Hash    `json:"mixHash"`
	Difficulty       Number  `json:"difficulty"`
	TotalDifficulty  Number  `json:"totalDifficulty"`
	ExtraData        Bytes   `json:"extraData"`
	Size             Number  `json:"size"`
	GasLimit         Number  `json:"gasLimit"`
	GasUsed          Number  `json:"gasUsed"`
	Timestamp        Number  `json:"timestamp"`
	Uncles           []Hash  `json:"uncles"`
}

// BlockTxHashes represents Ethereum block with transaction hashes.
type BlockTxHashes struct {
	Block
	Transactions []Hash `json:"transactions"`
}

// BlockTxObjects represents Ethereum block with full transactions.
type BlockTxObjects struct {
	Block
	Transactions []Transaction `json:"transactions"`
}
