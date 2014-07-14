package main

import (
	"fmt"
)

//START OMIT
type Animal struct {
	Name     string
	domestic bool
}

func (a Animal) isDomestic() bool {
	return a.domestic
}

func (a Animal) String() string {
	return fmt.Sprintf("This is %s", a.Name)
}

func main() {
	dog := Animal{Name: "Dog", domestic: true}
	fmt.Printf("%s\n", dog)
}

//END OMIT
