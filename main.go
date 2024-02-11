package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Welcome to %v booking system!", conferenceName)
	fmt.Println("Get your tickets here to attend ")
	fmt.Println("Total tickets available:", conferenceTickets)
	fmt.Println("Remaining tickets:", remainingTickets)
}
