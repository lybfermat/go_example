package main 
import "fmt"
import "sort"

func main() {
	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7,2,9}
	sort.Ints(ints)
	fmt.Println("ints:", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}