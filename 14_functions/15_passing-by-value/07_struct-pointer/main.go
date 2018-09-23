package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Dattatray", 44}
	fmt.Println(&c1.name) // 0x8201e4120
	fmt.Println(c1.name)

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x8201e4120
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Dattatray 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Bhau"
	fmt.Println(z)       // &{Bhau 44}
	fmt.Println(&z.name) // 0x8201e4120

}
