package coinconfig

type CoinConfig struct {
	ADX        *float64
	ATR        *float64
	ATRPercent *float64
	EMA        *float64
	MACD       *[]float64
	Patterns   *[]string
}
