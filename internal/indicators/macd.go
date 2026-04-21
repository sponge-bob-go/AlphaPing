package indicators

import "main/internal/signal_logic"

func GetMACD(candles []signal_logic.OHLCStruct) float64 {
	ema12 := GetEma(candles, 12)
	ema26 := GetEma(candles, 26)

	return ema12 - ema26
}
