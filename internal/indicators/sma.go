package indicators

import (
	"main/internal/signal_logic"
)

func GetSma(candles []signal_logic.OHLCStruct, count_of_candles int) float64 {
	var sum float64
	var ans float64

	for i := 0; i < len(candles[len(candles)-count_of_candles:]); i++ {
		sum += candles[len(candles)-count_of_candles:][i].Close
	}

	float_len := float64(len(candles[len(candles)-count_of_candles:]))
	ans = sum / float_len

	return ans
}
