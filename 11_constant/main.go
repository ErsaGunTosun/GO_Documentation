package main

import (
	"fmt"
	"math"
)

func main() {
	// constant: sabit değerler için kullanılır.
	// programn içinde sonradan değiştirilemezler.

	// const -> compile time = bilgisayarlar compile edilirken sabit değerlerin yerine yazılır.
	// var -> run time = program çalıştırılırken değerlerin yerine yazılır.

	// const x = math.Pow(5, 2) // hata verir. Fonksiyonlar run time da çalışır. Bundan dolayı const değişkene atılamaz.

	const pi float64 = 3.14

	area := pi * math.Pow(5, 2)

	fmt.Println(area)
}
