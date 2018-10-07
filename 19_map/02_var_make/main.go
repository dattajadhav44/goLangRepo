package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Dattatray"] = "Good morning."
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	

	fmt.Println(myGreeting)
}
