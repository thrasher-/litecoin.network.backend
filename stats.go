package main

import (
	"math"
	"net/url"

	"github.com/thrasher-/gocryptotrader/exchanges/bitfinex"
)

const (
	maxCoins            int64 = 84000000
	blockHalvingSubsidy int64 = 840000
	blockStartingReward int64 = 50
	blocksPerDay              = (60 / 2.5) * 24
)

func GetHalvings(blocks int64) int64 {
	return blocks / blockHalvingSubsidy
}

func GetRemainingBlocks(blocks int64) int64 {
	halvings := GetHalvings(blocks)
	if halvings == 0 {
		return blockHalvingSubsidy - blocks
	}
	halvings++
	return halvings*blockHalvingSubsidy - blocks
}

func GetRewardPerBlock(blocks int64) float64 {
	halvings := GetHalvings(blocks)
	if halvings == 0 {
		return float64(blockStartingReward)
	}

	blockReward := float64(blockStartingReward)
	for i := int64(0); i < halvings; i++ {
		blockReward = blockReward / 2
	}
	return blockReward
}

func GetTotalCoins(blocks int64) int64 {
	halvings := GetHalvings(blocks)
	if halvings == 0 {
		return blocks * blockStartingReward
	}

	coins := float64(0)
	blockReward := float64(blockStartingReward)

	for i := int64(0); i < halvings; i++ {
		coins += blockReward * float64(blockHalvingSubsidy)
		blocks -= blockHalvingSubsidy
		blockReward = blockReward / 2
	}
	coins += blockReward * float64(blocks)
	return int64(coins)
}

func GetInflationRate(totalCoins, blockReward float64) float64 {
	return math.Pow(float64(((totalCoins+blockReward)/totalCoins)), (365*blocksPerDay)) - 1
}

func GetRemainingCoins(coinCount int64) int64 {
	return maxCoins - coinCount
}

func GetBitfinexLastPrice() (float64, error) {
	bfx := bitfinex.Bitfinex{}
	result, err := bfx.GetTicker("LTCUSD", url.Values{})
	if err != nil {
		return 0, err
	}
	return result.Last, nil
}

func GetMarketCap(coinCount int64) (float64, error) {
	price, err := GetBitfinexLastPrice()
	if err != nil {
		return 0, err
	}
	return float64(coinCount) * price, nil
}

func GetHashrateDistribution() {
}
