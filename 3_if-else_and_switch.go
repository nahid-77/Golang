package main
import "fmt"

func main() {
	// relational operators
	// equal to -> ==,
	// not equal to -> !=,
	// greater than -> >,
	// less than -> <,
	// greater than or equal to -> >=,
	// less than or equal to -> <=

	// age := 20

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age < 18 {
	// 	fmt.Println("You are not an adult.")
	// } else{
	// 	fmt.Println("Invalid age.")
	// }

	// logical operators
	// and -> &&,
	// or -> ||,
	// not -> !

	// age = 24
	// if age >= 18 && age <= 30 {
	// 	fmt.Println("You are a young adult.")
	// } else if age < 18 || age > 30 {
	// 	fmt.Println("You are not a young adult.")
	// } else{
	// 	fmt.Println("Invalid age.")
	// }

	// isStudent := true
	// if !isStudent {
	// 	fmt.Println("You are a student.")
	// } else {
	// 	fmt.Println("You are not a student.")
	// }

	// switch statement

	day := "Monday"
	
	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	case "Friday":
		fmt.Println("Today is Friday.")
	case "Saturday":
		fmt.Println("Today is Saturday.")
	case "Sunday":
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Invalid day.")
	}
}