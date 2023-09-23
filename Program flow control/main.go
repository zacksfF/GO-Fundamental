package main

import "fmt"

func main() {
	price, instock := 1000, true
	if price > 800 {
		fmt.Println("Too expensive!")
	}
	_ = instock
}
