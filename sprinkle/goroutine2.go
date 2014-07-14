package main

import (
	"fmt"
	"time"
)

//START_LOOP1 OMIT
func loop() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}

//END_LOOP1 OMIT

func main() {

	//START_LOOP3 OMIT
	fmt.Println("Before calling loop..")
	go loop()
	fmt.Println("After calling loop..")

	time.Sleep(300 * time.Millisecond)

	//END_LOOP3 OMIT

}
