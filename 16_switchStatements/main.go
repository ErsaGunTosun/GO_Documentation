package main

import "fmt"

func main() {

	grade := 4

	switch grade := 5; grade {
	case 5:
		fmt.Println("Excellent")
	case 4:
		fmt.Println("Good")
	case 3:
		fmt.Println("Average")
	case 2:
		fmt.Println("Poor")
	case 1:
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid Grade")
	}
	fmt.Println(grade)

	fmt.Println("Switch Statements Finish")
}
