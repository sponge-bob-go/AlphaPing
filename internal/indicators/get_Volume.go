package logic

import (
	"strconv"
)

func GetVolume(candles [][]string, period int) string {
	lastCandles := candles[len(candles)-period:]
	var sum float64

	for _, c := range lastCandles {
		vol, err := strconv.ParseFloat(c[5], 64)
		if err != nil {
			panic(err)
		}
		sum += vol
	}

	avg := sum / float64(period)

	lastCandleVolStr := candles[len(candles)-1][5]
	lastCandleVol, err := strconv.ParseFloat(lastCandleVolStr, 64)
	if err != nil {
		panic(err)
	}

	if lastCandleVol > avg*3 {
		return "Strong Up"
	}
	if lastCandleVol > avg*1.5 {
		return "Volume Up"
	}
	return "Volume Down"
}
