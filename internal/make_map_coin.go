package internal

import "sync"

type CoinGeckoData struct {
	CurrentPrice      float64 `json:"CurrentPrice"`
	MarketCap         float64 `json:"MarketCap"`
	MarketCapRank     int     `json:"MarketCapRank"`
	TotalVolume       float64 `json:"TotalVolume"`
	High24h           float64 `json:"High24h"`
	Low24h            float64 `json:"Low24h"`
	CirculatingSupply float64 `json:"CirculatingSupply"`
	TotalSupply       float64 `json:"TotalSupply"`
}

type CoinInfo struct {
	Price    float64       `json:"Price"`
	CoinData CoinGeckoData `json:"CoinData"`
}

type CoinMap struct {
	Mu    sync.RWMutex
	Coins map[string]CoinInfo `json:"Coins"`
}

func MakeCoinMap() *CoinMap {
	return &CoinMap{
		Coins: make(map[string]CoinInfo),
	}
}
