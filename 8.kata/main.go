package main

import (
	"fmt"
	"strconv"
)

func fitzbuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizBuzZ"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "BuzZ"
	} else {
		return strconv.Itoa(n)
	}
}

func main() {
	num := 0
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)
	result := fitzbuzz(num)
	fmt.Println(result)

	for i := 1; i <= 15; i++ {
		fmt.Printf("%d: %s\n", i, fitzbuzz(i))
	}
}
