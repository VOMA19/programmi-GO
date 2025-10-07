package main

import (
	"fmt"
	"sort"
	"strings"
)

func decimalToRoman(num int) (roman string, err error) {

	if num < 1 || num > 3999 {
		return "", fmt.Errorf("numero fuori dal range (1-3999)")
	}

	romanNums := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	romanKeys := make([]int, 0)
	for k := range romanNums {
		romanKeys = append(romanKeys, k)
	}

	sort.Ints(romanKeys)

	for num > 0 {
		for i := 0; i < len(romanKeys); i++ {
			if i == len(romanKeys)-1 || (num >= romanKeys[i]) && num < romanKeys[i+1] {
				num -= romanKeys[i]
				roman += romanNums[romanKeys[i]]
				break
			}
		}
	}

	return roman, nil
}

func main() {
	var v int
	values := []int{37, 4, 9, 58, 1994, 2024, 3999, 0, -5, 4000}
	for _, v = range values {
		roman, err := decimalToRoman(v)

		if err != nil { 
			fmt.Println("il numero ", v, "è un ", err)
		} else {
			fmt.Println("il numero ", v, " convertito in numero romano è: ", roman)
		}
	}

	//continue vs break

	//continue salta l'iterazione corrente
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue //salta il 3
		}
		fmt.Println(i)
	}

	fmt.Println(strings.Repeat("#", 20))
	//break esce dal ciclo
	for i := 0; i < 10; i++ {
		if i == 6 {
			break //esce dal ciclo al 3
		}
		fmt.Println(i)
	}

	var input int
	fmt.Print("Inserisci un numero da convertire in romano: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Errore nella lettura del numero:", err)
	} else {
		roman, err := decimalToRoman(input)
		if err != nil {
			fmt.Println("il numero", input, "è un", err)
		} else {
			fmt.Println("il numero", input, "convertito in numero romano è:", roman)
		}
	}

}
