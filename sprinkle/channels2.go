package main

import (
	"fmt"
	"math/rand"
	"time"
)

//START OMIT
func waitAndShout(ch chan string, n int) {
	for {
		time.Sleep(time.Duration(rand.Int()) * time.Millisecond)
		ch <- fmt.Sprintf("Hello from %d", n)
	}
}

func display(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}

//END OMIT

func main() {

	//START_1 OMIT

	ch := make(chan string)

	for i := 1; i <= 1000; i++ {
		go waitAndShout(ch, i)
	}
	go display(ch)

	time.Sleep(2 * time.Second)

	//END_1 OMIT

}
