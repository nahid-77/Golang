package main
import "fmt"

func printWelcomeMessage() {
	fmt.Println("Welcome to the application.")
}

func getUserName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func displayResult(name string, sum int) {
	fmt.Println("Hello,", name)
	fmt.Println("The sum is :", sum)
}

func printGoodbyeMessage(name string) {
	fmt.Println("Thank you for using the application.")
	fmt.Println("Goodbye,", name)
}

func main() {
	// Print welcome message
	printWelcomeMessage()
	// Get user name as input
	name := getUserName()
	// Get two numbers as input
	num1, num2 := getTwoNumbers()
	// Calculate the sum 
	sum := add(num1, num2)
	// display the result
	displayResult(name, sum)
	// Print a goodbye message
	printGoodbyeMessage(name)
}