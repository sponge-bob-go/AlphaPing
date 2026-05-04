package handlers

import (
	"encoding/json"
	"fmt"
	"main/internal/indicators"
	bybit "main/internal/signal_logic/ByBit"
	mexc "main/internal/signal_logic/MEXC"
	"net/http"
)

func HandlerTechData(w http.ResponseWriter, r *http.Request) {
	market := r.URL.Query().Get("market")
	coin_name := r.URL.Query().Get("coin_name")
	interval := r.URL.Query().Get("interval")

	if market == "" {
		w.WriteHeader(403)
	}
	if coin_name == "" {
		w.WriteHeader(403)
	}
	if interval == "" {
		w.WriteHeader(403)
	}

	if market == "mexc" {
		ohlcModel := mexc.CreateMEXCModel()
		ohlc, err := ohlcModel.GetOHLC(coin_name, interval)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		indicator := indicators.GetFinalIndicators(ohlc)
		json.NewEncoder(w).Encode(indicator)
	}
	if market == "bybit" {
		ohlcModel := bybit.CreateBybitModel()
		ohlc, err := ohlcModel.GetOHLC(coin_name, interval)
		fmt.Println(ohlc)
		fmt.Println(ohlc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		indicator := indicators.GetFinalIndicators(ohlc)
		json.NewEncoder(w).Encode(indicator)
	}
}
