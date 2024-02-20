package main

import "fmt"

func main() {
	// 1- 1 ile 10 arasındaki sayıları tek - çift olarak ekrana yazdırın
	/*
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i, "even")
			} else {
				fmt.Println(i, "odd")
			}
		}
	*/

	// 2-  for yapısını kullanarak go'da olmayan while döngüsünü nasıl yaparız?
	/*
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}
	*/

	// 3- switch fallthrough ifadesini açıklayınız
	// fallthrough ifadesi kullanıldığında bir case bloğu çalıştıktan sonra bir sonraki case bloğu da çalıştırılır.
	/*
		switch x := 7; {
		case x < 5:
			fmt.Println("x < 5")
			fallthrough
		case x < 10:
			fmt.Println("x < 10")
			fallthrough
		case x < 20:
			fmt.Println("x < 20")
		}
	*/

	// 4- if döngüsünü daha idiomatic nasıl yazabiliriz?
	// idiomatic: Go dilinde yazılan kodların genel olarak kabul görmüş, yaygın olarak kullanılan, anlaşılabilir, okunabilir, performanslı ve mantıklı olması durumudur.
	/*
		if x := 20; x%2 == 0 {
			fmt.Println(x, "even")
		} else {
			fmt.Println(x, "odd")
		}
	*/
	x := 20
	if x%2 == 0 {
		fmt.Println(x, "even")
		return
	}

	fmt.Println(x, "odd")

	fmt.Println("App Finish")
}
