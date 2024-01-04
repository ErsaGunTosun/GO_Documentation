package main

import "fmt"

/*	1-) studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini
		3 farklı yöntem ile oluşturup, çıktısını yazdırınız.✔
	2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.✔
	3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)✔
	4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO üzerinden gösteriniz.
	5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration
*/

func main() {
	// 1 -
	/*
			var studentName string = "John Doe"
			var grade int = 77
			var isPassed bool = true

			var (
				studentName string = "John Doe"
				grade       int    = 77
				isPassed    bool   = true
			)


		studentName := "John Doe"
		grade := 77
		isPassed := true
	*/

	// 2-
	/*
		var studentName, grade, isPassed = "John Doe", 77, true
	*/

	// 3-
	/*
		- Declaaration: Bir değişkeni türü ve adı ile tanımlama işlemidir
			var number int, var name string gibi
		- Assign: Tanımlanan bir değişkene değer değiştirme işlemidir (İlk değer ataması assign olarak geçmez)
			var number int = 10, number= 5 gibi
		- Initialization: Bir değişkene ilk değer atama işlemidir
			var y int = 1 veya var x int, x = 2 gibi
		-Initial Value: Bir değişkene ilk değer ataması yapılırken, eğer değer atanmazsa, o değişkenin varsayılan değeridir
			var x int, x'in initial value'su 0'dır
	*/

	// 4-
	/*
		- Statically Typed: Bir programn çalıştırmadan önce değişkenlerin türlerinin belirlenmesidir. Kullanılan değişkenlerin
		türleri açıkca belirtilir ve bu sayede bir çok hata önlenmiş olur.
		- Dynamically Typed: Bir program çalıştırıldıktan sonra değişkenlerin türlerinin belirlenmesidir.

		-> Statically Typed
			var name string = "ErsaGun"
			var age int = 20
			var isDead bool = false
		-> Dynamically Typed
			name := "ErsaGun"
			age := 20
			isDead := false
	*/

	// 5-
	/*
		-"=" atama işlemidir. Bir değişkene sadece değer ataması yapar.
		-":=" operatörü bir değişkene tür ve değer atamasını aynı anda yapar, sadece local değişkenlerin
		atanması ve tanımlanmasında kullanılır
	*/

	fmt.Println("Finish")
}
