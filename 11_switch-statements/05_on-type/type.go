package main

import "fmt"

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}
func main() {
	SwitchOnType(7)
	SwitchOnType("Dattatray")
	var k = contact{" Glad to see you", "You  are also active in the Blockchain space"}
	SwitchOnType(k)
	fmt.Println(k)
	fmt.Printf("%T \n", k)
	SwitchOnType(k.greeting)
	SwitchOnType(k.name)
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
   switch x.(type) { // this is an assert; asserting, "x is of this type"
   case int:
	fmt.Println("This is integer type")
   case string:
	fmt.Println("This is a string type")
   case contact:
	fmt.Println("This is a contact type")
   default:
	fmt.Println("The default is here")
   }
}


