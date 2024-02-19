package main

import "fmt"

func main() {

	// if <boolean expression> { <statement> } else { <statement> }

	x := 10

	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	fmt.Println("------------")

	x = -3

	if x < 0 {
		fmt.Println("x is negative")
	} else if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	fmt.Println("------------")

	y := 13

	if y := 12; y%2 == 0 {
		fmt.Println("y is even")
	} else {
		fmt.Println("y is odd")
	}
	// yapılan işlemi if bloğu içinde değiştirdiğimiz için y değerinin ana blokta değişmediğini görebiliriz.
	fmt.Printf("y: %v\n", y)

	fmt.Println("Conditional Statements Finish")
}
