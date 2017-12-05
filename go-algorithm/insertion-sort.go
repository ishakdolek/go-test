 package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	numbers := os.Args[1]

// 	var arr []int
// 	var num int
// 	var err error
// 	for _, val := range strings.Split(numbers, ",") {
// 		num, err = strconv.Atoi(val)

// 		if err != nil {
// 			log.Fatal("Failed")
// 		}

// 		arr = append(arr, num)
// 	}

// 	for j := 1; j < len(arr); j++ {
// 		key := arr[j]
// 		i := j - 1

// 		for i >= 0 && arr[i] > key {
// 			arr[i+1] = arr[i]
// 			i = i - 1
// 		}

// 		arr[i+1] = key
// 	}

// 	fmt.Println(arr)
// 	fmt.Println("The test finished!")
// }

// func Insertion(items[]int){
	
// 	for j := 1; j < items.len; j++ {
// 		key := arr[j]
// 		i := j - 1

// 		for i >= 0 && arr[i] > key {
// 			arr[i+1] = arr[i]
// 			i = i - 1
// 		}

// 		arr[i+1] = key
// 	}

// }