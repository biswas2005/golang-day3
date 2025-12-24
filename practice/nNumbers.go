package practice

import "fmt"

func Numbers() {
	var num int
	fmt.Println("Enter any number: ")
	fmt.Scanln(&num)

	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}
