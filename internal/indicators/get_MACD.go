package logic

import (
	"strconv"
)

func GetMACD(candles [][]string) map[string]float64 {

	var closes []float64
	for _, c := range candles {

		price, err := strconv.ParseFloat(c[4], 64)
		if err != nil {
			return nil
		}
		closes = append(closes, price)
	}

	ema12 := GetEMA(candles, 12)

	ema26 := GetEMA(candles, 26)

	macdLine := ema12 - ema26

	macdValues := make([]float64, 0, len(closes)-25)
	for i := 26; i <= len(closes); i++ {
		if i < 26 {
			continue
		}
		ema12I := GetEMAFromCloses(closes[:i], 12)

		ema26I := GetEMAFromCloses(closes[:i], 26)

		macdValues = append(macdValues, ema12I-ema26I)
	}

	signal := GetEMAFromCloses(macdValues[len(macdValues)-9:], 9)

	histogram := macdLine - signal

	result := map[string]float64{
		"macdLine":   macdLine,
		"signalLine": signal,
		"histogram":  histogram,
	}

	return result
}
