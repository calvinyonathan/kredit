package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.Run(":8080")
}
