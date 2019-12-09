package main

import (
	"fmt"
	"github.com/cynic89/hackday/hobbybuddy"
	"github.com/cynic89/hackday/readSpreadsheet"
)

func main() {

	srv, err := hobbybuddy.Login()
	if err != nil {
		fmt.Printf("Error when trying to login: %s", err.Error())
	}
	hobbies := readSpreadsheet.ReadSpreadsheet();

	for hobby, emails := range hobbies {
		if len(emails) > 1 {
			hobbybuddy.CreateEvent(hobby, emails, srv)
		}
	}

}
