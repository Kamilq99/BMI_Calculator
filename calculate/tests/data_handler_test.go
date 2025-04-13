package handlers

import (
	"bytes"
	"calculate-bmi/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDataHandler_ValidData(t *testing.T) {
	router := gin.Default()
	router.POST("/data", handlers.DataHandler)

	body := map[string]float64{
		"weight": 70,
		"height": 1.75,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"weight":70`)
	assert.Contains(t, w.Body.String(), `"height":1.75`)
	assert.Contains(t, w.Body.String(), `"message":"All data has been saved successfully"`)
}

func TestDataHandler_MissingData(t *testing.T) {
	router := gin.Default()
	router.POST("/data", handlers.DataHandler)

	body := []byte(`{"height": 1.75}`)

	req, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}
