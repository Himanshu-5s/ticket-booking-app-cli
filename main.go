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

	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var phoneNumber string
		var userTickets int

		//userName = "John Doe"
		//userTickets = 5
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Please enter your phone number: ")
		fmt.Scan(&phoneNumber)
		fmt.Println("Please enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		// logic for updating the remaining tickets
		remainingTicket = remainingTicket - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets for the %v event\n", firstName, lastName, userTickets, eventName)
		fmt.Printf("Your ticket(s) will be sent to your email: %v and phone number: %v\n", email, phoneNumber)
		fmt.Printf("There are %v tickets remaining for the %v event\n", remainingTicket, eventName)
		fmt.Println("We are happy to have you here!")
		fmt.Printf("The list of all bookings for the %v event are: %v\n", eventName, bookings)
	}
}
