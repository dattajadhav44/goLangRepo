

package main
import "fmt"
const (
	_ = iota
	b = iota * 10
	c = iota * 20
	d = iota * 30
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}