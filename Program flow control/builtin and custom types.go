package main

import "fmt"

var (
	floatVar   float32 = 0.1
	floatVar64 float64 = 0.1
	name       string  = "FOO"
	intVar32   int32   = 1
	intvar64   int64   = 3536374
	intVAr     int     = 10
	uintvar    uint    = 1
	uinvart32  uint32  = 1
	uintvar64  uint64  = 34
	uintvar8   uint8   = 0x1
	bytevar    byte    = 0x2
	runvar     rune    = 'a'
)

type player struct {
	name        string
	health      int
	attackpower float64
}

//func main (){
	//users := map[string]int{
		//"zack": 21
	//users[varf]  := 24 
	//}
//}
 

func Gethealth(player player) int{
	return player.health
}

func mmain() {
	player := player{
		name:        "capitaine zack",
		health:      100,
		attackpower: 45.1,
	}
	fmt.Printf("This is the player health : %d\n", Gethealth(player))
}

func main(){
	users := make(map[string]int)
	//
	users ["foo"]= 10
	users["bar"] = 11

	age, ok := users["baz"]
	if !ok {
		fmt.Println("baz not exist in the map")
	}else {
		fmt.Println("baz exist in the map:", age)
	}

	fmt.Println("age : %D\n", age)
}