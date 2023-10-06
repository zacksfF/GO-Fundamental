package main

import "fmt"

func main() {
	url := "example.com"
	switch url {
	case "example.com":
		fmt.Println("test")
	case "google.com":
		fmt.Println("live")
	default:
		fmt.Println("dev")
	}

	x := 3
	switch x {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("other:", x)
	}
}
