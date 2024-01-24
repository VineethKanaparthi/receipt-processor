package main

import (
	"log"
	"time"

	"github.com/VineethKanaparthi/receipt-processor/internal/router"
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
	r := router.SetupRouter(db)
	r.Run(":8080")
}
