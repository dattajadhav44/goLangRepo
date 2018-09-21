package main 

import "fmt"

func main() {
	for i := 0; i<=100; i++ {
		if i % 3 == 0 {
			fmt.Println("This is 3's devident-", i)
		}else if i % 5 == 0 {
			fmt.Println("This is 5's devident-", i)
		}else {
			fmt.Println("DONT PRINT")
		}
	}
}
