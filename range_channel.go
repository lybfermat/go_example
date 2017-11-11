package main 
import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "tow"
	close(queue)

	for i := range queue {
		fmt.Println(i)
	}
}