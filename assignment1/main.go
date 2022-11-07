package main

import (
	"assignment1/controller"
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(bufio.NewReader(os.Stdin))
var attendees = controller.InitDefaultAttendees()

func main() {
	in := 1
	for in != 0 {

		fmt.Println("\t===============================")
		fmt.Println("\t===GLNG 33 ONL ATTENDEE MENU===")
		fmt.Println("\t===============================")
		fmt.Println("\t  1. Show All Attendees")
		fmt.Println("\t  2. Show Attendee Information by Name")
		fmt.Println("\t  3. Show Attendee Information by ID")
		fmt.Println("\t  4. Show Attendee Information by Address & Reason")
		fmt.Println("\t  5. Update Attendee Information")
		fmt.Println("\t  6. Add New Attendee ")
		fmt.Println("\t  7. Remove Attendee ")
		fmt.Println("\t  0. Exit")
		fmt.Println("")
		fmt.Print("\t Choose menu: ")

		var menu int
		_, err := fmt.Scanf("%d", &menu)

		if err == nil {
			switch menu {
			case 0:
				in = 0
				fmt.Println("ByeBye Thank you!")
			case 1:
				menu1()
				menuBackToMain()
			case 2:
				menu2()
				menuBackToMain()
			case 3:
				menu3()
				menuBackToMain()
			case 4:
				menu4()
				menuBackToMain()
			case 5:
				menu5()
				menuBackToMain()
			case 6:
				menu6()
				menuBackToMain()
			case 7:
				menu7()
				menuBackToMain()
			default:
				fmt.Println("Menu not listed!")
				menuBackToMain()
			}
		} else {
			fmt.Printf("Wrong input! (%s)\n", err)
		}
	}
}

func menuBackToMain() {
	scanner.Scan() // to ignore unwanted \n
	fmt.Println("--Press enter to back to main menu--")
	scanner.Scan() // to ignore unwanted \n
}

func menu1() {
	controller.ListAllAttendees(attendees)
}

func menu2() {
	scanner.Scan() // to ignore first \n

	fmt.Print("Input Attendee's Name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Finding \"%s\" Profile(s)...\n", name)

	controller.FindAttendeeByName(attendees, name)
}

func menu3() {
	scanner.Scan() // to ignore first \n

	fmt.Print("Input Attendee's ID: ")
	scanner.Scan()
	id := scanner.Text()
	fmt.Printf("Finding \"%s\" Profile(s)...\n", id)

	controller.FindAttendeeByID(attendees, id)
}

func menu4() {
	scanner.Scan() // to ignore first \n

	fmt.Print("Input Attendee's Address: ")
	scanner.Scan()
	address := scanner.Text()

	fmt.Print("Input Attendee's Reason: ")
	scanner.Scan()
	reason := scanner.Text()

	fmt.Printf("Finding Profile(s) based on Address[%s] & Reason[%s]...\n", address, reason)

	controller.FindAttendeeByAddressReason(attendees, address, reason)
}

func menu5() { fmt.Println("Menu under construction!") }
func menu6() { fmt.Println("Menu under construction!") }
func menu7() { fmt.Println("Menu under construction!") }
