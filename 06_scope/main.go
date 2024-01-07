package main

import "fmt"

var packvar = "packvar"

/*
Paket değişkenlerinde ":=" yani short declaration kullanılamaz.
short declaration sadece fonksiyon içerisinde kullanılabilir.
Bunun sebebi "var" anahtar kelimesi kullanılması gerektiğinden dolayıdır.
*/

func main() {

	if true {
		var blockvar = "blockvar"
		fmt.Println(blockvar)

	}

	var funcvar = "funcvar"

	fmt.Println(funcvar)
	fmt.Println(packvar)
	// fmt.Println(blockvar) scope alanı sadece if bloğu olduğu için burada kullanılamaz.

	fmt.Println("Scope")
}
