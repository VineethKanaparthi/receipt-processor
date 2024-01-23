package main

import "github.com/VineethKanaparthi/receipt-processor/internal/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
