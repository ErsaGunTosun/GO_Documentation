package main

import "fmt"

func main() {

	underArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(underArray)

	mySlc := underArray[2:5] // 2 3 4 indexteki elemanları alır.

	mySlc2 := underArray[:5] // 0 1 2 3 4 indexteki elemanları alır.
	fmt.Println(mySlc2)

	mySlc3 := underArray[4:] // 5 6 7 8 9 indexteki elemanları alır.
	fmt.Println(mySlc3)

	fmt.Println("--------------------")
	mySlc2[4] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc3)

	// eleman birleştirme
	slc1 := []int{1, 2, 3}
	slc2 := []int{4, 5, 6}

	slc1 = append(slc1, slc2...) // ... kullanma sebebimiz bütün elemanları teker teker almak.

	fmt.Println("Slices 2 Finish")
}
