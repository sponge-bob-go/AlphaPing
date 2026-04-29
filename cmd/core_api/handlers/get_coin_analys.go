package handlers

import (
	"encoding/json"
	"main/internal/indicators"
	mexc "main/internal/signal_logic/MEXC"
	"net/http"
)

func Handler_analys_coin(w http.ResponseWriter, r *http.Request) {
	market := r.URL.Query().Get("market")
	coin_name := r.URL.Query().Get("coin_name")
	interval := r.URL.Query().Get("interval")

	if market == "mexc" {
		ohlcModel := mexc.CreateMEXCModel()
		ohlc, err := ohlcModel.GetOHLC(coin_name, interval)
		ohlc60, _ := ohlcModel.GetOHLC(coin_name, "60m")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		indicator := indicators.GetFinalIndicators(ohlc, ohlc60)
		json.NewEncoder(w).Encode(indicator)
	}
}
