package database

import (
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

// NewBoltDatabase initializes the database
func NewBoltDatabase(dbname string) *bolt.DB {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the points bucket in the database
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("points"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
