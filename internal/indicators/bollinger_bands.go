package indicators

import (
	"main/internal/signal_logic"
	"math"
)

func sumOfSquares(candles []signal_logic.OHLCStruct, mean float64) float64 {
	var sum float64
	for _, candle := range candles {
		sum += math.Pow(candle.Close-mean, 2)
	}
	return sum
}

func GetBollingerBands(candles []signal_logic.OHLCStruct) []float64 {

	sma := GetSma(candles, 14)

	sd := math.Sqrt(sumOfSquares(candles, sma) / float64(len(candles)))

	upperBand := sma + 2*sd
	lowerBand := sma - 2*sd
	res := []float64{upperBand, sma, lowerBand}
	return res
}
