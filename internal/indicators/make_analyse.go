package indicators

import (
	"main/internal/signal_logic"
	"sync"
)

type OHLCAnalyse struct {
	Atr            []float64
	BollingerBands []float64
	CandlePatterns []string
	Ema            float64
	MACD           float64
	Rsi            float64
	Sma            float64
	Score          int
	mu             sync.Mutex
	wg             sync.WaitGroup
}

func calculateScore(analysis OHLCAnalyse) float64 {
	score := 0.0

	if len(analysis.Atr) > 0 && analysis.Atr[len(analysis.Atr)-1] > 0 {
		score += 1
	}

	if analysis.BollingerBands > 1 {
		score += 1
	} else if analysis.BollingerBands < -1 {
		score -= 1
	}

	for _, pattern := range analysis.CandlePatterns {
		switch pattern {
		case "Bullish Engulfing":
			score += 2
		case "Bearish Engulfing":
			score -= 2
		case "Morning Star":
			score += 1
		case "Evening Star":
			score -= 1
		case "Hammer":
			score += 1
		case "Inverted Hammer":
			score += 1
		case "Doji":
			score += 0
		case "Shooting Star":
			score -= 1
		}
	}

	if analysis.Ema > analysis.Sma {
		score += 1
	} else {
		score -= 1
	}

	if analysis.MACD > 0 {
		score += 2
	} else {
		score -= 2
	}

	if analysis.Rsi > 70 {
		score -= 1
	} else if analysis.Rsi < 30 {
		score += 1
	}

	if analysis.Sma > analysis.BollingerBands {
		score += 1
	} else {
		score -= 1
	}
	return score
}

func MakeAnalyse(data *OHLCAnalyse, candles []signal_logic.OHLCStruct) {
	data.mu.Lock()
	data.wg.Add(8)
	go func() {
		defer data.wg.Done()
		data.Atr = GetATR(candles, 14)
	}()
	go func() {
		defer data.wg.Done()
		data.BollingerBands = GetBollingerBands(candles)
	}()
}
