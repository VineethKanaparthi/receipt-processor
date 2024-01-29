package main

import (
	"github.com/VineethKanaparthi/receipt-processor/internal/server"
)

// main function initializes and runs the server
func main() {
	server := server.NewReceiptServer("receipts.db")
	server.Run(":8080")
	defer server.CloseDB()
}
