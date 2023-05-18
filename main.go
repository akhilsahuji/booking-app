package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go Conference 2023"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTicket := getUserInputs()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(userTicket, firstName, lastName, email)

		wg.Add(1)
		go sendEmail(userTicket, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our tickets are sold out. Come back next year. ")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v BookingApp!!\n", conferenceName)
	fmt.Println("Get your ticket now! ")
	fmt.Printf("We have total of %v tickets and %v tickets are still left.\n", conferenceTickets, remainingTickets)

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Please enter your name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Please enter number of ticket you want:")
	fmt.Scanln(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTickets(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTicket)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	// //create a map
	// user["firstName"] = firstName
	// user["lastName"] = lastName
	// user["email"] = email
	// user["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("User %v booked %v tickets, and its confirmation is sent to %v.\n", firstName, userTicket, email)
	fmt.Printf("%v tickets remaining of %v\n", remainingTickets, conferenceName)

}

func sendEmail(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Printf("##################")
	wg.Done()

}
