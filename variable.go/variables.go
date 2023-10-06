package main

import "fmt"

func main() {
	var mynames = "zakaria"
	fmt.Println("my name is :", mynames)

	var name string = "kathy"
	fmt.Println("name = ", name)

	username := "admin"
	fmt.Println("username= ", username)

	var sum int
	fmt.Println("The sum is ", sum)

	part1, other := 1, 5
	fmt.Println("part1 is ", part1, "other is ", other)

	part2, other2 := 1, 5
	fmt.Println("part2 is ", part2, "other2 is ", other2)

	sum = part1 + part2
	fmt.Println("sum = ", sum)

	var (
		textName = "variables"
		textType = "demo"
	)
	fmt.Println("textName is", textName)
	fmt.Println("textType is", textType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}
