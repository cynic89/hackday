package main

import (
	"fmt"
	"github.com/cynic89/hackday/hobbybuddy"
	"github.com/cynic89/hackday/readSpreadsheet"
	"math/rand"
)

func main() {

	srv, err := hobbybuddy.Login()
	if err != nil {
		fmt.Printf("Error when trying to login: %s", err.Error())
	}
	hobbies := readSpreadsheet.ReadSpreadsheet()
	randomize(hobbies)
	
	for hobby, emails := range hobbies {
		if len(emails) > 1 {
			hobbybuddy.CreateEvent(hobby, emails, srv)
		}
	}

}

func randomize(hobbies map[string][]string) {
	randUsers := hobbies["Random"]
	for _, u := range randUsers {
		randHobby := randHobby(u, hobbies)
		if randHobby != "" {
			hobbies[randHobby] = append(hobbies[randHobby], u)
		}
	}

	delete(hobbies, "Random")
}

func randHobby(user string, hobbies map[string][]string) string {
	possibleTargets := hobbiesNotPartOf(user, hobbies)
	if len(possibleTargets) > 0 {
		randNum := rand.Intn(len(possibleTargets))
		return possibleTargets[randNum]
	}

	return ""
}

func hobbiesNotPartOf(user string, hobbies map[string][]string) (possibleTargets []string) {
	for hobby, attendees := range hobbies {
		if !contains(user, attendees) {
			possibleTargets = append(possibleTargets, hobby)
		}
	}
}

func contains(user string, attendees []string) bool {
	for _, a := range attendees {
		if a == user {
			return true
		}
	}
	return false
}
