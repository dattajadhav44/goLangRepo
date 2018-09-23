package main

import "fmt"

func main(){
	fmt.Println(greet("Dattatray ", "Jadhav"))
}

func greet(fname string, lname string) (string, string,string) {
	return fmt.Sprint(fname,lname), fmt.Sprint("Block", "chain"), fmt.Sprint("Go ", "Program")
}

