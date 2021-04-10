package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	e := gin.Default()

	e.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"body": "#test proxy",
		})
	})

	e.Run(":12222")
}
