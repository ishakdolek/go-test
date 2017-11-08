package main

import (
    "fmt"
)

var toBeSorted [10]int = [10]int{1,3,2,4,8,6,7,2,3,0}

func bubbleSort(input [10]int) {
    n := 10
	swapped := true
	fmt.Println("Swapping")
    for swapped {
        swapped = false
        for i := 1; i < n-1; i++ {
            if input[i-1] > input[i] {             
                // swap values using Go's tuple assignment
                input[i], input[i-1] = input[i-1], input[i]
                swapped = true
            }
        }
    }
    fmt.Println(input)
}


func main() {
    fmt.Println("Sorting by Bubble Sort")
    bubbleSort(toBeSorted)
    
}