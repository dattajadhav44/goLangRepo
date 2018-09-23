package main

import "fmt"

func main() {
	callMe("Dj", "Call me")
}

func callMe(name string, action string){
	fmt.Println(name,action)
}
