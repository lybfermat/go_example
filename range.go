package main 
import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if(num == 1) {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]int{"a":1,"b":2,"c":3}
	for k,v := range kvs {
		fmt.Println("%s -> %d", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "hello" {
		fmt.Println(i,c)
	}
}