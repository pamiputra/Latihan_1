package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Pami Putra",
			"bio":  "Software Enginer",
		})
	})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Pami Putra",
			"bio":  "Software Enginer",
		})
	})
	router.Run()
}
