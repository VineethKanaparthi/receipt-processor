package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetPoints(t *testing.T) {
	server := NewReceiptServer(":memory:")
	defer server.CloseDB()
	// Test /receipts/:id/points endpoint with invalid id
	t.Run("GET /receipts/:id/points", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/123/points", nil)
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test /receipts/:id/points endpoint with non existent id
	t.Run("GET /receipts/:id/points", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/d49ae048-61cc-4236-a258-1c4b3c2362ab/points", nil)
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestProcessReceipt(t *testing.T) {
	server := NewReceiptServer(":memory:")
	defer server.CloseDB()
	// Test /receipts/process endpoint invalid json
	t.Run("POST /receipts/process", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", nil)
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test /receipts/process endpoint invalid total
	t.Run("POST /receipts/process", func(t *testing.T) {
		// Mock receipt JSON for testing
		receiptJSON := `{
			"retailer": "Target",
			"purchaseDate": "2022-01-01",
			"purchaseTime": "13:01",
			"items": [
			  {
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
			  },{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
			  },{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			  },{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			  },{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
			  }
			],
			"total": "a35.00"
		  }`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
		req.Header.Set("Content-Type", "application/json")
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test /receipts/process endpoint invalid item price
	t.Run("POST /receipts/process", func(t *testing.T) {
		// Mock receipt JSON for testing
		receiptJSON := `{
			"retailer": "Target",
			"purchaseDate": "2022-01-01",
			"purchaseTime": "13:01",
			"items": [
				{
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
				},{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
				},{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
				},{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
				},{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "a12.00"
				}
			],
			"total": "35.00"
			}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
		req.Header.Set("Content-Type", "application/json")
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test /receipts/process endpoint invalid date
	t.Run("POST /receipts/process", func(t *testing.T) {
		// Mock receipt JSON for testing
		receiptJSON := `{
				"retailer": "Target",
				"purchaseDate": "2022-01-62",
				"purchaseTime": "13:01",
				"items": [
				  {
					"shortDescription": "Mountain Dew 12PK",
					"price": "6.49"
				  },{
					"shortDescription": "Emils Cheese Pizza",
					"price": "12.25"
				  },{
					"shortDescription": "Knorr Creamy Chicken",
					"price": "1.26"
				  },{
					"shortDescription": "Doritos Nacho Cheese",
					"price": "3.35"
				  },{
					"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
					"price": "12.00"
				  }
				],
				"total": "35.00"
			  }`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
		req.Header.Set("Content-Type", "application/json")
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test /receipts/process endpoint invalid time
	t.Run("POST /receipts/process", func(t *testing.T) {
		// Mock receipt JSON for testing
		receiptJSON := `{
				"retailer": "Target",
				"purchaseDate": "2022-01-01",
				"purchaseTime": "26:01",
				"items": [
				  {
					"shortDescription": "Mountain Dew 12PK",
					"price": "6.49"
				  },{
					"shortDescription": "Emils Cheese Pizza",
					"price": "12.25"
				  },{
					"shortDescription": "Knorr Creamy Chicken",
					"price": "1.26"
				  },{
					"shortDescription": "Doritos Nacho Cheese",
					"price": "3.35"
				  },{
					"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
					"price": "12.00"
				  }
				],
				"total": "35.00"
			  }`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
		req.Header.Set("Content-Type", "application/json")
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("POST /receipts/process with valid JSON", func(t *testing.T) {
		// Mock receipt JSON for testing
		receiptJSON := `{
			"retailer": "Target",
			"purchaseDate": "2022-01-01",
			"purchaseTime": "13:01",
			"items": [
			  {
				"shortDescription": "Mountain Dew 12PK",
				"price": "6.49"
			  },{
				"shortDescription": "Emils Cheese Pizza",
				"price": "12.25"
			  },{
				"shortDescription": "Knorr Creamy Chicken",
				"price": "1.26"
			  },{
				"shortDescription": "Doritos Nacho Cheese",
				"price": "3.35"
			  },{
				"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
				"price": "12.00"
			  }
			],
			"total": "35.00"
		  }`

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(receiptJSON)))
		req.Header.Set("Content-Type", "application/json")
		server.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Unmarshal response to check if it contains 'id'
		receiptResponse := decodeResponse(w, t)
		assertUUID(receiptResponse.ID, t)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/receipts/"+receiptResponse.ID+"/points", nil)
		server.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]int
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assertNoErrorWhileDecodingJson(err, t, w)
		assert.Contains(t, response, "points")
	})

}

func decodeResponse(response *httptest.ResponseRecorder, t testing.TB) ReceiptResponse {
	t.Helper()
	var got ReceiptResponse
	err := json.NewDecoder(response.Body).Decode(&got)
	assertNoErrorWhileDecodingJson(err, t, response)
	return got
}

func assertContentTypeJson(response *httptest.ResponseRecorder, t testing.TB) {
	t.Helper()
	if response.Result().Header.Get("content-type") != "application/json" {
		t.Errorf("Expected content type json but found %v", response.Result().Header.Get("content-type"))
	}
}

func assertUUID(id string, t testing.TB) {
	t.Helper()
	_, err := uuid.Parse(id)
	if err != nil {
		t.Errorf("Expected a uuid but found %s", id)
	}
}

func assertNoErrorWhileDecodingJson(err error, t testing.TB, response *httptest.ResponseRecorder) {
	t.Helper()
	if err != nil {
		t.Errorf("Failure decoding json '%s', got error: %+v", response.Body.String(), err.Error())
	}
}

func assertStatusCode(response *httptest.ResponseRecorder, status int, t testing.TB) {
	t.Helper()
	if response.Code != status {
		t.Errorf("Expected response code %d, but got %d", status, response.Code)
	}
}
