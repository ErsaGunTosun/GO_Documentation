package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 1- iki rakam arasında toplama,çıkarma ve çarpma işlemleri yapacak bir fonksiyon yazın.
	// -> fmt.Println(numsOperation(10, 5))

	// 2- kullanıcı tarafından girilen sayıya göre geçti kaldı
	// -> passFail()

	// 3 - 1 ile 100 arasında bir sayı tahmin oyunu yapın maksimum tahmin hakkı 10 olsun
	// -> guessGame()

	fmt.Println("App Finish")
}

func numsOperation(x, y int) (int, int, int) {
	sum := x + y
	sub := x - y
	mul := x * y

	return sum, sub, mul
}

func passFail() {
	fmt.Print("Enter your first grade : ")
	reader := bufio.NewReader(os.Stdin)
	grade1, _ := reader.ReadString('\n')
	grade1 = strings.TrimSpace(grade1) // girilen değerlerin başında ve sonunda boşluk olmaması için
	num1, _ := strconv.Atoi(grade1)    // string olan değeri int'e çeviriyoruz

	fmt.Print("Enter your second grade : ")
	grade2, _ := reader.ReadString('\n')
	grade2 = strings.TrimSpace(grade2) // girilen değerlerin başında ve sonunda boşluk olmaması için
	num2, _ := strconv.Atoi(grade2)    // string olan değeri int'e çeviriyoruz

	result := (float64(num1) * 0.4) + (float64(num2) * 0.6)

	if result >= 50 {
		fmt.Println("Passed")
		return
	}

	fmt.Println("Failed")
}

func guessGame() {
	random := rand.Intn(99) + 1
	guessCount := 10
	for guessCount > 0 {
		fmt.Print("Enter your guess : ")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess) // girilen değerlerin başında ve sonunda boşluk olmaması için
		num, _ := strconv.Atoi(guess)    // string olan değeri int'e çeviriyoruz
		if num == random {
			fmt.Println("Congratulations! You found the number")
			return
		}
		if num > random {
			fmt.Println("Down")
			fmt.Println("Remaining right : ", guessCount-1)
		} else {
			fmt.Println("Up")
			fmt.Println("Remaining right : ", guessCount-1)
		}
		guessCount--
	}

	if guessCount == 0 {
		fmt.Println("You lost! The number was ", random)

	}
}
