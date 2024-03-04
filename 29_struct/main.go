package main

import (
	"fmt"
)

// type struct
// tüm paket içinde kullanılması için main dışında tanımlanır
type person struct {
	name     string
	age      int
	isMaried bool
}

func main() {
	// variable struct
	// var emp struct {
	// 	name     string
	// 	age      int
	// 	isMaried bool
	// }

	// fmt.Println(emp)
	// fmt.Printf("%#v \n", emp)

	// emp.age = 20
	// emp.name = "Ersagun"
	// emp.isMaried = false

	// fmt.Printf("%s %v %t \n", emp.name, emp.age, emp.isMaried)

	var emp person
	emp.name = "Ersagun"
	emp.age = 20
	emp.isMaried = false
	fmt.Printf("Name:%s, Age:%v, Maried:%t \n", emp.name, emp.age, emp.isMaried)

	var emp2 person
	emp2.name = "Fatih"
	emp2.age = 19
	emp2.isMaried = false
	fmt.Printf("Name:%s, Age:%v, Maried:%t \n", emp2.name, emp2.age, emp2.isMaried)

	// anonim struct

	boss := struct {
		name string
		age  int
	}{name: "Ersagun", age: 73}

	fmt.Println(boss)

	fmt.Println("Struct Finish")
}
