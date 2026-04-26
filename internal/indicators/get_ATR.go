package logic

import (
	"math"
	"strconv"
)

func GetATR(candles [][]string, period int) float64 {

	lastCandles := candles[len(candles)-period:]
	var sumTR float64

	for i := range lastCandles {
		high, err := strconv.ParseFloat(lastCandles[i][2], 64)
		if err != nil {
			return 0
		}

		low, err := strconv.ParseFloat(lastCandles[i][3], 64)
		if err != nil {
			return 0
		}

		if i == 0 {
			tr := high - low
			sumTR += tr
			continue
		}

		prevClose, err := strconv.ParseFloat(lastCandles[i-1][4], 64)
		if err != nil {
			return 0
		}

		rangeHL := high - low
		absHigh := math.Abs(high - prevClose)
		absLow := math.Abs(low - prevClose)

		tr := math.Max(rangeHL, math.Max(absHigh, absLow))
		sumTR += tr
	}

	return sumTR / float64(period)
}
