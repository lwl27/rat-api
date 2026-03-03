package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogHandler_MissingMessage(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "message is required")
}

func TestLogHandler_QueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=info&user=john", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ok")
	assert.Contains(t, w.Body.String(), "test")
}

func TestLogHandler_DefaultLevel(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)

	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "ok", resp["status"])
}

func TestLogHandler_DebugLevel(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=debug", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLogHandler_WarnLevel(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=warn", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLogHandler_WarningLevel(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=warning", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLogHandler_ErrorLevel(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=error", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLogHandler_JsonBody(t *testing.T) {
	body := `{"user": "john", "count": 42}`
	req := httptest.NewRequest(http.MethodPost, "/log?message=test&level=info", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)

	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "ok", resp["status"])
	assert.Equal(t, "test", resp["message"])
}

func TestLogHandler_JsonBodyWithInvalidJson(t *testing.T) {
	body := `{"user": "john"`
	req := httptest.NewRequest(http.MethodPost, "/log?message=test", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLogHandler_ResponseContentType(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func TestLogHandler_UppercaseLevel(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=INFO", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLogHandler_IgnoreMessageAndLevelInFields(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/log?message=test&level=info&user=john", nil)
	w := httptest.NewRecorder()

	logHandler(w, req)

	assert.Equal(t, 200, w.Code)
}
