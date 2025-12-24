package practice

import "fmt"

func Reverse() {
	var n int
	reverse := 0
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)

	for n > 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10

	}
	fmt.Println("Reverse is: ", reverse)
}
