package practice

import "fmt"

func Factorial() {

	var n int
	fmt.Println("Enter any number: ")
	fmt.Scan(&n)

	start := 1
	if n < 0 {
		fmt.Print("Invalid number")
		return
	}
	fmt.Print("The numbers are: ")
	for i := 1; i <= n; i++ {
		start *= i
		fmt.Print(i, " ")

	}
	fmt.Println("\nFactorial is: ", start)
}
