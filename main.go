package cake

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

func Ln(a float32) (float32, string) {
	if a <= 0 {
		return a, "Error: Logarithm of non-positive number"
	} else {

		return a, "Success"
	}
}

//func Acos(a float32) float32 {
//return (a * 3.14 / 180)
//}
