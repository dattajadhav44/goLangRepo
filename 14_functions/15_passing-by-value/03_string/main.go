package main 
import "fmt"

func main() {

	name := "Dattatray"
	country := "india"

	fmt.Println(name, country)
	fmt.Println(&name, &country)
	changeMe(&name, &country)
	//changeMe(name, country)
	fmt.Println(name, country)
	fmt.Println(&name, &country)
}

func changeMe(a *string, b *string) {
	fmt.Println("---------------------------")
	fmt.Println(a,b)
	fmt.Println(&a,&b)

	*a = "Bhausaheb"
	*b = "England"
}
