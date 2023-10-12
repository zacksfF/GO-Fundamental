package main

import "fmt"

func main(){
	firstanme, _ := getNames()
	fmt.Println("Welcome to Textio,", firstanme)
}
func getNames()(string, string) {
	return "john", " zack"
}