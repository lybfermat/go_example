package main 
import "fmt"

func main() {
	i := 100
	for j := 7; j < i; j++ {
		fmt.Println(j)
	}

	for i < 120 {
		fmt.Println(i)
		i = i + 1
	}

	for {
		fmt.Println("in loop")
		break
	}

	for k := 90; k < i; k++ {
		if(k%2 == 0) {
			continue
		}
		fmt.Println(k)
	}
}