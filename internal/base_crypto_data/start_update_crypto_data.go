package basecryptodata

import (
	internal "main/internal"
	"time"

	"github.com/k0kubun/pp/v3"
)

var coinConfig = map[string]struct {
	GeckoID     string
	BybitSymbol string
}{
	"BTCUSDT": {"bitcoin", "BTCUSDT"},
	"ETHUSDT": {"ethereum", "ETHUSDT"},
	"SOLUSDT": {"solana", "SOLUSDT"},
	"TONUSDT": {"the-open-network", "TONUSDT"},
}

func StartUpdateCoinsPrice(m *internal.CoinMap) error {
	for {
		for _, cfg := range coinConfig {

			err := UpdateCryptoPrice(m, cfg.BybitSymbol)
			if err != nil {
				return err
			}

			time.Sleep(3 * time.Second)
		}

		pp.Print(m.Coins)
		time.Sleep(15 * time.Second)
	}
}

func StartUpdateCoinsInfo(m *internal.CoinMap) error {
	for {
		err := UpdateCoinGeckoData(m)
		if err != nil {
			return err
		}

		pp.Print(m.Coins)
		time.Sleep(4 * time.Minute)
	}
}
