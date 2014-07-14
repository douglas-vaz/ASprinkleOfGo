package main

import (
	"fmt"
)

//START_1 OMIT

type Mover interface {
	Move()
}

func goSomewhere(m Mover) {
	m.Move()
}

//END_1 OMIT

//START_V OMIT

type Vehicle struct {
}

func (v Vehicle) Move() {
	fmt.Println("I'm Moving")
}

//END_V OMIT

//START OMIT
type Car struct {
	Vehicle
	driver string
}

func (c Car) Move() {
	fmt.Println(c.driver, "is driving")
}

//END OMIT
func main() {
	//MAIN OMIT
	thing := new(Vehicle)
	car := Car{driver: "Tom"}

	goSomewhere(thing)
	goSomewhere(car)
	//MAIN_END OMIT
}
