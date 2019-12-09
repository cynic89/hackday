package main

import (
	"fmt"
	"github.com/cynic89/hackday/hobbybuddy"
)

func main() {

	srv, err := hobbybuddy.Login()
	if err != nil {
		fmt.Printf("Error when trying to login: %s", err.Error())
	}

	hobbybuddy.CreateEvent("Pool",
		[]string{"ssankaran@pivotal.io", "rshenoy@pivotal.io", "neverma@pivotal.io"}, srv)

}
