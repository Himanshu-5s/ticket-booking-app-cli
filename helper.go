package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, phoneNumber string, userTickets int) (bool, bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidPhoneNumber := len(phoneNumber) == 11
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidPhoneNumber, isValidUserTickets
}
