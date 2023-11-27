package calculator

import (
	"calculatorWeb/mathops"
	"errors"
)

type Request struct {
	FirstNumber  float64 `json:"firstNumber"`
	Operation    string  `json:"operation"`
	SecondNumber float64 `json:"secondNumber"`
}

func Calculate(request Request) (float64, error) {
	switch request.Operation {
	case "+":
		return mathops.Add(request.FirstNumber, request.SecondNumber), nil
	case "-":
		return mathops.Subtract(request.FirstNumber, request.SecondNumber), nil
	case "*":
		return mathops.Multiply(request.FirstNumber, request.SecondNumber), nil
	case "/":
		return mathops.Divide(request.FirstNumber, request.SecondNumber)
	case "^":
		result := 1.0
		for i := 0.0; i < request.SecondNumber; i++ {
			result = mathops.Multiply(result, request.FirstNumber)
		}
		return result, nil
	default:
		return 0, errors.New("invalid operation")
	}
}
