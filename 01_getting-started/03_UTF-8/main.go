package main

import "fmt"

func main() {
	for i := 60; i < 120; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
		//fmt.Printf("%d \n",i)
		//fmt.Printf("%b \n",i)
		//fmt.Printf("%x \n",i)
		//fmt.Printf("%q \n",i)
	}
}
