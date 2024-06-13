package main

import (
	"fmt"
	"strings"
)

// package level variables
const eventTicket = 100

var eventName = "AI Event"
var eventDate = "2021-10-10"
var eventTime = "10:00 AM"
var remainingTicket = 50
var bookings []string

func main() {

	fmt.Printf("eventTicket: %T\n", eventTicket)

	// function to greet users and display event details
	greetUsers()

	for {
		// function to get user input
		firstName, lastName, email, phoneNumber, userTickets := getUserInput()

		// function to validate user input
		isValidName, isValidEmail, isValidPhoneNumber, isValidUserTickets := validateUserInput(firstName, lastName, email, phoneNumber, userTickets)

		if isValidName && isValidEmail && isValidPhoneNumber && isValidUserTickets {
			//
			remainingTicket, bookings = bookTickets(userTickets, firstName, lastName, email, phoneNumber)

			//  funtion to list all bookings for the event by their first names
			firstNames := getFirstNames()
			fmt.Printf("The list of all bookings for the %v event are: %v\n", eventName, firstNames)

			// check if all tickets have been booked(can it be a else if statement?)
			if remainingTicket == 0 {
				fmt.Println("All tickets have been booked for the event. Please come back next year")
				break
			}

		} else {
			// error messages to be displayed if user input is invalid
			if !isValidName {
				fmt.Println("Your first name and last name should be more than 2 characters")
			}
			if !isValidEmail {
				fmt.Println("Your email address should contian '@' and '.'")
			}
			if !isValidPhoneNumber {
				fmt.Printf("Your phone number should be 11 characters long\n")
			}
			if !isValidUserTickets {
				fmt.Println("You can only book between 1 and 50 tickets")
			}

		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application forum\n", eventName)
	fmt.Printf("The %v is scheduled for %v at %v\n", eventName, eventDate, eventTime)
	fmt.Printf("The total number of tickets for the event is %v and %v are still available\n", eventTicket, remainingTicket)
	fmt.Println("We are happy to have you here!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, phoneNumber string, userTickets int) (bool, bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidPhoneNumber := len(phoneNumber) == 11
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidPhoneNumber, isValidUserTickets
}

func getUserInput() (string, string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var phoneNumber string
	var userTickets int

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

	return firstName, lastName, email, phoneNumber, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string, phoneNumber string) (int, []string) {
	remainingTicket = remainingTicket - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets for the %v event\n", firstName, lastName, userTickets, eventName)
	fmt.Printf("Your ticket(s) will be sent to your email: %v and phone number: %v\n", email, phoneNumber)
	fmt.Printf("There are %v tickets remaining for the %v event\n", remainingTicket, eventName)
	fmt.Println("We are happy to have you here!")

	return remainingTicket, bookings
}
