/* 
1. Standard or Named Function :
A named function is a function declared with a name using the func keyword. You can call it whenever you need it.
*/
package main

import "fmt"

func greet() {
    fmt.Println("Hello, Golang!")
}

func add(a int, b int) int {
    return a + b
}

func main() {
	greet()
    result := add(10, 20)
    fmt.Println(result)
}