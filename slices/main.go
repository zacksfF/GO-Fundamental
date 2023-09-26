package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("Cities is equal to nil :", cities == nil)
	fmt.Printf("Cities %#v\n", cities)
	fmt.Println(len(cities))

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	for index, value := range numbers {
		fmt.Printf("index: %v, value :%\n", index, value)
	}
}
