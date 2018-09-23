package main 

import "fmt"

func main() {
	greeting := makeGreeting() 
	fmt.Println(greeting())
	fmt.Printf("%T\n", greeting)
}

func makeGreeting() func () string {
	return func() string {
		return "Hello Mr.Dj"
	}
}


