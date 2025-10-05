package test_programs

import (
	"fmt"
	"regexp"
)

func CoffeeTest() {

	fmt.Print("Coffee Test")
	// FindingAll()
	// FindingAllString()
	FindingSub()

}

func FindingAll() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))

}

func FindingAllString() {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
}

func FindingSub() {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1))

	re2 := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re2.FindAllStringSubmatch(`seafood fool`, -1))

	/*
	re3 := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re3.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re3.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re3.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re3.FindAllStringSubmatch("-axxb-ab-", -1))
	*/

}
