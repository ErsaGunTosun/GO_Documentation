package main

import "fmt"

func main() {

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["three"])

	// eleman kontrol etme
	value, ok := myMap["six"]
	fmt.Println(value, ok)

	myMap["seven"] = 7
	fmt.Println(myMap)

	delete(myMap, "seven")
	fmt.Println(len(myMap))

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	fmt.Println("Maps Finish")
}
