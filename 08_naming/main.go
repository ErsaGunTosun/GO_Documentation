package main

import "fmt"

func main() {
	// değişken isimlendirilirken go'da geçen özel anlamlı kelimeler kullanılmaz
	// örneğin: string, int, float64, bool, true, false, iota, nil, rune,

	// camel-case isimlendirme kullanılır
	var cointType string
	var coinName string

	// kısaltmalar büyük harfle yazılır

	var URL string
	var HTTP string

	fmt.Println(coinName, cointType, URL, HTTP)
	fmt.Println("Finish")
}
