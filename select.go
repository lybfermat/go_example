package main 
import "fmt"
import "time"
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second *2)
		c1 <- "two"
	}()

	for i := 0 ;i < 2; i++ {
		select {
		case msg := <- c1:
			fmt.Println("recv from c1:", msg)
		case msg := <- c2:
			fmt.Println("recv from c2:", msg)
		}
	}

}