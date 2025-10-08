package test_programs

import (
	"fmt"
	"regexp"
	"strings"
)

func CoffeeTest() {

	fmt.Print("Coffee Test\n\n")

	tstring := "5.10.2025 17:17:52 [Server Event] Player SueHex left."
	check := regexp.MustCompile(`Player`)
	checking := check.FindAllSubmatch([]byte(tstring), -1)
	testtype := []byte(tstring)
 
	var teststring string
	switch any(testtype).(type){
	case string:
		fmt.Printf("string\n")
		teststring = string(testtype)
	case []byte:
		fmt.Printf("byte\n")
		teststring = string(testtype)
	default:
		fmt.Printf("type: %T\n\n",testtype)
	}
	fmt.Println(teststring)

	if len(checking) > 0 {
		fmt.Print("\nplayer found\n\n")
		ev := regexp.MustCompile(`Player (.*?)\.`)

		event := ev.FindStringSubmatch(teststring)
		eventstring := event[1]

		split := strings.Split(eventstring, " "); player, action := split[0], split[1]

		fmt.Println(player)
		fmt.Println(action)
	}

}
