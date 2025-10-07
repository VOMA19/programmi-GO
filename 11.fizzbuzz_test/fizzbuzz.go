package fizzbuzztest

import "strconv"

func Fizzbuzz(n int) string {
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
