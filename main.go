package main

import (
	"github.com/gin-gonic/gin"
	r "movie_star/repository"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.NewClient()

	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080

}
