package test_programs

import(
	"fmt"
	"regexp"
)

func CoffeeTest(){
	
	fmt.Print("Coffee Test")
	FindingAll()
	FindingAllString()

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