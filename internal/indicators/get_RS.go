package indicators

import "main/internal/signal_logic"

func GetRS(candles []signal_logic.OHLCStruct) ([]float64, []float64) {

	var supports []float64
	var resistances []float64

	for i := 1; i < len(candles)-1; i++ {

		floatHigh := candles[i].High
		floatLow := candles[i].Low
		floatPrevHigh := candles[i-1].High
		floatNextHigh := candles[i+1].High

		floatPrevLow := candles[i-1].Low

		floatNextLow := candles[i+1].Low
		if floatHigh > floatPrevHigh && floatHigh > floatNextHigh {
			resistances = append(resistances, floatHigh)
		}

		if floatLow < floatPrevLow && floatLow < floatNextLow {
			supports = append(supports, floatLow)
		}
	}

	return supports, resistances
}
