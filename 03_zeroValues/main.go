package main

import "fmt"

func main() {
	/*
		var name string = "ErsaGun"
		var name2 = "Fikri"
		name3 := "Tosun"
	*/

	// var name, name2, name3 string = "ErsaGun", "Fikri", "Tosun"

	/*
		var name string = "ErsaGun"
		var age int = 20
		var isDead = true
	*/
	var (
		name string = "ErsaGun"
		age  int    = 20

		isDead = true
	)

	var (
		spaceTxt  string
		SpaceNum  int
		SpaceBool bool
	)

	fmt.Println("Zero")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isDead)

	fmt.Println(spaceTxt)
	fmt.Println(SpaceNum)
	fmt.Println(SpaceBool)
}
