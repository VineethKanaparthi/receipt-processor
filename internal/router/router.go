package router

import (
	"github.com/VineethKanaparthi/receipt-processor/internal/handler"
	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

func SetupRouter(db *bolt.DB) *gin.Engine {
	r := gin.Default()
	// POST /receipts/process endpoint
	r.POST("/receipts/process", func(c *gin.Context) {
		handler.ProcessReceipt(c, db)
	})

	// GET /receipts/:id/points endpoint
	r.GET("receipts/:id/points", func(c *gin.Context) {
		handler.GetPoints(c, db)
	})
	return r
}
