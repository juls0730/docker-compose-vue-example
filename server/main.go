package main

import (
	"context" // [!code hl:2]
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9" // [!code hl]
)

var ctx = context.Background() // [!code hl]

func main() {
	client := redis.NewClient(&redis.Options{ // [!code hl:5]
		Addr:     "keydb:6379",
		Password: "",
		DB:       0,
	})

	r := gin.Default()

	// This route returns a count from the database and increments it for every request
	r.GET("/count", func(c *gin.Context) {
		count, err := client.Incr(ctx, "count").Result() // [!code hl:8]
		if err != nil {
			log.Println("Error incrementing count in Redis:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment count"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"count": count})
	})

	r.Run(":8080")
}
