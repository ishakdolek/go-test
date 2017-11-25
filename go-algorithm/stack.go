package main

import (
	"fmt"
)


type Stack struct{
	data [4]int
	top int
}
 func Push(stk Stack,number int){
	if(stk.top<=4){
fmt.Println("Dolu!")

	}
	stk.top++;
	stk.data[stk.top]=number
}
