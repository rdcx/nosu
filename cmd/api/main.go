package main

import (
	"log"
	"nosu/pkg/domain"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

type RegisterDomainRequest struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func main() {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a writable transaction.
	tx, _ := db.Begin(true)
	defer tx.Rollback()
	// Use the transaction...
	tx.CreateBucket([]byte("domains"))
	tx.Commit()

	r := gin.Default()

	r.POST("/domains", func(c *gin.Context) {
		var req RegisterDomainRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		domain, err := domain.Register(req.Name, req.Owner, db)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, domain)
	})

	r.GET("/domains", func(c *gin.Context) {
		doms, err := domain.GetAll(db)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, doms)
	})

	r.Run(":8080")
}
