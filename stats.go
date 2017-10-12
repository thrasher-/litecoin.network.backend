package main

import (
	"log"
	"math"
	"net/url"

	"github.com/thrasher-/gocryptotrader/exchanges/bitfinex"
)

const (
	maxCoins            int64 = 84000000
	blockHalvingSubsidy int64 = 840000
	blockStartingReward int64 = 50
	blocksPerDay              = (60 / 2.5) * 24
	daysPerYear               = 365.2425
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

func GetEnergyConsumption() error {
	A4Hashrate := 550         // MH/s
	A4PowerConsumption := 750 // watts
	A4Price := float64(3150)  //USD
	avgPowerPerKW := float64(12)
	totalHashrate, err := GetNetworkHashRate()
	if err != nil {
		return err
	}

	// convert to megahashes per second
	totalHashrate = totalHashrate / 1000 / 1000
	log.Printf("Calculations using most efficient scrypt chip (Innosilicon A4) Stats per unit. Hashrate: %d MH/s, Power consumption: %d Watts, Price per unit: $%.0f USD",
		A4Hashrate, A4PowerConsumption, A4Price)
	log.Printf("Using %.0f cents per kilowatt-hour (average price people in the U.S. pay for electricity is about 12 cents per kilowatt-hour)", avgPowerPerKW)
	networkEquivA4MinerAmount := totalHashrate / float64(A4Hashrate)
	networkPowerConsumption := networkEquivA4MinerAmount * float64(A4PowerConsumption)
	log.Printf("Network hash rate A4 equiv: %.0f", networkEquivA4MinerAmount)
	log.Printf("Mining infrastructure A4 equiv cost $%.2f", networkEquivA4MinerAmount*A4Price)
	log.Printf("Network power consumption: %.2f Watts", networkPowerConsumption)
	networkPowerConsumptionKW := networkPowerConsumption / 1000
	log.Printf("Network power consumption: %.2f KWatts", networkPowerConsumptionKW)
	log.Printf("Network energy cost per kilowatt-hour: $%.2f", networkPowerConsumptionKW*avgPowerPerKW/100)
	log.Printf("Network energy cost per kilowatt-day: $%.2f", (networkPowerConsumptionKW*avgPowerPerKW/100)*24)
	log.Printf("Network energy cost per kilowatt-year: $%.2f", (networkPowerConsumptionKW*avgPowerPerKW/100)*24*daysPerYear)
	return nil
}
