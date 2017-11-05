package main 
import "fmt"

func main() {

	m := make(map[string]int)
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println("map: ", m)

	v := m["key1"]
	fmt.Println(v)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)		
}