package main 

import "fmt"

func main() {

	res,_bool := executeMe(7)
	fmt.Println(res,_bool)
}

func executeMe(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}
