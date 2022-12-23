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

	greetUser( conferenceName, conferenceTickets, remainingTickets)
	var cont string


	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("Invalid Name")
		}

		if !isValidEmail {
			fmt.Println("Invalid Email")
		}

		if !isValidTicket {
			fmt.Println("Invalid Ticket")
		}
		//if there user booked more ticket than available ticket, print error message
		if userTickets > remainingTickets {
			fmt.Printf("Sorry, There is only %v remainingTicket left. Please try again.\n", remainingTickets)
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v. You booked %v tickets. \n Your information will be be sent to %v \n. ", firstName, lastName, userTickets, email)
		fmt.Printf("The remaining ticket %v for %v\n", remainingTickets, conferenceName)
		fmt.Println("Do you want to continue.....")
		fmt.Scan(&cont)
		if cont == "no" {
			break
		}
		//if there is no more remaining ticket, then quit the program
		if remainingTickets == 0 {
			fmt.Println("There is no more ticket left!!!!")
			break
		}

	}
	printName := getFirstName(bookings)
	//print out only the firstName from the slice
	fmt.Printf("Here is the list of totel booking %v\n\n", printName)
	fmt.Println("****Good Bye***")

}

//writing a function
func greetUser(conferenceName string, conferenceTickets uint, remainingTickets uint ) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstName(bookings []string) []string {
	firstNameBookings := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking) //split the string with space ["crystal", "song"]
		//only save the first name in the slice(go term for array)
		firstNameBookings = append(firstNameBookings, name[0])

	}
	return firstNameBookings
}

//in Go, you can return multiple value
func validateUserInput (firstName string, lastName string, email string, userTickets uint, remainingTickets uint ) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets > remainingTickets

	return isValidName, isValidEmail, isValidTicket;

}

//get user input
func getUserInput () (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter email address")
	fmt.Scan(&email)

	fmt.Println("please enter the number of ticket you want to purchase")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}
