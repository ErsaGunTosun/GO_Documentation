package main

import "fmt"

func main() {

	// 1- x = x -10 vs x -= 10
	/*
		x := 52

		x = x - 10 // assingmnet statement

		x -= 10 // asignmnet operator

		fmt.Printf("Type: %T, Value %v\n", x, x)
	*/

	// 2- K = (F - 32) / 1.8 + 273  -> F = -40 ise K kaçtır?
	/*
		F := -40

		k := float64(F-32)/1.8 + 273

		fmt.Printf("Type: %T, Value: %v\n", k, k)
	*/

	// 3-
	/*
		age := 20

		const myAge = age // constant compile time'da belirlenir. Bu yüzden age değişkeni const değişkenine atanamaz.

		fmt.Printf("Type: %T, Value: %v\n", myAge, myAge)
	*/

	// 4 const x float64 = 6.4, y := 4 + x, y =?

	const x float64 = 6.4

	y := 4 + x

	fmt.Printf("Type: %T, Value: %v\n", y, y)

	fmt.Println("App finish")

}
