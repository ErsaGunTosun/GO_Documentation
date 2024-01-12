package main

import "fmt"

func main() {

	x, y := 15, 10

	fmt.Printf("%d, %d; %T, %T\n", x, y, x, y)
	fmt.Printf("%d, %T\n", (x + y), (x + y))
	fmt.Printf("%d, %T\n", (x - y), (x - y))
	fmt.Printf("%d, %T\n", (x * y), (x * y))
	fmt.Printf("%d, %T\n", (x / y), (x / y)) // Burada iki değerde int olduğu için sonuç int olur. 1.5 değil 1 olur.
	var z float64 = float64(x) / float64(y)
	fmt.Printf("%f, %T\n", z, z) // Burada float64 olarak tanımladığımız için sonuç 1.500000 olur.

	x++
	fmt.Printf("%d, %T\n", x, x)
	// fmt.Printf("%d, %T\n",x++,x++) şeklinde olursa çalışmaz nedeni x++ bir statementtir. Expression değildir.
}
