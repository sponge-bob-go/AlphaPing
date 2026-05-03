package handlers

import (
	"encoding/json"
	"main/internal/indicators"
	"main/internal/signal_logic"
	"net/http"
)

type Request struct {
	Candles [][]string `json:"candles"`
}

func OHLCHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Candles [][]string `json:"candles"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if len(req.Candles) == 0 {
		http.Error(w, "empty candles", 400)
		return
	}

	data, err := signal_logic.ParseOHLCFromRaw(req.Candles)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ind := indicators.GetFinalIndicators(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ind)
}
