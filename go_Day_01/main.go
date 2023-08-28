package main

import (
	"fmt"
	"strings"
)

func main() {

		var conferenceName string = "Go Conference"
		const conferenceTickets int = 50
		var remainingTickets uint = 50
		var bookings [] string

	for {
		
		fmt.Printf("Welcome to %v booking application!\n", conferenceName)
		fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")

		//Ask for Names, Email and Tickets

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

		remainingTickets = remainingTickets - uint(userTickets)
		bookings = append(bookings, firstName+" "+lastName)

		firstNames := [] string {} //var firstNames [] string
		for _, booking:= range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames,names[0])
		}

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail to %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v!\n", remainingTickets, conferenceName)
		// fmt.Printf("These are all the bookings we have %v\n", bookings)

		fmt.Printf("The first names of the bookings : %v\n", firstNames)
	}
}
