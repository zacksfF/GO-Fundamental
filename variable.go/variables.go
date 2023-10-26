package main

import "fmt"

var x = 10

//var (
//	name             = "FOO"
//	firstname string = "FOO"
//	lastName  string
//)
var name = "FOO"
var firsttname string = "FOO"

const (
	version = 1
	KeyLen = 10
)

func main() {

	//var version int
	//version = 10
	//fmt.Println(name, firstname, lastName)
	//version := 1 // infer ibt
	//otherversion :=  "BAR" // string
	//anoter_version := 10.1 //float
	fmt.Println(version, KeyLen)
}
