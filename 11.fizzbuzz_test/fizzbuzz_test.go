package fizzbuzztest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz(t *testing.T) {
	t.Run("When1_then1", func(t *testing.T) {
		res := Fizzbuzz(1)
		assert.Equal(t, "1", res)
	})

	t.Run("When3_thenFizz", func(t *testing.T) {
		res := Fizzbuzz(3)
		assert.Equal(t, "Fizz", res)
	})

	t.Run("When5_thenBuzz", func(t *testing.T) {
		res := Fizzbuzz(5)
		assert.Equal(t, "BuzZ", res)
	})

	t.Run("When15_thenFizzBuzz", func(t *testing.T) {
		res := Fizzbuzz(15)
		assert.Equal(t, "FizBuzZ", res)
	})
}

/*
func TestFizzbuzz(t *testing.T) {
	//arrange
	var n int = 5
	var expected string = "FizBuzZ"
	//act
	var result string = Fizzbuzz(n)
	//assert
	if result != expected {
		t.Errorf("Fizzbuzz(%d) = %s; want %s", n, result, expected)
	}
}
*/
