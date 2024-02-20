package main

import "fmt"

func main() {

	// for <initialization statement>; <condition>; <post statement> {<code>}

	for i := 0; i <= 5; i += 2 {
		fmt.Println(i)
	}

	// infinite loop
	/*for {
		fmt.Print("A")
	}*/

	fmt.Println("-----------------")

	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("-----------------")

	for i := 0; i <= 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("For Loop Finish")
}
