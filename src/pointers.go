package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

// Creat funcion a una estructura
func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}


// Cuando imprimamos el objeto
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	a := 50
	b := &a

	fmt.Println(a, &a, b, *b)

	*b = 100
	fmt.Println(a)


	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)
	myPc.ping()

	myPc.duplicateRAM()
	fmt.Println(myPc)
	
	myPc.duplicateRAM()
	fmt.Println(myPc)
	fmt.Printf("2, %T", 2)
}