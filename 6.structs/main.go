package main

import "fmt"

//structs in go
//player scoped to this package
//Player can be accessed outside this package (e.g. repo. Player)

type player struct {
	name    string
	age     int
	country string
}

func main() {

	//initialize struct

	player1 := player{
		name:    "Federer",
		age:     38,
		country: "Switzerland",
	}

	//accessing struct fields
	fmt.Printf("Player 1: %v\n", player1)
	fmt.Println(player1.name)
	fmt.Println(player1.age)
	fmt.Println(player1.country)

}
