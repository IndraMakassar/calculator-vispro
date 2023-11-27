package calculator

import "calculatorWeb/mathops"

func Add(firstNumber, secondNumber float64) float64 {
	return mathops.Add(firstNumber, secondNumber)
}

func Subtract(firstNumber, secondNumber float64) float64 {
	return mathops.Subtract(firstNumber, secondNumber)
}

func Multiply(firstNumber, secondNumber float64) float64 {
	return mathops.Multiply(firstNumber, secondNumber)
}

func Divide(firstNumber, secondNumber float64) float64 {
	hasil, _ := mathops.Divide(firstNumber, secondNumber)
	return hasil
}

func Power(firstNumber, secondNumber float64) float64 {
	result := 1.0
	for i := 0.0; i < secondNumber; i++ {
		result = Multiply(result, firstNumber)
	}
	return result
}
