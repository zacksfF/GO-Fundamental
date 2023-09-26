package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Println("%v\n", numbers)
	fmt.Println("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a3)

	man := [5]bool{true}
	fmt.Printf("%#v\n", man)

	// ellipsis operator
	M5 := [...]int{1, 2, 3, 5, 4, -11}
	fmt.Printf("%#v\n", M5)
	fmt.Printf("the length of M5 is%d\n", len(M5))

	var M6 = [...]int{1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(M6)

}
