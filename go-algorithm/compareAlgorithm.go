package main

import (
	"time"
	"fmt"
	"log"
)

func main(){
	fmt.Println("Sıralama Algoritmaları Karşılaştırma!")
	list := []int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}
	fmt.Println(list)		
	
	start := time.Now()
	fmt.Println("Quick Sort Start...")
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)		
	fmt.Println("Quick Sort End!")
	elapsed := time.Since(start)
 	log.Printf("Quick took %s", elapsed)

	start1:=time.Now()
	fmt.Println("Selection  Sort Start...")
	selectionsort(list)
    fmt.Println(list)		
	elapsed1:=time.Since(start1)
	fmt.Println(elapsed1)
	fmt.Println("Selection Sort End!:")
	
    fmt.Println("İnsertion  Sort Start...")
	// MergeSort(list)
    fmt.Println(MergeSort(list))		
	fmt.Println("İnsertion Sort End!")

	fmt.Println(elapsed)
    fmt.Println(elapsed1)
	
    fmt.Println("İnsertion  Sort Start...")
	bubbleSort(list)
    fmt.Println(list)		
	fmt.Println("İnsertion Sort End!")
}