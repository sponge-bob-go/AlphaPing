package indicators

import "main/internal/signal_logic"

func GetATRPercent(candles []signal_logic.OHLCStruct, period int) float64 {
	atr := GetATR(candles, period)

	lastClose := candles[len(candles)-1].Close

	return (atr / lastClose) * 100
}
