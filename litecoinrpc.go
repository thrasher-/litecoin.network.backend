package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func BuildLitecoinServerURL() string {
	return fmt.Sprintf("http://user:pass@localhost:9332")
}

func SendRPCRequest(method, req interface{}) (map[string]interface{}, error) {
	var params []interface{}
	if req != nil {
		params = append(params, req)
	} else {
		params = nil
	}

	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     1,
		"params": params,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(BuildLitecoinServerURL(), "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result["error"] != nil {
		errorMsg := result["error"].(map[string]interface{})
		return nil, fmt.Errorf("Error code: %v, message: %v\n", errorMsg["code"], errorMsg["message"])
	}
	return result, nil
}

func GetBlockHeight() (int64, error) {
	result, err := SendRPCRequest("getinfo", nil)
	if err != nil {
		return 0, err
	}
	result = result["result"].(map[string]interface{})
	block := result["blocks"].(float64)
	return int64(block), nil
}

func GetBlockHash(block int64) (string, error) {
	result, err := SendRPCRequest("getblockhash", block)
	if err != nil {
		return "", err
	}
	return result["result"].(string), nil
}

func GetBlockTime(block string) (int64, error) {
	result, err := SendRPCRequest("getblock", block)
	if err != nil {
		return 0, err
	}

	result = result["result"].(map[string]interface{})
	blockTime := result["time"].(float64)
	return int64(blockTime), nil
}

func GetDifficulty() (float64, error) {
	result, err := SendRPCRequest("getdifficulty", nil)
	if err != nil {
		return 0, err
	}

	difficulty := result["result"]
	return difficulty.(float64), nil
}

func GetNetworkHashRate() (float64, error) {
	result, err := SendRPCRequest("getnetworkhashps", nil)
	if err != nil {
		return 0, err
	}

	hashrate := result["result"]
	return hashrate.(float64), nil
}

func TestBlockHeight() (BlockInfo, error) {
	var blockInfo BlockInfo
	blockHeight, err := GetBlockHeight()
	if err != nil {
		return blockInfo, err
	}

	blockHash, err := GetBlockHash(blockHeight)
	if err != nil {
		return blockInfo, err
	}

	blockTime, err := GetBlockTime(blockHash)
	if err != nil {
		return blockInfo, err
	}

	blockInfo.BlockHeight = blockHeight
	blockInfo.BlockHash = blockHash
	blockInfo.BlockTime = blockTime
	blockInfo.TimeElapsed = GetSecondsElapsed(blockTime)
	blockInfo.Status = TimeSinceLastBlock(blockTime)
	return blockInfo, nil
}

func TimeSinceLastBlock(blockTime int64) string {
	seconds := GetSecondsElapsed(blockTime)
	if seconds >= int64(60*2.5) && seconds < 60*10 {
		return "Block not found within 2.5 minutes."
	}
	if seconds >= 60*10 && seconds < 60*30 {
		return "Block not found within 10 minutes."
	}
	if seconds > 60*30 {
		return "POTENTIAL ISSUE: Block not found within 30 minutes."
	}
	return "OK"
}

func BlockMonitor() {
	bi, err := TestBlockHeight()
	if err != nil {
		log.Fatal(err)
	}

	errCounter := 0
	blockHeight := bi.BlockHeight

	for {
		bInfo, err := TestBlockHeight()
		if err != nil {
			errCounter++
			log.Println(err)

			if errCounter > 5 {
				log.Fatal(err)
			}
		} else {
			errCounter = 0
			if bInfo.BlockHeight != blockHeight {
				fmt.Printf("New block! Height: %d Hash: %s Time: %d - %s\n", bInfo.BlockHeight, bInfo.BlockHash, bInfo.BlockTime, TimeSinceLastBlock(bInfo.BlockTime))
			}
		}
		time.Sleep(time.Second * 10)
	}
}
