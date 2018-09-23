
package main 

import "fmt"

func main() {
	var arr [256]int
	for i := 0; i < 256; i++ {
		arr[i] = i
	}
    for i, m := range arr {
		fmt.Printf("%v - %T - %b\n", m, m, m)
		if i > 50 { 
		  break
		}
	}
}
