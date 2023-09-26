package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)         

	// slices with the same element type can be assigned to each other
	var n []int
	n = numbers
	_ = n

	//** COMPARING SLICES **//


	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true
	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
