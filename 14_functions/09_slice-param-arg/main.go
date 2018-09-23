package main
import "fmt"
func main() {
	data := []float64{10,20,30,40,50}

	res := sliceArg(data)
	fmt.Println(res)
}

func sliceArg(sf []float64) float64 {
	fmt.Println(sf)

	var total float64

	for _, v := range sf {
		total += v 
	}

	fmt.Println(total)
	fmt.Println(float64(len(sf)))

	return total / (float64(len(sf)))

}
