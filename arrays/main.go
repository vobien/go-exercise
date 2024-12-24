package main

import (
	"fmt"
	"math"
)

type funcType func(int) float64

func main() {

	// urlMap := map[string] string {}
	// urlMap["java"] = "https://www.java.com"
	// urlMap["go"] = "https://www.go.com"

	// for k, v := range urlMap {
	// 	fmt.Println(k, v)
	// }

	// // allocate memory for slice
	// courses := make([]string, 3)
	// courses[0] = "go"
	// courses[1] = "python"
	// courses[2] = "java"

	// courses = append(courses, "rust")

	// fmt.Println(courses)
	// fmt.Println(len(courses), cap(courses))

	// // array
	// otherCourses := [3]string{"go", "python", "java"}
	// fmt.Println(otherCourses)
	// fmt.Println(len(otherCourses), cap(otherCourses))

	nums := []int{1, 2, 3, 4, 5}
	calculate(nums, double)
	calculate(nums, triple)

}

func calculate(nums []int, function funcType) {
	for _, num := range nums {
		fmt.Printf("%v ", function(num))
	}
	fmt.Println()
}

func double(x int) float64 {
	return float64(x * 2)
}

func triple(x int) float64 {
	return math.Pow(float64(x), float64(3))
}
