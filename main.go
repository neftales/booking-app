package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
