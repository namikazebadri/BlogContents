package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file not found")
	}

	envVars, err := godotenv.Read()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Hello from " + envVars["SERVER_NAME"],
		})
	})

	errRun := router.Run()

	if errRun != nil {
		return
	}
}
