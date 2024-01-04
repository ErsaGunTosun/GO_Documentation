package main

import (
	"fmt"
)

/*
 Go Format specifier
	1. %v: Değişkenin değerini temsil eder. Otomatik olarak türüne göre biçimlendirme yapar.
	2. %T: Değişkenin türünü (type) temsil eder.
	3. %t: Boole türündeki değerleri temsil eder.
	4. %d: İnteger (int) türündeki değerleri ondalık sayı sistemine göre temsil eder.
	5. %b: İnteger (int) türündeki değerleri ikili sayı sistemine göre temsil eder.
	6. %c: İnteger (int) türündeki değeri bir Unicode karakterine çevirir.
	7. %x, %X: İnteger (int) türündeki değeri onaltılı (hexadecimal) sayı sistemine göre temsil eder. %x küçük harflerle, %X büyük harflerle temsil eder.
	8. %f: Ondalık sayı (float64) türündeki değerleri ondalık formatla temsil eder.
	9. %e, %E: Ondalık sayı (float64) türündeki değeri bilimsel gösterimle temsil eder. %e küçük harflerle, %E büyük harflerle temsil eder.
	10. %s: Dize (string) türündeki değerleri temsil eder.
	11. %p: İşaretçi (pointer) türündeki değerleri temsil eder.
	12. %v: Verbose (detaylı) olarak, değişkenin değerini otomatik olarak türüne göre temsil eder.
	13. %%: Yüzde işareti (%) karakterini yazdırmak için kullanılır.
*/

func main() {
	// var name, age, isDead = "ErsaGun", 20, true

	//	-Print(args...), geriye yazdırılan byte sayısını ve hata varsa hatayı döndürür.
	/*
		printSize, printErr := fmt.Print(name, ", he is ", age, " years old, is he dead? ", isDead, "\n")
		fmt.Print("print size: ", printSize, ", print error: ", printErr, "\n")
	*/

	// -Printf(args...) geriye yazdırılan byte sayısını ve hata varsa hatayı döndürür.
	/*
		printSize, printErr := fmt.Printf("%s, he is %d years old, is he dead? %t\n", name, age, isDead)
		fmt.Printf("print size: %d print error: %s \n", printSize, printErr)
	*/

	// -Println(args...) geriye yazdırılan byte sayısını ve hata varsa hatayı döndürür.
	/*
		printSize, printErr := fmt.Println(name, ", he is ", age, " years old, is he dead? ", isDead)
		fmt.Println("print size:", printSize, "print error:", printErr)
		fmt.Println("fmt.Println finish")
	*/

	// -Errorf(format, args...): hata mesajı oluşturmak için kullanılır, geriye string ve hata döndürür.
	/*
		_, printErr := fmt.Println("ERROR MESSAGE")
		err := fmt.Errorf("error: %s", printErr)
		fmt.Println(err.Error())
	*/

	// -Sprint(args...): formatlanmış string oluşturmak için kullanılır, geriye string döndürür
	/*
		msg := fmt.Sprint(name, ", he is ", age, " years old, is he dead? ", isDead)
		fmt.Println(msg)
	*/

	// -Sprintf(format,args...): formatlanmış string oluşturmak için kullanılır, geriye string döndürür.
	/*
		message := fmt.Sprintf("%s, he is %d years old, is he dead? %t", name, age, isDead)
		fmt.Println(message)
	*/

	// -Sprintln(args...): formatlanmış string oluşturmak için kullanılır, geriye string döndürür.
	/*
		messge := fmt.Sprintln(name, ", he is ", age, " years old, is he dead? ", isDead)
		fmt.Println(messge)
	*/
	fmt.Println("fmt finish")
}
