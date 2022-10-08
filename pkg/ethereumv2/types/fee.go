package types

// FeeHistory represents the result of the feeHistory RPC call.
type FeeHistory struct {
	OldestBlock   Number     `json:"oldestBlock"`
	Reward        [][]Number `json:"reward"`
	BaseFeePerGas []Number   `json:"baseFeePerGas"`
	GasUsedRatio  []float64  `json:"gasUsedRatio"`
}
