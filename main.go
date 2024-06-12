package main

import "fmt"

func main() {
	var eventName = "AI Event"
	var eventDate = "2021-10-10"
	var eventTime = "10:00 AM"
	const eventTicket = 100
	var remainingTicket = 50

	fmt.Printf("eventTicket: %T\n", eventTicket)

	fmt.Printf("Welcome to the %v booking application forum\n", eventName)
	fmt.Printf("The %v is scheduled for %v at %v\n", eventName, eventDate, eventTime)
	fmt.Printf("The total number of tickets for the event is %v and %v are still available\n", eventTicket, remainingTicket)
	fmt.Println("We are happy to have you here!")

	var userName string
	// ask user to enter their name
	var userTickets int

	userName = "John Doe"
	userTickets = 5
	fmt.Println("Please enter your name: ")
	fmt.Scan(&userName)

	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

}
