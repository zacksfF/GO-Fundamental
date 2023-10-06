package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eegs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cercereal ?")
	if !found{
		fmt.Println("nope")
	}else{
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList{
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "on the shopping list")
}

