package main

import (
	"fmt"
)

//START OMIT
func evenNums() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func main() {
	nextEven := evenNums()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

//END OMIT
