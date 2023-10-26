package main

import "fmt"

//weapon type
//axe
//sword
//wooden stick
//knife
type weoponType int 

const (
	axe weoponType = 1
	sword weoponType  = 2
	woodenStick weoponType = 3
	knife weoponType = 4
)

func getDamage(weaponTyope weoponType) int {
	switch weaponTyope {
	case axe:
		return 100
	case sword:
		return 90
	case woodenStick:
		return 1
	case knife:
		return 40
	default:
		panic("weapon does not exist")
	}
}
func main (){
	fmt.Println("damage of weopon :", getDamage(sword))
}