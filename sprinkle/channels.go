package main

import (
	"fmt"
	"time"
)

//START OMIT
func printMessage(ch chan int) {
	fmt.Println("In printMessage: Waiting for message...")
	fmt.Println("In printMessage: Received: ", <-ch)
}

//END OMIT

func main() {

	//START_1 OMIT

	ch := make(chan int)

	go printMessage(ch)

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Sending message..")
	ch <- 200

	//Stop main thread for 1s (DON'T DO THIS IN A REAL PROGRAM)
	time.Sleep(1000 * time.Millisecond) 

	//END_1 OMIT
}
