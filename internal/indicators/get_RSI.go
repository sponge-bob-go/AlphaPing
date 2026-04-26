package logic

import (
	"math"
	"strconv"
)

func GetRSI(candles [][]string) float64 {
	period := 14

	var closes []float64
	for _, c := range candles {
		if len(c) < 5 {
			return 0
		}
		price, err := strconv.ParseFloat(c[4], 64)
		if err != nil {
			return 0
		}
		closes = append(closes, price)
	}

	var gains []float64
	var losses []float64

	for i := 1; i < len(closes); i++ {
		diff := closes[i] - closes[i-1]
		if diff > 0 {
			gains = append(gains, diff)
			losses = append(losses, 0)
		} else {
			gains = append(gains, 0)
			losses = append(losses, math.Abs(diff))
		}
	}

	sumGains := 0.0
	sumLosses := 0.0

	for _, g := range gains[len(gains)-period:] {
		sumGains += g
	}
	for _, l := range losses[len(losses)-period:] {
		sumLosses += l
	}

	avgGain := sumGains / float64(period)
	avgLoss := sumLosses / float64(period)

	if avgLoss == 0 {
		return 100.0
	}

	rs := avgGain / avgLoss
	rsi := 100 - (100 / (1 + rs))

	return rsi
}
