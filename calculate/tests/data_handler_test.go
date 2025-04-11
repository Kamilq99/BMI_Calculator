package handlers

import (
	"bytes"
	"calculate-bmi/handlers"
	"calculate-bmi/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test for valid data passed to the DataHandler
func TestDataHandler_ValidData(t *testing.T) {
	// Set up the Gin router
	router := gin.Default()
	router.POST("/data", handlers.DataHandler)

	// Prepare the valid input data
	human := models.Human{
		Weight: 70,
		Height: 1.75,
	}
	body, _ := json.Marshal(human)

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Send the request to the router
	router.ServeHTTP(w, req)

	// Assert the response status code and body content
	assert.Equal(t, 200, w.Code)                                                            // Check if the status is 200 OK
	assert.Contains(t, w.Body.String(), `"weight":70`)                                      // Check if the response contains weight
	assert.Contains(t, w.Body.String(), `"height":1.75`)                                    // Check if the response contains height
	assert.Contains(t, w.Body.String(), `"message":"All data has been saved successfully"`) // Check if the success message is present
}

// Test for invalid data passed to the DataHandler
func TestDataHandler_InvalidData(t *testing.T) {
	// Set up the Gin router
	router := gin.Default()
	router.POST("/data", handlers.DataHandler)

	// Prepare invalid input data (missing "weight" field)
	body := []byte(`{"height": 1.75}`)

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Send the request to the router
	router.ServeHTTP(w, req)

	// Assert the response status code and error message
	assert.Equal(t, 400, w.Code)                   // Check if the status is 400 Bad Request
	assert.Contains(t, w.Body.String(), `"error"`) // Check if the response contains an error message
}
