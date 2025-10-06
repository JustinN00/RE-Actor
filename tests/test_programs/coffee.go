package test_programs

import (
	"fmt"
	"regexp"
	"strings"
)

func CoffeeTest() {

	fmt.Print("Coffee Test\n\n")

	teststring := "5.10.2025 17:17:52 [Server Event] Player SueHex left."
	check := regexp.MustCompile(`Player`)
	checking := check.FindAllSubmatch([]byte(teststring), -1)
	testtype := []byte(teststring)
 
	switch any(testtype).(type){
	case string:
		fmt.Println("string")
	case byte:
		fmt.Println("byte")
	}

	if len(checking) > 0 {
		fmt.Print("player found\n")
		ev := regexp.MustCompile(`Player (.*?)\.`)

		event := ev.FindStringSubmatch(teststring)
		eventstring := event[1]
		fmt.Printf("%T\n\n",eventstring)

		split := strings.Split(eventstring, " "); player, action := split[0], split[1]

		fmt.Println(player)
		fmt.Println(action)
	}

}
