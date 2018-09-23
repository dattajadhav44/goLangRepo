package main

import "fmt"

func main() {
	greet("Dealer", "Track")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
