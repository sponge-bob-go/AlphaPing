package indicators

import "main/internal/signal_logic"

func GetEMA(candles []signal_logic.OHLCStruct, count int) float64 {

	k := 2.0 / float64(count+1)

	sum := 0.0
	for i := 0; i < count; i++ {
		sum += candles[i].Close
	}

	ema := sum / float64(count)

	for i := count; i < len(candles); i++ {
		ema = candles[i].Close*k + ema*(1-k)
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
