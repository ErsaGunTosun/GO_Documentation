// Package clause(Paket Bildirimi)
package main

/*
Kodları paketler haline getirerek tekrar kullanılabilirliği arttırırız.
Burada 'main' verme sebebimiz bu kodu çalıştırmak için kullanacağımız paketin adının main olması gerektiğidir.
*/

// Import statement (import ifadesi)
import "fmt"

/*
GO standart paketleri veya başka paketleri kullanmak için import ifadesini kullanırız.
*/

// My Code (Kodlarımız)
func main() {

	fmt.Println("I'm computer")
}

/*
- go run
- go build = compile packages and dependencies (bağımlılıkları ve paketleri derler)
*/
