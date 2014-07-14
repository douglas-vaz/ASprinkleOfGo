package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

//START_FEED OMIT
func feed(dir string) <-chan string {
	out := make(chan string)
	files, _ := ioutil.ReadDir(dir)
	go func() {
		for _, f := range files {
			path := dir + "/" + f.Name()
			out <- path
		}
		close(out)
	}()
	return out
}

//END_FEED OMIT
//START_READ OMIT
func read(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for f := range in {
			b, err := ioutil.ReadFile(f)
			if err != nil {
				panic(err)
			}
			out <- string(b)
		}
		close(out)
	}()
	return out
}

//END_READ OMIT
//START_COUNT OMIT
func count(in <-chan string) <-chan int64 {
	out := make(chan int64)
	go func() {
		for text := range in {
			out <- int64(len(strings.Fields(text)))
		}
		close(out)
	}()
	return out
}

//END_COUNT OMIT
func main() {

	//START_MAIN OMIT
	//set concurrency
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	//setup pipeline
	files := feed("CORPUS_TXT")
	text := read(files)
	out := count(text)

	//Reduce
	var total int64
	for n := range out {
		total = total + n
	}

	//Output
	fmt.Println(total)

	//END_MAIN OMIT
}
