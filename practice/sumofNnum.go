package practice

import (
	"fmt"
)

func SumofN() {

	var n, sum int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Print("Numbers are: ")
	for i := 1; i <= n; i++ {
		sum += i
		fmt.Print(i, " ")

	}

	fmt.Println("\nSum is:", sum)

}
