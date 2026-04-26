package logic

func GetTrend60(ema50 float64, ema200 float64) string {
	if ema50 > ema200 {
		return "Up"
	} else if ema50 < ema200 {
		return "Down"
	} else {
		return "Flet"
	}
}
