package main

import "fmt"

func main() {
	for i := 0; i<=100; i++ {
		if i % 3 == 0 {
			fmt.Println("There is no reminder for this ", i)
		}
	}
}

