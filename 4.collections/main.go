package main

import "fmt"

func main() {

	var grades [5]int
	grades[0] = 85
	grades[1] = 90
	grades[2] = 78
	grades[3] = 92
	grades[4] = 88

	fmt.Printf("Grades in the array: %v\n", grades)

	for i := 0; i < 5; i++ {
		println("Item:", grades[i])
	}

	//Slices

	friends := []string{"Alice", "Bob", "Charlie"}
	friends = append(friends, "David", "Eve")
	fmt.Printf("Friends in the slice: %v\n", friends)

	for i := 0; i < len(friends); i++ {
		println("Friend:", friends[i])
	}

	for k, v := range friends {
		fmt.Printf("Index: %d: \t Friend: %s\n", k, v)
	}

	for _, v := range friends {
		fmt.Printf("Friend: %s\n", v)
	}

	//Maps

	studentGrades := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 78,
		"David":   92,
		"Eve":     88,
	}

	studentsClass := make(map[string]string)
	studentsClass["Alice"] = "3A"
	studentsClass["Bob"] = "3B"
	studentsClass["Charlie"] = "1C"
	studentsClass["David"] = "2A"
	studentsClass["Eve"] = "5B"

	for k, v := range studentsClass {
		fmt.Printf("Student: %s, Class: %s\n", k, v)
	}

	fmt.Printf("Student Grades: %v\n", studentGrades)
}
