package main

import (
	"fmt"
	"strings"
)

func main() {
	// Pointers in Go
	var intVal int = 34
	fmt.Println("Value of intVal:", intVal) // prints 34
	//address of intVal using &
	fmt.Println("Address of intVal:", &intVal) // prints address of intVal

	//pointer variable
	var ptr *int = &intVal // pointer variable ptr holds the address of intVal
	fmt.Println("Value of ptr (address of intVal):", ptr)
	//dereferencing ptr to get value of intVal using * --> dereferencing operator
	fmt.Println("Value at address stored in ptr:", *ptr) // dereferencing ptr to get value of intVal

	//changing value of intVal using pointer

	intVal = 45
	fmt.Println(strings.Repeat("#", 60))
	fmt.Println("Value of intVal:", intVal)
	fmt.Println("Address of intVal:", &intVal)
	fmt.Println("Value of ptr (address of intVal):", ptr)
	fmt.Println("Value at address stored in ptr:", *ptr)

}
