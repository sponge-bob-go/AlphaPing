package indicators

import "main/internal/signal_logic"

func GetRsi(candles []signal_logic.OHLCStruct) float64 {
	rsList := candles[len(candles)-14:]
	gain := []int{}
	loss := []int{}
	for i := 1; i < len(rsList); i++ {
		diff := candles[i].Clo
	}
}
