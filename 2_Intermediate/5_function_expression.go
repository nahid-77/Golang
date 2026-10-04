/*
5. Function Expression (Assigning a Function to a Variable) :
A function expression, as the term is commonly used in Go tutorials, is an anonymous function assigned to a variable.
Go treats functions as values, so a function can be stored in a variable and called through that variable.
*/

package main

import "fmt"

func main() {

	greet := func(name string) string {
        return "Hello, " + name
    }

    add := func(a int, b int) int {
        return a + b
    }

    result := add(10, 20)

	fmt.Println(greet("Nahid"))
    fmt.Println(result)
}