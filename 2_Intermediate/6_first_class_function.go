/*
6. First-class functions :
In Go, functions are first-class values.

This means functions can be:
a) Assigned to variables.
b) Passed as arguments to other functions.
c) Returned from other functions.
d) Stored in data structures, subject to their types.
*/

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	operation := add
	fmt.Println(operation(5, 7)) // Output: 12
}