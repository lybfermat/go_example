package main 
import "fmt"

func MultiReturn(a int, b int) (int,int) {
	return a+b, a-b
}

func main() {
	plus, sub := MultiReturn(10, 20);
	fmt.Println("plus: ", plus)
	fmt.Println("sub: ", sub)
}