package main

import "fmt"

func main(){
	const premiumplan = "Premium Plan"
	const basicPalne ="BasicPlanName"

	fmt.Println("plan:", premiumplan)
	fmt.Println("plan:", basicPalne)

	const secondInMin = 60
	const minutesInhour = 60
	const HourInDay = secondInMin * minutesInhour * 24
	fmt.Println("In a sec in min we have:", secondInMin)
	fmt.Println("In a minute in Hour we have:", secondInMin)
	fmt.Println("In a day we have :", HourInDay)
}