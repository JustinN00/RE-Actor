package main

import (
	"fmt"
	"os"

	"github.com/SusanHex/RE-Actor/tests/test_programs"
)

func main() {
	user := os.Getenv("DEV_USER")

	switch user {

	case "Kohi-Kamisama":

		test_programs.CoffeeTest()

	case "SusanHex":

		test_programs.SusanTest()

	case "":

		fmt.Println("User Not Found")
		fmt.Println("\nCreat a custom enviroment varibal called DEV_USER and set it to your GitHub name")

	default:

		fmt.Println("Unknown User")
		fmt.Println("\nUser not in switch")

	}

}
