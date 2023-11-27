package mathops

import "fmt"

func Add(firstNumber, secondNumber float64) float64 {
	return firstNumber + secondNumber
}

func Subtract(firstNumber, secondNumber float64) float64 {
	return firstNumber - secondNumber
}

func Multiply(firstNumber, secondNumber float64) float64 {
	return firstNumber * secondNumber
}

func Divide(firstNumber, secondNumber float64) (float64, error) {
	if secondNumber == 0 {
		return 0, fmt.Errorf("can't divide %g by zero", firstNumber)
	}
	return firstNumber / secondNumber, nil
}
