package main

import ("fmt"
        //"bufio"
       ) 
func main() {
	
	var name string
	fmt.Print("Enter the user name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)

	var fullName string 
	fmt.Print("Enter the Full name: ")
	fmt.Scanln(&fullName)
	fmt.Println("Print name", fullName)


}