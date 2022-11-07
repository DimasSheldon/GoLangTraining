package controller

import (
	"assignment1/model"
	"fmt"
	"strings"
)

func CreateAttendee(id, name, address, occupation, reason string) model.Attendee {
	return model.Attendee{
		ID:         id,
		Name:       name,
		Address:    address,
		Occupation: occupation,
		Reason:     reason}
}

func InitDefaultAttendees() []model.Attendee {
	attendees := []model.Attendee{
		CreateAttendee("GLNG033ONL1000", "Cantika Nanda", "Bandung", "Administrator", "Manage GoLang Class!"),
		CreateAttendee("GLNG033ONL2000", "Thomas Darmawan", "Tangerang", "Trainer", "Teach GoLang Class!"),
		CreateAttendee("GLNG033ONL1003", "Dhimas Adi", "Jakarta", "Trainee", "Who is GoLang!"),
		CreateAttendee("GLNG033ONL0009", "Stefanus Christanto", "Banten", "Trainee", "Where GoLang!"),
		CreateAttendee("GLNG033ONL1004", "Joshua Elias", "Dumai", "Trainee", "When GoLang!"),
		CreateAttendee("GLNG033ONL1002", "A. Faishal Fadhil", "Bekasi", "Trainee", "Learn GoLang!"),
		CreateAttendee("GLNG033ONL1001", "Evelline Krist.", "Depok", "Trainee", "What is GoLang!"),
		CreateAttendee("GLNG033ONL1008", "Bagus Kurnia", "Jawa Barat", "Trainee", "How GoLang!"),
		CreateAttendee("GLNG033ONL1007", "Fakhrul Muhammad Rijal", "Riau", "Trainee", "Belajar GoLang!"),
		CreateAttendee("GLNG033ONL1005", "Ricky Tri", "Bogor", "Trainee", "Learn GoLang!"),
		CreateAttendee("GLNG033ONL1006", "Dimas Sheldon", "Tangerang", "Trainee", "Why GoLang!"),
		CreateAttendee("GLNG033ONL1010", "Joshua Sumardi", "Kajarta", "Trainee", "What is GoLang!"),
	}
	return attendees
}

type AttendeeListModel model.AttendeeList

func (attendee *AttendeeListModel) AddAttendee(profile model.Attendee) []model.Attendee {
	attendee.Attendees = append(attendee.Attendees, profile)
	return attendee.Attendees
}

func PrintAttendee(attendee model.Attendee) {
	fmt.Println("---------------------------------------------")
	fmt.Println("ID\t\t: ", attendee.ID)
	fmt.Println("Name\t\t: ", attendee.Name)
	fmt.Println("Address\t\t: ", attendee.Address)
	fmt.Println("Occupation\t: ", attendee.Occupation)
	fmt.Println("Reason\t\t: ", attendee.Reason)
	fmt.Println("ID(index)\t: ", attendee.Index)
}

func ListAllAttendees(attendees []model.Attendee) {
	for idx, attendee := range attendees {
		attendee.Index = idx
		PrintAttendee(attendee)
	}
}

func FindAttendeeByName(attendees []model.Attendee, name string) {

	notFound := 0
	found := 0
	for idx, attendee := range attendees {
		// if strings.TrimRight(name, "\n") == attendee.Name { /* Exact Name */
		if strings.Contains(strings.ToLower(attendee.Name),
			strings.ToLower(strings.TrimRight(name, "\n"))) { /* Similar Name */
			attendee.Index = idx
			PrintAttendee(attendee)
			found++
		} else {
			notFound++
		}
	}

	if notFound == len(attendees) {
		fmt.Println("Not Found!")
	} else {
		fmt.Printf("Found [%d] record(s)\n", found)
	}
}

func FindAttendeeByID(attendees []model.Attendee, id string) {

	notFound := 0
	found := 0
	for idx, attendee := range attendees {
		// if strings.TrimRight(id, "\n") == attendee.Name { /* Exact Name */
		if strings.Contains(strings.ToLower(attendee.ID),
			strings.ToLower(strings.TrimRight(id, "\n"))) { /* Similar Name */
			attendee.Index = idx
			PrintAttendee(attendee)
			found++
		} else {
			notFound++
		}
	}

	if notFound == len(attendees) {
		fmt.Println("Not Found!")
	} else {
		fmt.Printf("Found [%d] record(s)\n", found)
	}
}

func FindAttendeeByAddressReason(attendees []model.Attendee, address string, reason string) {

	notFound := 0
	found := 0
	for idx, attendee := range attendees {
		// if strings.TrimRight(id, "\n") == attendee.Name { /* Exact x */
		if strings.Contains(strings.ToLower(attendee.Address), strings.ToLower(strings.TrimRight(address, "\n"))) &&
			strings.Contains(strings.ToLower(attendee.Reason), strings.ToLower(strings.TrimRight(reason, "\n"))) {
			attendee.Index = idx
			PrintAttendee(attendee)
			found++
		} else {
			notFound++
		}
	}

	if notFound == len(attendees) {
		fmt.Println("Not Found!")
	} else {
		fmt.Printf("Found [%d] record(s)\n", found)
	}
}
