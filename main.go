package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking system!\n", conferenceName)
	fmt.Println("Get your tickets here to attend ")
	fmt.Printf("Tickets remaining: %v | Tickets available: %v\n", remainingTickets, conferenceTickets)

	var userName string
	var userTickets int

	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	/* userName = "John Doe"
	userTickets = 5 */

	fmt.Printf("User: %v | Tickets: %v\n", userName, userTickets)
}
