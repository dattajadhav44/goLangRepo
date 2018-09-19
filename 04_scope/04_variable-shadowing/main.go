package main
import "fmt"
func addTo(x int) int {
	return 50 + x
}

func main() {
	result := addTo(8)
	fmt.Println(result)
}

