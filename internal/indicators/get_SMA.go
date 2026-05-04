package indicators

import "main/internal/signal_logic"

func GetSMA(candles []signal_logic.OHLCStruct, count int) float64 {
	var total float64
	last := candles[len(candles)-count:]

	for _, c := range last {
		total += c.Close
	}

	return total / float64(count)
}
