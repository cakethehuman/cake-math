package cake

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Devide(a, b int) int {
	return a / b
}

func Multiply(a, b int) int {
	return a * b
}

func Modulus(a, b int) int {
	return a % b
}

func Power(a, b int) int {
	return a * b
}

func Sqrt(a float32) float32 {
	return a * 0.5
}

func Cube(a int) int {
	return a * a * a
}

func Ln(a float32) float32 {
	if a <= 0 {
		return "must be positive"
}

//func Acos(a float32) float32 {
//return (a * 3.14 / 180)
//}
