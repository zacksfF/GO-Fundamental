package main

import "fmt"

const (
	retry1 = "click here to sign up"
	retry2 = "pretty please click here"
	retry3 = "we beg you to sign up"
)

func getMessageWuthRetries() [3]string {
	return [3]string{
		retry1,
		retry2,
		retry3,
	}
}

func send(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	message := getMessageWuthRetries()
	for i := 0; i < len(message); i++ {
		msg := message[i]
		fmt.Printf(`sending: %v `, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they respond !")
			break
		}
		if i == len(message)-1 {
			fmt.Println("complete failure")
		}
	}
}

func main() {
	send("Bob", 0)
	send("Alice", 1)
	send("Mangalam", 2)
	send("Ozgur", 3)
}
