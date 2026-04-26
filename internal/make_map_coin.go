package internal

import "sync"

type CoinGeckoData struct {
	CurrentPrice      float64
	MarketCap         float64
	MarketCapRank     int
	TotalVolume       float64
	High24h           float64
	Low24h            float64
	CirculatingSupply float64
	TotalSupply       float64
}

type CoinInfo struct {
	Price    float64
	CoinData CoinGeckoData
}

type CoinMap struct {
	Mu    sync.RWMutex
	Coins map[string]CoinInfo
}

func MakeCoinMap() *CoinMap {
	return &CoinMap{
		Coins: make(map[string]CoinInfo),
	}
}
