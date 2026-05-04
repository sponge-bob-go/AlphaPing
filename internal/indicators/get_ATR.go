package indicators

import (
	"main/internal/signal_logic"
	"math"
)

func GetATR(candles []signal_logic.OHLCStruct, period int) float64 {

	lastCandles := candles[len(candles)-period:]
	var sumTR float64

	for i := range lastCandles {
		high := lastCandles[i].High
		low := lastCandles[i].Low

		if i == 0 {
			tr := high - low
			sumTR += tr
			continue
		}

		prevClose := lastCandles[i-1].Close

		rangeHL := high - low
		absHigh := math.Abs(high - prevClose)
		absLow := math.Abs(low - prevClose)

		tr := math.Max(rangeHL, math.Max(absHigh, absLow))
		sumTR += tr
	}

	return sumTR / float64(period)
}
