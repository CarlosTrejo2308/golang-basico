package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola")
	normalFunction("Hola")
	normalFunction("Hola")

	tripleArgument(1, 2, "Hola")

	value := returnValue(3)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println("1:", value1, "\n2:", value2)

}
