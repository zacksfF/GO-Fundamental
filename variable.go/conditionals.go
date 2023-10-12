package main

import "fmt"

func main(){
	messagelen := 10
	maxMsg := 20
	fmt.Println("trying to send a message of length:", messagelen)
	
	if messagelen < maxMsg {
		fmt.Println("Message sent")
	}else {
		fmt.Println("message not sent")
	}
}