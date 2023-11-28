package ui

import (
	"calculatorWeb/calculator"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
)

func StartWebCalculator() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Static("/static", "./static")
	router.POST("/calculate", calculateHandler)
	router.Run(":80")
}

func calculateHandler(c *gin.Context) {
	var request calculator.Request

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		result, err := calculator.Calculate(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": result})
		}
	}
}
