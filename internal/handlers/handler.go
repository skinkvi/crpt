// pkg/handlers/handlers.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/skinkvi/crpt/internal/crypto"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func InitHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/crypto", func(w http.ResponseWriter, r *http.Request) {
		cryptoName := r.URL.Query().Get("name")
		if cryptoName == "" {
			http.Error(w, "Missing crypto name", http.StatusBadRequest)
			return
		}

		cryptoData, err := crypto.GetCryptoData(cryptoName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cryptoData)
	})

	logger.GetLogger().Info("Handlers initialized")
}
