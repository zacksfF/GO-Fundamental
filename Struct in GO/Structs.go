package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}
type Bus struct {
	FrontSeat Passenger
}

func main() {
	caesy := Passenger{"Casey", 1, false}
	fmt.Println(caesy)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Hiedi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	caesy.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("bill has boarded the bu ")
	}
	if caesy.Boarded {
		fmt.Println(caesy.Name, "has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
