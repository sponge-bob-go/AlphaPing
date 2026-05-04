package indicators

func GetTrend15(ema50, ema200 float64) string {

	diff := (ema50 - ema200) / ema200 * 100

	if diff > 0.2 {
		return "Up"
	}

	if diff < -0.2 {
		return "Down"
	}

	return "Flat"
}
