package miniproject

import "fmt"

func Project() {

	for {
		fmt.Println("<Console based number analyzer>")
		fmt.Println("1.Check Even or Odd.")
		fmt.Println("2.Check if Prime or not.")
		fmt.Println("3.Reverse number.")
		fmt.Println("4.Sum of digits.")
		fmt.Println("5.Exit")
		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		var input int
		fmt.Println("Enter any number: ")
		fmt.Scan(&input)

		switch choice {
		case 1:
			if input%2 == 0 {
				fmt.Println("It is a EVEN number.")

			} else {
				fmt.Println("It is a ODD number.")
			}
		case 2:
			if input <= 1 {
				fmt.Println("Invalid number.")
				return

			}
			isPrime := true

			for i := 2; i*i <= input; i++ {
				if input%i == 0 {
					fmt.Println("It is not a PRIME number.")
					isPrime = false
					break

				}

			}
			if isPrime {
				fmt.Println("It is a PRIME number.")
			}
		case 3:
			source := 0

			for input > 0 {
				digit := input % 10
				source = source*10 + digit
				input = input / 10
			}
			fmt.Println(source)
		case 4:

			source1 := 0
			for input > 0 {
				digit := input % 10
				source1 = source1 + digit
				input = input / 10
			}
			fmt.Println(source1)
		case 5:
			fmt.Println("Exiting loop....")
			return
		default:
			fmt.Println("Invalid choice,please choose a valid option.")

		}
	}

}
