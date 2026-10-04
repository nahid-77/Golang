/*
3. Anonymous Function :
An anonymous function is a function that has no name.
You can define it wherever a function value is expected.
*/
package main

import "fmt"

func main() {
    message := func() {
        fmt.Println("Hello from an anonymous function")
    }

	// Anonymous function with parameters :
	/*
	func(name string) {
        fmt.Println("Hello,", name)
    }("Nahid")
	*/

    message()
}



