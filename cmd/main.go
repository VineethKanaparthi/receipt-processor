package main

import (
	"github.com/VineethKanaparthi/receipt-processor/internal/server"
)

// main function initializes the database and sets up the router to handle API requests.
func main() {
	server := server.NewReceiptServer("receipts.db")
	server.Run(":8080")
	defer server.CloseDB()
}
