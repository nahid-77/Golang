// 2. Block Scope
// A block is a section of code enclosed in curly braces {}.
// Variables declared inside a block are accessible within that block and its nested blocks, but not outside it.

package main

import "fmt"

func main() {
    age := 23

    if age >= 18 {
        message := "You are an adult."
        fmt.Println(message)
    }

     fmt.Println(message) // ERROR!
}