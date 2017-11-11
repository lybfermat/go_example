package main 
import "fmt"
import "sort"

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	fruits := []string{"apple", "pear", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}