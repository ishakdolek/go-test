
// Sum divide string multiple subsract
package main

import (
	"fmt"
)
func SumFunc(n1 int, n2 int)(int){
	return n1+n2
}

func DivisionFunc(n1 float32, n2 float32)(float32){
	return n1/n2
}
func SubsractFunc(n1 int, n2 int)(int){
	return n1-n2
}

func MultipleFunc(n1 int, n2 int)(int){
	return n1*n2
}

func main(){
	fmt.Println("Sum:",SumFunc(12,16))
	fmt.Println("Division:",DivisionFunc(8.0,25.0))
	fmt.Println("Substract:",SubsractFunc(23,12))
	fmt.Println("Multiple:",MultipleFunc(23,12))
}