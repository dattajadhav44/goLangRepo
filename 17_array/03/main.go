
package main

import "fmt"

func main() {
	var arr [256]int

	fmt.Println(len(arr))
	fmt.Println(arr[42])
	for i := 0; i < 256; i++ {
		arr[i] = i
	}
	for i, v := range arr {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
