package logic

import (
	"strconv"
)

func GetRS(candles [][]string) ([]float64, []float64) {

	var supports []float64
	var resistances []float64

	for i := 1; i < len(candles)-1; i++ {

		floatHigh, err := strconv.ParseFloat(candles[i][2], 64)
		if err != nil {
			return nil, nil
		}

		floatLow, err := strconv.ParseFloat(candles[i][3], 64)
		floatPrevHigh, err := strconv.ParseFloat(candles[i-1][2], 64)
		floatNextHigh, err := strconv.ParseFloat(candles[i+1][2], 64)

		floatPrevLow, err := strconv.ParseFloat(candles[i-1][3], 64)

		floatNextLow, err := strconv.ParseFloat(candles[i+1][3], 64)
		if floatHigh > floatPrevHigh && floatHigh > floatNextHigh {
			resistances = append(resistances, floatHigh)
		}

		if floatLow < floatPrevLow && floatLow < floatNextLow {
			supports = append(supports, floatLow)
		}
	}

	return supports, resistances
}
