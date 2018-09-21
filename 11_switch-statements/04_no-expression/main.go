package main

import "fmt"

func main() {

	myMatchingCase := "TestCase"
	switch {
	case len(myMatchingCase) == 4:
		fmt.Println("The lenght is matched")
	case myMatchingCase == "Sehwag":
		fmt.Println("This is a viru")
	case myMatchingCase == "GoLang":
		fmt.Println("This is a go language")
	default:
		fmt.Println("Nothing matched, This is a default")
	}
}


/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/
