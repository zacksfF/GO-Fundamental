package main

import "fmt"

func main() {
	price, instock := 100, true
	if price >= 80 {
		fmt.Println("To expensive !")
	}
	if price <= 100 && instock == true {
		fmt.Println("Buy it!")
	}
	if price < 100 {
		fmt.Println("it's cheap !")
	} else if price != 100 {
		fmt.Println("It's expensive")
	} else {
		fmt.Println("on the edge")
	}
}
