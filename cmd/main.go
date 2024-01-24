package main

import (
	"log"
	"time"

	"github.com/VineethKanaparthi/receipt-processor/internal/handler"
	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

// main function initializes the database and sets up the router to handle API requests.
func main() {
	// initialize the database
	db, err := bolt.Open("receipts.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
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

	// SetupRouter handles API routing
	r := gin.Default()
	// POST /receipts/process endpoint
	r.POST("/receipts/process", func(c *gin.Context) {
		handler.ProcessReceipt(c, db)
	})

	// GET /receipts/:id/points endpoint
	r.GET("receipts/:id/points", func(c *gin.Context) {
		handler.GetPoints(c, db)
	})

	r.Run(":8080")
}
