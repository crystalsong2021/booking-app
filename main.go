package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var cont string
	var email string
	var userTickets uint

	for {
		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Please enter email address")
		fmt.Scan(&email)

		fmt.Println("please enter the number of ticket you want to purchase")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v. You booked %v tickets. \n Your information will be be sent to %v \n. ", firstName, lastName, userTickets, email)
		fmt.Printf("The remaining ticket %v for %v\n", remainingTickets, conferenceName)
		fmt.Println("Do you want to continue.....")
		fmt.Scan(&cont)
		if cont == "no" {
			break
		}

	}
	firstNameBookings := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNameBookings = append(firstNameBookings, name[0])

	}
	fmt.Printf("Here is the list of totel booking %v\n\n", firstNameBookings)
	fmt.Println("****Good Bye***")

}
