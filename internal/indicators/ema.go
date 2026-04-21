package indicators

import "main/internal/signal_logic"

func GetEma(candles []signal_logic.OHLCStruct, count_of_candles int) float64 {

	ema := make([]float64, len(candles))
	sma := 0.0
	for i := 0; i < count_of_candles; i++ {
		sma += candles[i].Close
	}

	ema[count_of_candles-1] = sma / float64(count_of_candles)

	k := 2.0 / float64(count_of_candles+1)

	for i := count_of_candles; i < len(candles); i++ {
		ema[i] = candles[i].Close*k + ema[i-1]*(1-k)
	}

	return ema[len(ema)-1]
}
