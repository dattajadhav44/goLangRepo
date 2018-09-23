package main
import "fmt"
func main() {
	var x [60]int
	fmt.Println(x)
	x[50] = 555
	fmt.Println(x[50])
	fmt.Println(x)
}
