package handler

import (
	"encoding/json"

	"net/http"

	"github.com/Omotolani98/k8ly/services/logly/models"
	"gorm.io/gorm"
)

type LogHandler struct {
	DB *gorm.DB
}

func NewLogHandler(db *gorm.DB) *LogHandler {
	return &LogHandler{DB: db}
}

func (h *LogHandler) CreateLog(w http.ResponseWriter, r *http.Request) {
	var logEntry models.Log
	if err := json.NewDecoder(r.Body).Decode(&logEntry); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&logEntry).Error; err != nil {
		http.Error(w, "Failed to create log entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(logEntry)
}

func (h *LogHandler) GetLogs(w http.ResponseWriter, r *http.Request) {
	var logs []models.Log
	if err := h.DB.Find(&logs).Error; err != nil {
		http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)
}
