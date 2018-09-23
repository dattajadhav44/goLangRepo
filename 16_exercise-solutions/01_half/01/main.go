package main 
import "fmt"

func main() {
	res, _odd := execute(6)
	fmt.Println(res,_odd)
}

func execute(n int) (int, bool) {
	return n/2, n%2 == 0
}
