package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"lograt"
)

var logger = lograt.New()

func main() {
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	fields := make(map[string]any)

	for k, v := range r.URL.Query() {
		if k == "message" || k == "level" {
			continue
		}
		if len(v) > 0 {
			fields[k] = v[0]
		}
	}

	message := r.URL.Query().Get("message")
	level := r.URL.Query().Get("level")
	if level == "" {
		level = "info"
	}

	if r.Header.Get("Content-Type") == "application/json" {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
			for k, v := range body {
				if k != "message" && k != "level" {
					fields[k] = v
				}
			}
		}
	}

	if message == "" {
		http.Error(w, `{"error": "message is required"}`, 400)
		return
	}

	switch strings.ToLower(level) {
	case "debug":
		logger.Debug(message, fields)
	case "warn", "warning":
		logger.Warn(message, fields)
	case "error":
		logger.Error(message, fields)
	default:
		logger.Info(message, fields)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": message,
	})
}
