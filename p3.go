package main

import (
	"fmt"
)

func add(x float64, y float64) float64 {
	return x + y
}
func multiple(a, b string) (string, string) {
	return a, b
}
func main() {
	// var num1 float64 = 5.5
	// var num2 float64 = 9.9
	s1, s2 := "Hey", "bitch"
	// fmt.Println(add(num1, num2))
	fmt.Println(multiple(s1, s2))
}
