package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	for i, v := range numbers {
		fmt.Println("index:", i, "value", v)
	}
	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index :", i, "value :", numbers[i])
	}

}
