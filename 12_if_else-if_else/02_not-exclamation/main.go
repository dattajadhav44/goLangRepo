package main

import "fmt"

func main() {

	if !true {
		fmt.Println("This should print False")
	}

	if !false{
		fmt.Println("This should print True")
	}
}
