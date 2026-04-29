package indicators

import "main/internal/signal_logic"

func GetVolume(candles []signal_logic.OHLCStruct, period int) string {
	lastCandles := candles[len(candles)-period:]
	var sum float64

	for _, c := range lastCandles {
		sum += c.Volume
	}

	avg := sum / float64(period)

	lastCandleVol := candles[len(candles)-1].Volume

	if lastCandleVol > avg*3 {
		return "Strong Up"
	}
	if lastCandleVol > avg*1.5 {
		return "Volume Up"
	}
	return "Volume Down"
}
