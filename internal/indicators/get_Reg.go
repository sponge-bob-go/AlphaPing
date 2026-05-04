package indicators

import (
	"main/internal/signal_logic"
	"math"
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

type OHLCStruct signal_logic.OHLCStruct

func RegCandles(candles []signal_logic.OHLCStruct) []OHLC {
	var res []OHLC

	for _, c := range candles {
		open_ := c.Open
		high := c.High
		low := c.Low
		close_ := c.Close

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
