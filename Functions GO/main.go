package main

import "fmt"

func concat(s1, s2 string) string {
    return s1 + s2
}


func main() {
    test("Me,", " happy birthday!")
    test("zack,", " out of the context")
    test("Go& rust & python", " is fantastic")
}

func test(s1 string, s2 string) {
    fmt.Println(concat(s1, s2))
}
