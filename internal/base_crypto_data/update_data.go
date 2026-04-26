package basecryptodata

import (
	"encoding/json"
	"fmt"
	"io"
	internal "main/internal"
	"net/http"
	"strconv"
)

type BybitResponse struct {
	Result struct {
		List []struct {
			Symbol    string `json:"symbol"`
			LastPrice string `json:"lastPrice"`
		} `json:"list"`
	} `json:"result"`
}

type CoinGeckoResponse struct {
	ID                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	CurrentPrice      float64 `json:"current_price"`
	MarketCap         float64 `json:"market_cap"`
	MarketCapRank     int     `json:"market_cap_rank"`
	TotalVolume       float64 `json:"total_volume"`
	High24h           float64 `json:"high_24h"`
	Low24h            float64 `json:"low_24h"`
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
}

func UpdateCryptoPrice(m *internal.CoinMap, coinName string) error {
	url := fmt.Sprintf(
		"https://api.bybit.com/v5/market/tickers?category=spot&symbol=%s",
		coinName,
	)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Print(string(body))
	var result BybitResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	if len(result.Result.List) == 0 {
		return fmt.Errorf("empty response")
	}

	price, err := strconv.ParseFloat(result.Result.List[0].LastPrice, 64)
	if err != nil {
		return err
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()

	coin := m.Coins[coinName]
	coin.Price = price
	m.Coins[coinName] = coin

	return nil
}

func UpdateCoinGeckoData(m *internal.CoinMap) error {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin,ethereum,solana,the-open-network"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result []CoinGeckoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()

	for key, cfg := range coinConfig {

		for _, c := range result {
			if c.ID == cfg.GeckoID {

				coin := m.Coins[key]

				coin.CoinData = internal.CoinGeckoData{
					CurrentPrice:      c.CurrentPrice,
					MarketCap:         c.MarketCap,
					MarketCapRank:     c.MarketCapRank,
					TotalVolume:       c.TotalVolume,
					High24h:           c.High24h,
					Low24h:            c.Low24h,
					CirculatingSupply: c.CirculatingSupply,
					TotalSupply:       c.TotalSupply,
				}

				m.Coins[key] = coin
			}
		}
	}

	return nil
}
