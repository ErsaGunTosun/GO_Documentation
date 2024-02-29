package main

import "fmt"

func main() {
	//cities := [4]string{"New York", "Boston", "Atlanta", "Los Angeles"}
	//cities := []string{"New York", "Boston", "Atlanta", "Los Angeles"}
	cities := [...]string{"New York", "Boston", "Atlanta", "Los Angeles"}

	fmt.Println(cities)
	fmt.Println(cities[0])
	fmt.Println(cities[3])

	var myInt [5]int

	fmt.Println(myInt)
	fmt.Printf("Array Length: %v \n", len(myInt))

	var myInt2 [5]int
	fmt.Println(myInt2)

	if myInt == myInt2 {
		fmt.Println("equals")
	}

	// for

	for i := 0; i < len(cities); i++ {
		fmt.Printf("%v. : %s \t", i, cities[i])
	}
	fmt.Print("\n")

	myArr := [5]int{1, 2, 3, 4, 5}

	myArr = mySquare(myArr)

	// for range

	for _, item := range myArr {
		fmt.Printf("%v \t", item)
	}
	fmt.Print("\n")

	fmt.Println("Arrays Finish")
}

func mySquare(arr [5]int) [5]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}

	return arr
}
