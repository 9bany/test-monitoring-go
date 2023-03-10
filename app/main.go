package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")

	p.Use(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello	",
		})
	})
	r.Run("0.0.0.0:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
