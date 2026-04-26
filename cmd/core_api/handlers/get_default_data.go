package handlers

import (
	"encoding/json"
	"main/internal"
	"net/http"
)

func Handler_default_data(cm *internal.CoinMap) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cm.Mu.RLock()
		defer cm.Mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cm.Coins)
	}
}
