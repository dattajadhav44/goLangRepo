package main

import "fmt"

func main() {
       switch "Marcus" {
	   case "Tim":
		   fmt.Println(" whats up Tim")
	   case "Jenny":
		   fmt.Println(" whats up Jenny")
	   case "Marcus":
			fmt.Println(" Whats up Marcus")
			fallthrough
	   case "Stev":
			fmt.Println(" whats up Stev")
			fallthrough
	   case "Scott":
			fmt.Println(" Whats up Scott")
	   case "Tendulkar":
		    fmt.Println(" Whats up Sachin. The country has a huge respect for you!!!")
	   }

}
