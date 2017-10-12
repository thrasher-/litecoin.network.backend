package main

import (
	"log"
)

func main() {
	bi, err := TestBlockHeight()
	if err != nil {
		log.Fatal(err)
	}

	err = GetEnergyConsumption()
	if err != nil {
		log.Fatal(err)
	}

	difficulty, err := GetDifficulty()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Network Difficulty: %v", difficulty)

	hashrate, err := GetNetworkHashRate()
	if err != nil {
		log.Fatal(err)
	}
	hashrate = hashrate / 1000 / 1000 / 1000 / 1000
	log.Printf("Network Hashrate: %f TH/s", hashrate)
	height := bi.BlockHeight
	log.Printf("Height: %d\n", height)
	log.Printf("Block reward halvings: %d\n", GetHalvings(height))
	log.Printf("Remaining blocks until halving: %d\n", GetRemainingBlocks(height))
	reward := GetRewardPerBlock(height)
	log.Printf("Current block reward: %v\n", reward)
	totalCoins := GetTotalCoins(height)
	log.Printf("Total coins: %d", totalCoins)
	ifrate := GetInflationRate(float64(totalCoins), reward)
	ifrate = ifrate * 100 / 1
	log.Printf("Current inflation rate: %.2f%%", ifrate)
	log.Printf("Remaining coins to mine: %d", GetRemainingCoins(int64(totalCoins)))

	price, err := GetBitfinexLastPrice()
	if err != nil {
		log.Fatal(err)
	}
	marketCap, _ := GetMarketCap(int64(totalCoins))
	log.Printf("Market cap: $%.2f USD", marketCap)
	log.Printf("Block reward value: $%.2f USD", reward*price)
}
