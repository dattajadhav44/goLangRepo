package main

import "fmt"

func main() {

	name := "Dattatray"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Bhau
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Dattatray
	*z = "Bhau"
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Bhau
}
