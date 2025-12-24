package practice

import "fmt"

func Table() {
	var n int
	fmt.Println("Enter a number to print it's table: ")
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		var result int
		result = n * i
		fmt.Printf("%d*%d=%d\n", n, i, result)
	}
}
