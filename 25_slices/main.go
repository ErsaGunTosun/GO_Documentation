package main

import "fmt"

func main() {
	// varsayılan değerleri yoktur
	// Değer yapısı esnektir
	// Boyutu sonradan küçültülebilir veya büyültülebilir
	// Yeni bir slice tanımlamak için make() fonksiyonu kullanılır

	var mySlc []int

	fmt.Println(mySlc)
	fmt.Println(len(mySlc))
	fmt.Println(mySlc == nil)

	// make() ile slice oluşturmadan slice değer ataması yapılamaz
	mySlc = make([]int, 5)
	fmt.Println(mySlc)
	fmt.Println(len(mySlc))
	fmt.Println(mySlc == nil)

	// slices call by reference ile çalışır
	// arrays call by value ile çalışır

	var intSlc1 = []int{1, 2, 3}
	var intSlc2 = intSlc1

	fmt.Println(intSlc1)
	fmt.Println(intSlc2)
	intSlc2[2] = 10
	fmt.Println(intSlc1)
	fmt.Println(intSlc2)

}
