package main

import "fmt"


func main() {

	var a bool
	a = false 

	if false {
		fmt.Println("This should print False")
	} else if true {
		fmt.Println("This should print True----")
	} else {
		fmt.Println("This should print out of the BOX :D")
	}
}

