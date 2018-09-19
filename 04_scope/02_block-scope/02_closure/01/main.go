
package main

import "fmt"

func main() {
   x := 42
   fmt.Println(x) 
   {
	   fmt.Println(x)
	   y := "The credit belongs to one who is in the ring, not accessable for outsiders"
	   fmt.Println(y)
   }

   //fmt.Println(y)// these will not compile as the y is outside of scope
}