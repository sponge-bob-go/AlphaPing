package indicators

import "main/internal/signal_logic"

type Indicators struct {
	SMA     float64 `json:"SMA"`
	EMA     float64 `json:"EMA"`
	Trend15 string  `json:"Trend15"`
	Trend60 string  `json:"Trend60"`
	RSI     float64 `json:"RSI"`
	MACD    struct {
		MACDLine   float64 `json:"MACDLine"`
		SignalLine float64 `json:"SignalLine"`
		Histogram  float64 `json:"Histogram"`
	}
	ATR         float64         `json:"ATR"`
	ATRPercent  float64         `json:"ATRPercent"`
	ADX         float64         `json:"ADX"`
	Volume      string          `json:"Volume"`
	Supports    []float64       `json:"Supports"`
	Resistances []float64       `json:"Resistances"`
	Patterns    map[string]bool `json:"Patterns"`
}

func GetFinalIndicators(candles []signal_logic.OHLCStruct, candles60 []signal_logic.OHLCStruct) Indicators {
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
		chADX <- GetADX(candles, 14)
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

	ema15 := GetEMA(candles, 50)
	result.EMA = ema15
	result.Trend15 = GetTrend15(ema15, sma200)

	ema60 := GetEMA(candles60, 200)
	sma60 := GetSMA(candles60, 50)
	result.Trend60 = GetTrend15(ema60, sma60)

	return result

}
