package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))

	s.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
}
