package logic

import (
	"math"
	"strconv"
)

type OHLC struct {
	Open        float64
	High        float64
	Low         float64
	Close       float64
	Body        float64
	UpperShadow float64
	LowerShadow float64
	Range       float64
	Color       string
}

func RegCandles(candles [][]string) []OHLC {
	var res []OHLC

	for _, c := range candles {
		open_, _ := strconv.ParseFloat(c[1], 64)
		high, _ := strconv.ParseFloat(c[2], 64)
		low, _ := strconv.ParseFloat(c[3], 64)
		close_, _ := strconv.ParseFloat(c[4], 64)

		body := math.Abs(close_ - open_)
		upperShadow := high - math.Max(open_, close_)
		lowerShadow := math.Min(open_, close_) - low
		fullRange := high - low

		var color string
		if close_ > open_ {
			color = "green"
		} else if close_ < open_ {
			color = "red"
		} else {
			color = "doji"
		}

		candle := OHLC{
			Open:        open_,
			High:        high,
			Low:         low,
			Close:       close_,
			Body:        body,
			UpperShadow: upperShadow,
			LowerShadow: lowerShadow,
			Range:       fullRange,
			Color:       color,
		}

		res = append(res, candle)
	}

	return res
}
