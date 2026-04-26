package logic

import (
	"strconv"
)

func GetSMA(candles [][]string, count int) float64 {
	var total float64
	last := candles[len(candles)-count:]

	for _, c := range last {
		closeStr := c[4]
		closeVal, err := strconv.ParseFloat(closeStr, 64)
		if err != nil {
			return 0
		}

		total += closeVal
	}

	return total / float64(count)
}
