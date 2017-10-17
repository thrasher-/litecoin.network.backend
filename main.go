package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.MarshalIndent(GenerateOutput(), "", "\t")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})

	log.Println("Starting HTTP server on port http://localhost:8444")
	log.Fatal(http.ListenAndServe(":8444", nil))
}

func GenerateOutput() Output {
	var op Output
	bi, err := TestBlockHeight()
	if err != nil {
		log.Fatal(err)
	}
	op.BlockInformation = bi

	mc, err := GetEnergyConsumption()
	if err != nil {
		log.Fatal(err)
	}
	op.MiningInfo = mc

	difficulty, err := GetDifficulty()
	if err != nil {
		log.Fatal(err)
	}
	op.NetworkDifficulty = difficulty

	hashrate, err := GetNetworkHashRate()
	if err != nil {
		log.Fatal(err)
	}
	op.NetworkHashrate = hashrate
	op.NetworkHashrateTH = hashrate / 1000 / 1000 / 1000 / 1000
	op.BlockRewardHalvings = GetHalvings(op.BlockInformation.BlockHeight)
	op.RemainingBlocksHalving = GetRemainingBlocks(op.BlockInformation.BlockHeight)
	reward := GetRewardPerBlock(op.BlockInformation.BlockHeight)
	op.BlockRewardCoins = reward
	op.CoinsTotal = GetTotalCoins(op.BlockInformation.BlockHeight)
	ifrate := GetInflationRate(float64(op.CoinsTotal), reward)
	ifrate = ifrate * 100 / 1
	op.InflationRate = ifrate
	op.CoinsRemaining = GetRemainingCoins(int64(op.CoinsTotal))
	price, err := GetBitfinexLastPrice()
	if err != nil {
		log.Fatal(err)
	}
	op.MarketCap = float64(op.CoinsTotal) * price
	op.BlockRewardUSD = reward * price
	return op
}
