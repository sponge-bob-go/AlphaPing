package handlers

import (
	mexc "main/internal/signal_logic/MEXC"
	"net/http"
)

func HandlerTechData(w http.ResponseWriter, r *http.Request) {
	market := r.URL.Query().Get("market")
	timeFrame := r.URL.Query().Get("timeframe")
	coinName := r.URL.Query().Get("coin_name")

	if market == "" {
		w.WriteHeader(403)
	}

	if timeFrame == "" {
		w.WriteHeader(403)
	}

	if coinName == "" {
		w.WriteHeader(403)
	}

	if market == "mexc" {
		mexcModel := mexc.CreateMEXCModel()
		mexcOHLC, err := mexcModel.GetOHLC(coinName, timeFrame)
		if err != nil {
			w.WriteHeader(403)
		}
	}
}
