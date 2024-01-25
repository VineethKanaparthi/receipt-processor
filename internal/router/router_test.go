package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	bolt "go.etcd.io/bbolt"
)

func TestSetupRouter(t *testing.T) {
	// initialize the database
	db, err := bolt.Open(":memory:", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the points bucket in the database
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("points"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	router := SetupRouter(db)

	// Test /receipts/process endpoint invalid json
	t.Run("POST /receipts/process", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/receipts/process", nil)
		router.ServeHTTP(w, req)

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
		router.ServeHTTP(w, req)

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
		router.ServeHTTP(w, req)

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
		router.ServeHTTP(w, req)

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
		router.ServeHTTP(w, req)

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
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Unmarshal response to check if it contains 'id'
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		fmt.Println(response)
		assert.NoError(t, err)
		assert.Contains(t, response, "id")

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/receipts/"+response["id"]+"/points", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response, "points")
	})

	// Test /receipts/:id/points endpoint with invalid id
	t.Run("GET /receipts/:id/points", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/123/points", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test /receipts/:id/points endpoint with non existent id
	t.Run("GET /receipts/:id/points", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/receipts/d49ae048-61cc-4236-a258-1c4b3c2362ab/points", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
