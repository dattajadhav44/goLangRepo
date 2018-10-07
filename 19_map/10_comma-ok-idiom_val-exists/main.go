package main 

import "fmt"

func main() {

    takeGreeting := map[int]string{
		0 : "Zero",
		1 : "One",
		2 : "Two",
		3 : "Three",
	}

	fmt.Println(takeGreeting)
	delete(takeGreeting, 2)

	if val, exists := takeGreeting[2]; exists {
		fmt.Println("The value is exist")
		fmt.Println("The value is :", val)
		fmt.Println("Indicator :", exists)

	} else {
		fmt.Println("The value does not exists")
		fmt.Println("The value is :", val)
		fmt.Println("Indicator :", exists)
	}

}

