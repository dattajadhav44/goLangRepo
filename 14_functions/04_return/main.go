package main

import "fmt"

func main() {
	fmt.Println(meeting("Blockchain ","folks"))
}

func meeting(tech,people string) string {
	return fmt.Sprint(tech,people)
}


