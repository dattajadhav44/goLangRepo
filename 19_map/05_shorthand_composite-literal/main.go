package main

import "fmt"

func main() {

    MyGreeting := map[string] string {
		"Datta" : "Jadhav",
		"Kuldeep" : "Yadav",
	}

	fmt.Println(MyGreeting["Kuldeep"])

}

