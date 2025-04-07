// handlers/request_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Omotolani98/k8ly/services/reqly/models"
	"gorm.io/gorm"
)

func HandleWebhook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := io.ReadAll(r.Body)
		headers, _ := json.Marshal(r.Header)

		req := models.LoggedRequest{
			Method:    r.Method,
			URL:       r.RequestURI,
			Headers:   string(headers),
			Body:      string(bodyBytes),
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		if result := db.Create(&req); result.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "❌ Error saving request: %v", result.Error)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "✅ Webhook received")
	}
}

func HandleInspect(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var logs []models.LoggedRequest
		db.Order("created_at desc").Limit(50).Find(&logs)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logs)
	}
}

func HandleClear(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db.Exec("DELETE FROM logged_requests")
		fmt.Fprintln(w, "🧹 Cleared all requests")
	}
}
