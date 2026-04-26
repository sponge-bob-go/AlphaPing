package indicators

type Indicators struct {
	SMA     float64
	EMA     float64
	Trend15 string
	Trend60 string
	RSI     float64
	MACD    struct {
		MACDLine   float64
		SignalLine float64
		Histogram  float64
	}
	ATR         float64
	ATRPercent  float64
	ADX         float64
	Volume      string
	Supports    []float64
	Resistances []float64
	Patterns    map[string]bool
}

func GetFinalIndicators(candles [][]string) Indicators {
	result := Indicators{}
	result.Patterns = make(map[string]bool)

	chSMA50 := make(chan float64)
	chSMA200 := make(chan float64)
	chEMA := make(chan float64)
	chRSI := make(chan float64)
	chMACD := make(chan map[string]float64)
	chATR := make(chan float64)
	chATRPercent := make(chan float64)
	chADX := make(chan float64)
	chVolume := make(chan string)
	chRS := make(chan []float64)
	chResistances := make(chan []float64)
	chPatterns := make(chan map[string]bool)

	go func() {
		chSMA50 <- GetSMA(candles, 50)
	}()

	go func() {
		chSMA200 <- GetSMA(candles, 200)
	}()

	go func() {
		chEMA <- GetEMA(candles, 200)
	}()

	go func() {
		chRSI <- GetRSI(candles)
	}()

	go func() {
		chMACD <- GetMACD(candles)
	}()

	go func() {
		chATR <- GetATR(candles, 14)
	}()

	go func() {
		chATRPercent <- GetATRPercent(candles, 14)
	}()

	go func() {
		ohlc := RegCandles(candles)
		chADX <- GetADX(ohlc, 14)
	}()

	go func() {
		chVolume <- GetVolume(candles, 20)
	}()

	go func() {
		supports, resistances := GetRS(candles)
		chRS <- supports
		chResistances <- resistances
	}()

	go func() {
		chPatterns <- GetPatterns(candles)
	}()

	result.SMA = <-chSMA50
	sma200 := <-chSMA200
	result.EMA = <-chEMA
	result.RSI = <-chRSI

	macd := <-chMACD
	result.MACD.MACDLine = macd["macdLine"]
	result.MACD.SignalLine = macd["signalLine"]
	result.MACD.Histogram = macd["histogram"]

	result.ATR = <-chATR
	result.ATRPercent = <-chATRPercent
	result.ADX = <-chADX
	result.Volume = <-chVolume
	result.Supports = <-chRS
	result.Resistances = <-chResistances
	result.Patterns = <-chPatterns

	ema50 := result.EMA
	ema200 := sma200
	result.Trend15 = GetTrend15(ema50, ema200)
	result.Trend60 = GetTrend60(ema50, ema200)

	return result
}
