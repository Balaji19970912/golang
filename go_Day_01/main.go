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

		fmt.Printf("Welcome to %v booking application!\n", conferenceName)
		fmt.Printf("We have a total of %v tickets and %v tickets are still available!\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend!")

		//Ask for Names, Email and Tickets

		var firstName string
		var lastName string
		var email string
		var userTickets int
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

			fmt.Println("Enter your First Name:")
			fmt.Scan(&firstName)
			fmt.Println("Enter your Last Name:")
			fmt.Scan(&lastName)
			fmt.Println("Enter your Email Id:")
			fmt.Scan(&email)
			fmt.Println("Enter Number of tickets:")
			fmt.Scan(&userTickets)

			isValidName := len(firstName) >= 2 && len(lastName) >= 2
			isValidEmail := strings.Contains(email, "@")
			isValidTicketNumber := (userTickets > 0) && (userTickets <= int(remainingTickets))

			if isValidName && isValidEmail && isValidTicketNumber {
				remainingTickets = remainingTickets - uint(userTickets)
				bookings = append(bookings, firstName+" "+lastName)

				firstNames := []string{} //var firstNames [] string
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])
				}

				// fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail to %v\n", firstName, lastName, userTickets, email)
				fmt.Printf("%v tickets are remaining for %v!\n", remainingTickets, conferenceName)
				// fmt.Printf("These are all the bookings we have %v\n", bookings)

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
