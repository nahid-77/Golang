/*
4. IIFE (Immediately Invoked Function Expression) :
An IIFE is an anonymous function that is defined and executed immediately.
In Go, this is commonly written as an anonymous function followed directly by ().
*/
package main

import "fmt"

func main() {
    func() {
        fmt.Println("I am an IIFE")
    }()

	// IIFE with parameters
	func(a int, b int) {
        fmt.Println("Sum:", a+b)
    }(10, 20)
}