package main

import "fmt"

func main() {

	var car string = "Audi"         // Declare a variable named car of type string and assign it the value "Audi"
	fmt.Println("The car is:", car) // Print the value of the car variable to the console

	cheapcar := "Toyota"                       // Declare a variable named cheapcar using short variable declaration and assign it the value "Toyota"
	fmt.Println("The cheap car is:", cheapcar) // Print the value of the cheapcar

	var myage int = 19    // Declare a variable named myage of type int and assign it the value 19
	var majorage int = 18 // Declare another variable named majorage of type int and assign it the value 18

	isAdult := myage >= majorage // Use a short variable declaration to check if myage is greater than or equal to majorage

	if isAdult { // Check if myage is greater than or equal to majorage
		fmt.Println("You are an adult") // Print that the user is an adult
	} else {
		fmt.Println("You are not an adult") // Print that the user is not an adult
	}
}

// This program demonstrates variable declaration and printing in Go.
