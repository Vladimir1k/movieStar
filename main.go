package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"movie_star/internal/director"
	"movie_star/repository"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/POST|new_director", func(c *gin.Context) {
		var person director.Director

		err := c.BindJSON(&person)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		conn, err := repository.NewClient()
		defer conn.Close(context.Background())

		row := conn.QueryRow(context.Background(),
			`INSERT INTO directors (name, date_of_birth, bio) 
				VALUES ($1, $2, $3) RETURNING id, name, date_of_birth, bio`,
			person.Name, person.DateOfBirth.Format("2006-01-02"), person.Bio)

		err = row.Scan(&person.Id, &person.Name, &person.DateOfBirth, &person.Bio)
		if err != nil {
			log.Fatalf("Unable to INSERT: %v\n", err)
		}

		c.JSON(200, gin.H{"id": person.Id, "name": person.Name,
			"date_of_birth": person.DateOfBirth.Format("2006-01-02"), "bio": person.Bio})
	})

	router.Run("localhost:8080")
}
