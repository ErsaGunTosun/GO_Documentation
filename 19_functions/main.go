package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* Fonksiyonun amaçları
	Moduler Programing
	Readable Code
	From Complex to Simple
	*/
	x, y := 4, 5
	fmt.Println("Sum of", x, "and", y, "is", sum(x, y))

	hi()
	fmt.Println("Functions Finish")

	/* Fonksiyon isimlendirme
	ilk karakter harf olmalıdır
	camel case kullanılmalıdır myFunction gibi
	paket dışı ise büyük harfle başlamalıdır
	*/

	/* Parametre vs Argüman
	Parametre: Fonksiyon tanımında kullanılan değişken
	Argüman: Fonksiyon çağrısında kullanılan değişken
	*/

	/* Fonskiyon vs Metod
	Fonksiyon: kendi çalıştığımız dosya içinde tanımlanmış fonksiyonlardır
	Metod: Başka bir paketin içindeki fonksiyonlardır
	*/

	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // _ kullanılmazsa hata verir _= blank identifier yani gelen değeri kullanmıyorum demek
	fmt.Println("You entered", input)

	section, remainder := div(10, 3)
	fmt.Println("Section:", section, "Remainder:", remainder)

}

// func function_name( <parameter list> ) <return_types> {<code>}

func sum(x, y int) int {
	return x + y
}

func hi() {
	fmt.Println("Hi")
}

func div(x, y int) (section, remainder int) { // multiple return
	section = x / y
	remainder = x % y
	return section, remainder
}
