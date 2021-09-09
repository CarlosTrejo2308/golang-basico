package main

import "fmt"

func main() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}