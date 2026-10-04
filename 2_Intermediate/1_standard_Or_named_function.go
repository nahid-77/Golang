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