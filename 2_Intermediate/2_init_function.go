package main

import "fmt"

func init() {
    fmt.Println("Init function executed")
}

func main() {
    fmt.Println("Main function executed")
}

// Multiple init functions

// func init() {
//     fmt.Println("First init")
// }

// func init() {
//     fmt.Println("Second init")
// }

// func main() {
//     fmt.Println("Main")
// }