//os.Args is a slice (dynamic array) in the Go programming language that provides access to the command-line arguments passed to a Go program. It is part of the os package and is commonly used when you want to work with command-line arguments provided to your Go application.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args", os.Args)
}
