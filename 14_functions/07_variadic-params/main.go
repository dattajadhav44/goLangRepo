package main 
import "fmt"
func main() {
	s := variadicFunction(43,45,65,70,80)
	fmt.Println(s)
}

func variadicFunction(sf ...float64) float64 {
	fmt.Println(sf)

	var total float64

	for err, v := range sf {
		total += v
	}

	fmt.Println(total)
	fmt.Println(float64(len(sf)))

	return total / float64(len(sf))
}

