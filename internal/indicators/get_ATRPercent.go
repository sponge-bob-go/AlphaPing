package logic

import (
	"strconv"
)

func GetATRPercent(candles [][]string, period int) float64 {
	atr := GetATR(candles, period)

	lastCloseStr := candles[len(candles)-1][4]
	lastClose, err := strconv.ParseFloat(lastCloseStr, 64)
	if err != nil {
		return 0
	}

	return (atr / lastClose) * 100
}
