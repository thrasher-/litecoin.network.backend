package main

type BlockInfo struct {
	BlockHeight int64  `json:"block_height"`
	BlockTime   int64  `json:"block_time"`
	BlockHash   string `json:"block_hash"`
	TimeElapsed int64  `json:"time_elapsed"`
	Status      string `json:"status"`
}

type MiningCalculations struct {
	DeviceName                       string  `json:"device_name"`
	DeviceHashrate                   int64   `json:"device_hashrate_mhashpsec"`
	DevicePowerConsumption           int64   `json:"device_power_consumption_watts"`
	DevicePriceUSD                   float64 `json:"device_price_usd"`
	DeviceNetworkHashrateEquivilant  float64 `json:"device_network_hashrate_equivilant"`
	MiningInfrastructureCost         float64 `json:"network_mining_infrastructure_cost"`
	MiningInfrastructureCost51Attack float64 `json:"network_mining_infrastructure_cost_51percent_attack"`
	NetworkPowerConsumptionWatts     float64 `json:"network_power_consumption_watts"`
	NetworkPowerConsumptionKWatts    float64 `json:"network_power_consumption_kwatts"`
	NetworkPowerCostKilowattHour     float64 `json:"network_power_cost_kilowatthr"`
	NetworkPowerCostKilowattDay      float64 `json:"network_power_cost_kilowattday"`
	NetworkPowerCostKilowattYear     float64 `json:"network_power_cost_kilowattyear"`
	AvgCostPerKilowattCents          float64 `json:"average_cost_per_kilowatt_cents"`
}

type Output struct {
	MiningInfo             MiningCalculations `json:"mining_calculations"`
	NetworkDifficulty      float64            `json:"network_difficulty"`
	NetworkHashrate        float64            `json:"network_hashrate"`
	NetworkHashrateTH      float64            `json:"network_hashrate_thashpsec"`
	RemainingBlocksHalving int64              `json:"remaining_blocks_till_halving"`
	InflationRate          float64            `json:"inflation_rate"`
	CoinsTotal             int64              `json:"coins_total"`
	CoinsRemaining         int64              `json:"coins_remaining"`
	MarketCap              float64            `json:"market_cap"`
	BlockInformation       BlockInfo          `json:"block_info"`
	BlockRewardHalvings    int64              `json:"block_reward_halvings"`
	BlockRewardCoins       float64            `json:"block_reward"`
	BlockRewardUSD         float64            `json:"block_reward_usd"`
}
