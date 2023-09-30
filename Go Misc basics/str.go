package main

import "fmt"

func main() {
	var A = 1 + 1
	fmt.Println("1 + 1 is =", A)

	fmt.Println("Hello world")
	fmt.Println("Hello Wold"[1])
	fmt.Println("Zakaria saif"[1])
	fmt.Println("Hello" + "YES ist's me")

	fmt.Println(true || true)   // or
	fmt.Println(false && false) //and
	fmt.Println(!false)         // not

	a := 2
	b := 3.5
	fmt.Printf("a,  %[1]v %T \n", a, a)
	fmt.Printf("b,  %[1]v %T \n", b, b)

	var i int = -129
	fmt.Printf(" %T\n", i)

	var kal rune = 'M'
	fmt.Printf("%T\n", kal)
	fmt.Println(kal)
}
