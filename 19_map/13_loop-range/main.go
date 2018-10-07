package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Datatray!",
		2: "Go Programming",
		3: "Rocks!",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
