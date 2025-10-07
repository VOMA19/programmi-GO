package sum

import "testing"

//file called sum_test.go
//to test the function sum in sum.go
//white box testing
//name of the test function must start with Test
//accept t *testing.T as parameter

func TestSum(t *testing.T) {
	//arrange
	a := 2
	b := 3

	//act
	result := sum(a, b)

	//assert
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
