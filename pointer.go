package main 

import "fmt"

func zeroVal(v int) {
	v = 0;
}

func zeroPtr(v *int) {
	*v = 0
}
func main() {

	i := 1

	fmt.Println("initial i:", i)
	zeroVal(i)
	fmt.Println("zeroVal i:", i)
	zeroPtr(&i)
	fmt.Println("zeroPtr i:", i)
	fmt.Println("pointer i:", &i)
}