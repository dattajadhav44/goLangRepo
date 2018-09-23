package main

import "fmt"

func main() {

     xs := wrapper([]int{1,2,3,4}, func(n int) bool {
		 fmt.Println(n) 
		   return n > 1
	 } )
	 fmt.Println(xs)
}

func wrapper(number []int, callback func(int) bool) []int {
	var res []int
	for _,n := range number {
		if callback(n) {
			res = append(res,n)
			fmt.Println(res)
		} else {
			fmt.Println("the number is ONE",n)
		}
	}
	return res
}
