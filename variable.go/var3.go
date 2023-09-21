// Declaration with type inference

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = 10
	var s = "france"
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}
