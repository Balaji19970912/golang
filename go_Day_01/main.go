package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	for {

		greetUser(conferenceName, conferenceTickets, remainingTickets) // greeting user function call
		
		var toExit string

		fmt.Println("PRESS 1 to EXIT and PRESS 0 to CONTINUE")
		fmt.Scan(&toExit)

		isValidExit := len(toExit) == 1 && toExit == "1"
		isValidContinue := len(toExit) == 1 && toExit == "0"

		if isValidExit {
			fmt.Println("Thank you! Go Conference")
			fmt.Println("----------****-----------")
			break
		} else if isValidContinue {
			
			firstName, lastName, email, userTickets := getUserInput() // calling user input function
			isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets) // validation of user input
			
			if !isValidName {
				fmt.Println("More than 2 letter has to be given for firstname or lastname!")
				break
			} else if !isValidEmail {
				fmt.Println("More than 2 letter has to be given for firstname or lastname!")
				break
			} else if !isValidTicketNumber {
				fmt.Println("Kindly enter positive value for the tickets or tickets soldout!")
				break
			}

			if isValidName && isValidEmail && isValidTicketNumber {
				//a function will come below
				bookings, remainingTickets = bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

				// fmt.Printf("These are all the bookings we have %v\n", bookings)

				//call firstname function
				firstNames := getFirstNames(bookings)

				fmt.Printf("The first names of the bookings : %v\n", firstNames)

				// var noTicketsRemaining bool = remainingTickets == 0  //another method of declaring
				noTicketsRemaining := remainingTickets == 0

				if noTicketsRemaining {
					fmt.Println("All tickets are sold out for our conference! Come back next year!")
					break
				}
			} else {
				fmt.Println("You have given an invalid input! Try again")
				break
			}
		} else {
			fmt.Println("Invalid Input for exit or continue!")
			break
		}
	}
}

func greetUser(confName string, confTickets int, remainTickets uint) {
	fmt.Printf("------***------\n")
	fmt.Printf("Welcome to the %v\n", confName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available!\n", confTickets, remainTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email Id:")
	fmt.Scan(&email)
	fmt.Println("Enter Number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{} //var firstNames [] string
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := (userTickets > 0) && (userTickets <= int(remainingTickets))
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(remainingTickets uint, userTickets int, bookings []string, firstName string, lastName string, email string, conferenceName string) ([]string, uint) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v!\n", remainingTickets, conferenceName)

	return bookings, remainingTickets
}