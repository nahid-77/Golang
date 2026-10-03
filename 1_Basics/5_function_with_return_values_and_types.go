package main
import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main() {
	a := 30
	b := 10
	result := add(a, b)
	fmt.Println(result)

	x, y := getNumbers(a, b)
	fmt.Println(x, y)
}