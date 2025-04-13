package handlers

import (
	"calculate-bmi/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBMIHandler_ValidData(t *testing.T) {
	// Ustaw dane globalne
	handlers.Weight = 70
	handlers.Height = 1.75

	router := gin.Default()
	router.GET("/bmi", handlers.CalculateBMIHandler)

	req, _ := http.NewRequest("GET", "/bmi", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"BMI":22.857142857142858`)
}

func TestCalculateBMIHandler_ZeroHeight(t *testing.T) {
	// Ustaw niewłaściwe dane
	handlers.Weight = 70
	handlers.Height = 0

	router := gin.Default()
	router.GET("/bmi", handlers.CalculateBMIHandler)

	req, _ := http.NewRequest("GET", "/bmi", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"error"`)
}
