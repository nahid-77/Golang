// 1. Local Scope:
// A variable declared inside a function can generally be accessed only within that function.

package main

import "fmt"

func main() {
    name := "Nahid"
    age := 23

    fmt.Println(name)
    fmt.Println(age)
}

func display() {
    fmt.Println(name) // ERROR!
}
// Because name is declared inside main(). The display() function cannot access it.
// Remember: A variable declared inside a function is accessible only within its scope.