/*
8. Callback Function :
A callback is a function passed to another function so that the receiving function can call it at an appropriate time.
Callbacks are possible in Go because functions are first-class values.
*/
package main
import "fmt"

func process_2(callback_2 func()) {
    fmt.Println("Before callback")

    callback_2()

    fmt.Println("After callback")
}

func process(name string, callback func(string)) {
	fmt.Println("Processing started")

	callback(name)

	fmt.Println("Processing completed")
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	process("Nahid", greet) 
	process_2(func() {
        fmt.Println("Callback executed")
    })
}

// Output:
// Processing started
// Hello Nahid
// Processing completed