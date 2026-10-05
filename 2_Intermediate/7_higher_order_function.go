/*
7. Higher-order functions :
A higher-order function is a function that accepts another function as an argument, returns a function, or does both.
*/
package main

import "fmt"

func calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(calculate(2, 4, add))	  // Output: 6
	fmt.Println(calculate(7, 9, multiply)) // Output: 63
}

/*
Remember: First-class describes how functions behave in the language. Higher-order describes what a particular function does.
*/