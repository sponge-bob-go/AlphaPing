package indicators

import (
	"main/internal/signal_logic"
	"math"
)

func getTr(candleCurr signal_logic.OHLCStruct, candlePrev signal_logic.OHLCStruct) float64 {
	High_low_diff := candleCurr.High - candleCurr.Low
	Abs_high_close := math.Abs(candleCurr.High - candlePrev.Close)
	Abs_low_close := math.Abs(candleCurr.Low - candlePrev.Close)

	AbsMax := math.Max(Abs_high_close, Abs_low_close)
	return math.Max(AbsMax, High_low_diff)
}

func GetATR(candles []signal_logic.OHLCStruct, period int) []float64 {
	n := len(candles)
	if n <= period {
		return nil
	}

	tr := make([]float64, n)
	atr := make([]float64, n)

	for i := 0; i < n; i++ {
		tr[i] = getTr(candles[i], candles[i-1])
	}

	sum := 0.0
	for i := 1; i < n; i++ {
		sum += tr[i]
	}

	atr[period] = sum / float64(period)

	for i := period + 1; i < n; i++ {
		atr[i] = (atr[i-1]*float64(period-1) + tr[i]) / float64(period)
	}

	return atr
}
