package service

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	model "github.com/VineethKanaparthi/receipt-processor/pkg"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

// ErrIdNotFound is an error indicating that the ID was not found in the database.
var ErrIdNotFound = errors.New("id not found")

// ProcessReceipt processes a receipt, calculates points, and stores the points in the database.
func ProcessReceipt(receipt *model.Receipt, db *bolt.DB) (string, error) {
	fmt.Printf("%+v\n", receipt)
	points := calculatePoints(receipt)
	fmt.Println(points)
	id := uuid.New().String()
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("points"))
		return bucket.Put([]byte(id), []byte(strconv.Itoa(points)))
	})
	if err != nil {
		return id, err
	}
	return id, nil
}

// Calculate points for a receipt based on the defined rules
func calculatePoints(receipt *model.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	points += countAlphanumericCharacters(receipt.Retailer)
	fmt.Printf("Rule 1: %d\n", points)
	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if cents := receipt.Total[len(receipt.Total)-2:]; cents == "00" {
		points += 50
	}
	fmt.Printf("Rule 2: %d\n", points)

	// Rule 3: 25 points if the total is a multiple of 0.25
	totalFloat := 0.0
	fmt.Sscanf(receipt.Total, "%f", &totalFloat)
	fmt.Println(totalFloat)
	if totalFloat/0.25 == float64(int(totalFloat/0.25)) {
		points += 25
	}
	fmt.Printf("Rule 3: %d\n", points)

	// Rule 4: 5 points for every two items on the receipt
	points += 5 * (len(receipt.Items) / 2)
	fmt.Printf("Rule 4: %d\n", points)

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2
	// and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.Trim(item.ShortDescription, " "))
		if trimmedLength%3 == 0 {
			priceFloat, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(priceFloat * 0.2))
		}
	}
	fmt.Printf("Rule 5: %d\n", points)

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDay, _ := strconv.Atoi(receipt.PurchaseDate[len(receipt.PurchaseDate)-2:])
	fmt.Println(purchaseDay)
	if purchaseDay%2 != 0 {
		points += 6
	}
	fmt.Printf("Rule 6: %d\n", points)

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	fmt.Printf("%+v\n", purchaseTime)
	if purchaseTime.After(time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)) &&
		purchaseTime.Before(time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)) {
		points += 10
	}
	fmt.Printf("Rule 7: %d\n", points)

	return points
}

// countAlphanumericCharacters counts the number of alphanumeric characters in a string.
func countAlphanumericCharacters(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

// GetPoints retrieves points from the database based on the provided ID.
func GetPoints(id string, db *bolt.DB) (int, error) {
	var points int
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("points"))
		data := bucket.Get([]byte(id))
		if data == nil {
			return ErrIdNotFound
		}
		var converr error
		points, converr = strconv.Atoi(string(data))
		return converr
	})

	return points, err
}
