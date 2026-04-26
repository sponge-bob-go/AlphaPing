package logic

func GetPatterns(candles [][]string) map[string]bool {
	regsCandles := RegCandles(candles)
	patternsCandles := make(map[string]bool)

	if len(regsCandles) < 3 {
		return patternsCandles
	}

	n := len(regsCandles)
	last := regsCandles[n-1]
	prev := regsCandles[n-2]
	third := regsCandles[n-3]

	patternsCandles["white_solders"] =
		last.Color == "green" &&
			prev.Color == "green" &&
			third.Color == "green" &&
			third.Close < prev.Close &&
			prev.Close < last.Close

	patternsCandles["black_solders"] =
		last.Color == "red" &&
			prev.Color == "red" &&
			third.Color == "red" &&
			third.Close > prev.Close &&
			prev.Close > last.Close

	patternsCandles["doji"] =
		last.Body <= last.Range*0.1

	patternsCandles["dragonfly_doji"] =
		patternsCandles["doji"] &&
			last.LowerShadow > last.Range*0.6

	patternsCandles["gravestone_doji"] =
		patternsCandles["doji"] &&
			last.UpperShadow > last.Range*0.6

	patternsCandles["hammer"] =
		last.Color == "green" &&
			last.LowerShadow > last.Body*2 &&
			last.UpperShadow < last.Body*0.3

	patternsCandles["builish_engulfing"] =
		prev.Color == "red" &&
			last.Color == "green" &&
			last.Open < prev.Close &&
			last.Close > prev.Open

	patternsCandles["bearish_engulfing"] =
		prev.Color == "green" &&
			last.Color == "red" &&
			last.Open > prev.Close &&
			last.Close < prev.Open

	return patternsCandles
}
