package main

import "fmt"

func main() {

	x := 3
	y := 3.1

	fmt.Printf("x = %v, type of %T\n", x, x)
	fmt.Printf("y = %v, type of %T\n", y, y)

	// Type Conversion: type(value) -> int(value) = int
	fmt.Println(x + int(y)) // Burada veri kaybı olabilir bundan dolayı float ile dönüşüm yapma işlemi yapılır.
	fmt.Println(float64(x) + y)

	q := 3
	p := "3"

	fmt.Printf("x = %v, type of %T\n", q, q)
	fmt.Printf("y = %v, type of %T\n", p, p)

	// fmt.Println(q + int(p))  type conversion string'i int'a döndüremez

	numCode := 99
	characterCode := string(numCode) // int to string

	fmt.Printf("x = %v, type of %T\n", numCode, numCode)
	fmt.Println(characterCode)
	fmt.Printf("y = %v, type of %T\n", characterCode, characterCode)

	fmt.Println("Type Conversion")
}
