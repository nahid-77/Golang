package main
import "fmt"

func printSomething(){
	fmt.Println("This is a function that prints something.")
}

func sayHello(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	printSomething()
	sayHello("Nahid Mondol")
}