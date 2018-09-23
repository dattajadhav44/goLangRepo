package main

import "fmt"

func hello() {
	fmt.Print("hello \n")
}

func world() {
	fmt.Println("world")
}

func main() {
	world()
	hello()
}
