package main

import (
	"fmt"
)

func main() {
	x := 15
	a := &x
	fmt.Println(a)  //adress
	fmt.Println(*a) //pointer value
	*a = 5          //change pointer value
	fmt.Println(x)  //print new x value
	*a = *a * *a    //change value again
	fmt.Println(a)  //new adress
	fmt.Println(*a) //new value 25
}
