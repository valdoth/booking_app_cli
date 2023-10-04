package main

import (
	"booking-app/helper"
	"fmt"
	// "strconv"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	const conferenceTickets int = 50
	var conferenceName string = "Go conference"
	var remainingTickets uint = 50
	// var bookings = make([]map[string]string, 0)
	var bookings = make([]UserData, 0)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidTicketNumber || !isValidEmail || !isValidName {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you enterd doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println()
			continue
		}

		bookTicket(&bookings, &remainingTickets, userTickets, firstName, lastName, email, conferenceName)

		sendTicket(userTickets, firstName, lastName, email)

		allFirstNames := getFirstNames(bookings)
		fmt.Printf("These are all our bookings: %v\n", allFirstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
		fmt.Println()
	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	fmt.Println()
}

// func getFirstNames(bookings []map[string]string) []string {
func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// func bookTicket(bookings *[]map[string]string, remainingTickets *uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) {
func bookTicket(bookings *[]UserData, remainingTickets *uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	*remainingTickets -= userTickets

	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lasttName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	*bookings = append(*bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", *remainingTickets, conferenceName)
	fmt.Printf("List of bookings is %v\n", *bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################################")
	fmt.Printf("Sending ticket: %v to email address %v\n", ticket, email)
	fmt.Println("####################################")
}