package main

import "fmt"

func parOImpar(num int) string {
	switch modulo := num % 2; modulo {
	case 0:
		return "Es par"
	default:
		return "Es impar"
	}
}

func main() {
	valor1 := 1

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	queEs := parOImpar(24)
	fmt.Println(queEs)

	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}
