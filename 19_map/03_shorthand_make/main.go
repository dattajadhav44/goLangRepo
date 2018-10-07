package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)    // another way of declaration
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	myGreeting["Dattatray"] = "Good morning!"

	fmt.Println(myGreeting)

	myGreeting2 := map[string]string{}
	myGreeting2["Steph"] = "Good morning."
	myGreeting2["Paul"] = "Good afternoon."

	fmt.Println(myGreeting2)
}
