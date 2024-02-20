package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// GO'da errolar bir değişken olarak alınır.

	rslt, err := sRoot(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rslt)
	}

	fmt.Println("Error Handling Finish")
}

func sRoot(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Negative number") // errors içinde yeni bir hata oluşturduk
	}
	return int(math.Sqrt(float64(num))), nil // nil = null yani değer yok demek
}
