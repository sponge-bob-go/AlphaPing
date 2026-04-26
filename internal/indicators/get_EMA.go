package logic

import (
	"strconv"
)

func GetEMA(candles [][]string, count int) float64 {

	var closes []float64
	for _, c := range candles {

		price, err := strconv.ParseFloat(c[4], 64)
		if err != nil {
			return 0
		}
		closes = append(closes, price)
	}

	k := 2.0 / float64(count+1)

	sum := 0.0
	for _, price := range closes[:count] {
		sum += price
	}
	ema := sum / float64(count)

	for _, price := range closes[count:] {
		ema = price*k + ema*(1-k)
	}

	return ema
}

func GetEMAFromCloses(closes []float64, count int) float64 {

	k := 2.0 / float64(count+1)

	sum := 0.0
	for _, price := range closes[:count] {
		sum += price
	}
	ema := sum / float64(count)

	for _, price := range closes[count:] {
		ema = price*k + ema*(1-k)
	}

	return ema
}
