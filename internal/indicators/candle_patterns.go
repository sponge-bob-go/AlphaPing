package indicators

import (
	"main/internal/signal_logic"
	"math"
)

func IsBearishEngulfing(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 2 {
		return false
	}
	return candles[len(candles)-1].Close < candles[len(candles)-1].Open &&
		candles[len(candles)-2].Close > candles[len(candles)-2].Open &&
		candles[len(candles)-1].Close < candles[len(candles)-2].Open &&
		candles[len(candles)-1].Open > candles[len(candles)-2].Close
}

func IsBullishEngulfing(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 2 {
		return false
	}
	return candles[len(candles)-1].Close > candles[len(candles)-1].Open &&
		candles[len(candles)-2].Close < candles[len(candles)-2].Open &&
		candles[len(candles)-1].Close > candles[len(candles)-2].Open &&
		candles[len(candles)-1].Open < candles[len(candles)-2].Close
}

func IsHammer(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 1 {
		return false
	}
	candle := candles[len(candles)-1]
	body := candle.Close - candle.Open
	upperShadow := candle.High - math.Max(candle.Close, candle.Open)
	lowerShadow := math.Min(candle.Close, candle.Open) - candle.Low
	return lowerShadow > 2*body && upperShadow < body
}

func IsInvertedHammer(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 1 {
		return false
	}
	candle := candles[len(candles)-1]
	body := candle.Close - candle.Open
	upperShadow := candle.High - math.Max(candle.Close, candle.Open)
	lowerShadow := math.Min(candle.Close, candle.Open) - candle.Low
	return upperShadow > 2*body && lowerShadow < body
}

func IsDoji(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 1 {
		return false
	}
	candle := candles[len(candles)-1]
	return math.Abs(candle.Close-candle.Open) <= 0.1*(candle.High-candle.Low)
}

func IsMorningStar(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 3 {
		return false
	}
	return candles[len(candles)-3].Close < candles[len(candles)-3].Open &&
		math.Abs(candles[len(candles)-2].Close-candles[len(candles)-2].Open) < 0.1*(candles[len(candles)-2].High-candles[len(candles)-2].Low) &&
		candles[len(candles)-1].Close > candles[len(candles)-2].Close &&
		candles[len(candles)-1].Close > candles[len(candles)-3].Close
}

func IsEveningStar(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 3 {
		return false
	}
	return candles[len(candles)-3].Close > candles[len(candles)-3].Open &&
		math.Abs(candles[len(candles)-2].Close-candles[len(candles)-2].Open) < 0.1*(candles[len(candles)-2].High-candles[len(candles)-2].Low) &&
		candles[len(candles)-1].Close < candles[len(candles)-2].Close &&
		candles[len(candles)-1].Close < candles[len(candles)-3].Close
}

func IsShootingStar(candles []signal_logic.OHLCStruct) bool {
	if len(candles) < 1 {
		return false
	}
	candle := candles[len(candles)-1]
	body := candle.Close - candle.Open
	upperShadow := candle.High - math.Max(candle.Close, candle.Open)
	lowerShadow := math.Min(candle.Close, candle.Open) - candle.Low
	return upperShadow > 2*body && lowerShadow < body
}

func GetCandlePatterns(candles []signal_logic.OHLCStruct) []string {
	var patterns []string

	if IsBullishEngulfing(candles) {
		patterns = append(patterns, "Bullish Engulfing")
	}
	if IsBearishEngulfing(candles) {
		patterns = append(patterns, "Bearish Engulfing")
	}
	if IsHammer(candles) {
		patterns = append(patterns, "Hammer")
	}
	if IsInvertedHammer(candles) {
		patterns = append(patterns, "Inverted Hammer")
	}
	if IsDoji(candles) {
		patterns = append(patterns, "Doji")
	}
	if IsMorningStar(candles) {
		patterns = append(patterns, "Morning Star")
	}
	if IsEveningStar(candles) {
		patterns = append(patterns, "Evening Star")
	}
	if IsShootingStar(candles) {
		patterns = append(patterns, "Shooting Star")
	}

	return patterns
}
