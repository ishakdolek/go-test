package main


// func main(){
// 	fmt.Println("Result:",fact(12))
// }

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
