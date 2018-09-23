package main

import "fmt"

func main() {
	fmt.Println(greet("Dattatray ", "Jadhav"))
}

func greet(name, surname string) (s string) {
	s = fmt.Sprint(name + surname)
	return 
}




/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
