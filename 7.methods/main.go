package main

import (
	"errors"
	"fmt"
	"methods/repo"
	"strings"
)

type player struct {
	name    string
	age     int
	country string
}

// function to print player info
func printPlayerInfo(input player) {
	fmt.Println("Name:", input.name)
	fmt.Println("Age:", input.age)
	fmt.Println("Country:", input.country)
}

// function with return method
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// method to print player info
//value receiver type 	  => func (p player) printInfo(){}
//reference receiver type => func (p *player) printInfo(){}

func (p player) printInfo() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Country:", p.country)
}

func (p *player) printInfoPtr() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Country:", p.country)
}

func main() {
	//methods and functions in go

	player1 := player{
		name:    "Federer",
		age:     38,
		country: "Switzerland",
	}

	car1 := repo.Car{
		Brand: "AUDI",
		Model: "rs3",
		Year:  2025,
	}

	printPlayerInfo(player1) // calling function
	fmt.Println(strings.Repeat("#", 20))
	player1.printInfoPtr() // calling method with pointer receiver
	fmt.Println(strings.Repeat("#", 20))
	player1.printInfo() // calling method with value receiver

	fmt.Println(strings.Repeat("#", 20))
	fmt.Printf("Car details: %s %s %d\n", car1.Brand, car1.Model, car1.Year)

	var a, b int = 10, 2
	result, err := divide(a, b) // calling function with return type
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of %d / %d = %d\n", a, b, result)
	}
}
