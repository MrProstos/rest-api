package server

import (
	"encoding/json"
	"net/http"
	"time"
)

func parseDate(str string) (string, error) {
	date, err := time.Parse("2006-01-02", str)
	if err != nil {
		return "", err
	}
	str = date.Format("2006-01-02")
	return str, nil
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
