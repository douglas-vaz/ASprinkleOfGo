package main

import (
	"fmt"
)

//START_LOOP1 OMIT
func loop() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

//END_LOOP1 OMIT

func main() {

	//START_LOOP2 OMIT
	fmt.Println("Before calling loop..")
	go loop()
	fmt.Println("After calling loop..")

	//END_LOOP2 OMIT

}
