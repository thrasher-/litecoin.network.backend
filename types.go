package main

type BlockInfo struct {
	BlockHeight int64  `json:"block_height"`
	BlockTime   int64  `json:"block_time"`
	BlockHash   string `json:"block_hash"`
	TimeElapsed int64  `json:"time_elapsed"`
	Status      string `json:"status"`
}
