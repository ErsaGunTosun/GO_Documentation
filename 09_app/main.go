package main

func main() {
	// 1- x :int y : float64 type conversion
	/*
		x := 10
		fmt.Printf("x: %d, type: %T \n", x, x)
		y := (float64)(x)
		fmt.Printf("y: %f, type: %T \n", y, y)
	*/

	// 2- multiple assignment x, y = y, x
	/*
		x := 3
		y := 7
		fmt.Printf("x: %d, y: %d \n", x, y)
		x, y = y, x // multiple assignment
		fmt.Printf("x: %d, y: %d \n", x, y)
	*/

	// 3- non english variable names
	/*
		yaş := 10
		возраст := 20

		fmt.Println(yaş)
		fmt.Println(возраст)
	*/

	// 4- 37 as a string
	/*
		x := 37
		fmt.Printf("x: %d, type: %T \n", x, x)
		y := string(x)
		fmt.Printf("y: %v, type: %T \n", y, y)
	*/
}
