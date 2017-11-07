package main 
import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1;
	} 
	return n * fact(n-1)
}

func main() {

	n := fact(1)
	n1 := fact(10)

	fmt.Println(n, n1)
}