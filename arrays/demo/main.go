package main

import (
	"fmt"
)

type Product struct {
	title string
	id string
	price float64
}

func main() {
	hobbies := [3]string{"coding", "reading", "fucking"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	slices := hobbies[0:2]
	fmt.Println(slices)

	slices = hobbies[1:]
	fmt.Println(slices)

	courses := []string{"go", "python", "java"}
	fmt.Println(courses)

	courses[1] = "rust"
	fmt.Println(courses)

	courses = append(courses, "javascript")
	fmt.Println(courses)


	products := []Product {
		Product{"linux", "1", 100.0},
		Product{"go", "2", 200.0},
	}
	fmt.Println(products)

	products = append(products, Product{"rust", "3", 300.0})
	fmt.Println(products)
}
