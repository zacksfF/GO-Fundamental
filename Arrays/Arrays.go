package main

import(
	"fmt"
)

type Room struct{
	name string
	cleaned bool
}
func checkCleanlines(rooms[4]Room){
	for i:=0; i < len(rooms); i++{
		room := rooms[i]
		if room.cleaned{
			fmt.Println(room.name, "is clean")
		}else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

func main(){
	rooms := [...]Room{
		{name:  "office"},
		{name: "warehouse"},
		{name: "reception"},
		{name: "Ops"},
	}

	checkCleanlines(rooms)
	fmt.Println("Performing cleaning ...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanlines(rooms)
}









