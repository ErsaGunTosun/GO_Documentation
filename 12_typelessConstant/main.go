package main

import "fmt"

func main() {
	// Typeless constant
	// Sabitin değeri belirtilen bir sabit türü yoktur. Sabitlerin türü, kullanıldıkları bağlamdan belirlenir.

	/*
		const x int8 = 10
		var y int16 = 15
		fmt.Printf("Type: %T, Value: %v\n", x, x)
		fmt.Printf("Type: %T, Value: %v\n", y, y)
		fmt.Printf("Type: %T,  Value: %v\n", x+y, x+y)

		Bu şekilde yapıldığında x int8 ve y int16 olduğu için toplama işlemi yapamaz AMA
	*/

	const x = 10
	var y int16 = 15

	fmt.Printf("Type: %T, Value: %v\n", x, x) // x'in herhangi bir türü olmadığı için default olan tür yani int türünü atar
	fmt.Printf("Type: %T, Value: %v\n", y, y)
	fmt.Printf("Type: %T,  Value: %v\n", x+y, x+y) // x'i typeless constant olarak ayarladığımız i.çin toplama işlemi yapabilir
	// x burada type conversion yapar ve y'nin türüne dönüşür ve toplama işlemi yapar

	fmt.Println("Typeless constant finish")
}
