
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

//using switch
func SelectOperation(number int){
	switch number{
	case 1:
	fmt.Println(MultipleFunc(10,10))
	case 2:
	fmt.Println(DivisionFunc(12,1))	
	case 3:
		fmt.Println(SumFunc(10,10))
	case 4:
		fmt.Println(SubsractFunc(12,1))
	default:
		fmt.Println("Bilinmeyen Numara!")
	}
}

//using for loop
func forLoop(start int, limit int,count int)int{
		toplam:= 0

for i:=start;i<=limit;i+=count{
	fmt.Println(i)
	toplam=toplam+i
}
 return toplam;
}

func main(){
	SelectOperation(5)
	// fmt.Println("Sum:",SumFunc(12,16))
	// fmt.Println("Division:",DivisionFunc(8.0,25.0))
	// fmt.Println("Substract:",SubsractFunc(23,12))
	// fmt.Println("Multiple:",MultipleFunc(23,12))
	fmt.Println("ForLoop:",forLoop(0,4,1))
}