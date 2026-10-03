package main

import (
    "fmt"
    "example/mathlib"
)

var name = "Nahid"

var (
    a = 10
    b = 20
)

func main() {
    fmt.Println(name)

    result := mathlib.Add(a, b)

    fmt.Println(result)
}