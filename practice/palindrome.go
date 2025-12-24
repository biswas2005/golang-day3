package practice

import "fmt"

func Palindrome() {

	var n int
	fmt.Println("<Enter a number to check if it's palindrome>")
	fmt.Scan(&n)

	start := 0
	original := n

	for n > 0 {
		digit := n % 10
		start = start*10 + digit
		n = n / 10
	}
	fmt.Println(start)
	if start == original {
		fmt.Println("Is a palindrome.")
	} else {
		fmt.Println("It is not a palindrome.")
	}
}
