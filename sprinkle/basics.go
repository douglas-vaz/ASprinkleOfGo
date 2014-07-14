package main

import (
	"fmt"
)

//START FUNC OMIT
func add(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

//END FUNC OMIT

func main() {
	//START OMIT
	var x int
	x = 20

	y := 10
	//END OMIT

	fmt.Println(add(x, y))

}
