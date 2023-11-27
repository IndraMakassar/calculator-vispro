package ui

import (
	"calculatorWeb/calculator"
	"errors"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	router.Run(":8080")
}

func calculateHandler(c *gin.Context) {
	var num1, num2 float64
	var operation string

	// Get data from the request
	num1, _ = strconv.ParseFloat(c.PostForm("num1"), 64)
	num2, _ = strconv.ParseFloat(c.PostForm("num2"), 64)
	operation = c.PostForm("operation")

	// Perform the calculation
	result, err := calculate(float64(num1), float64(num2), operation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return calculator.Add(a, b), nil
	case "-":
		return calculator.Subtract(a, b), nil
	case "*":
		return calculator.Multiply(a, b), nil
	case "/":
		return calculator.Divide(a, b), nil
	case "^":
		return calculator.Power(a, b), nil
	default:
		return 0, errors.New("invalid operation")
	}
}
