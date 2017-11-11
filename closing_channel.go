package main 
import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		 for{
		 	j, more := <- jobs
		 	if more {
		 		fmt.Println("recv job:", j)
		 	} else {
		 		fmt.Println("recv all")
		 		done <- true
		 		return
		 	}
		 }

	}()

	for j := 0 ; j< 10; j++ {
		jobs <- j
		fmt.Println("send-job", j)
	}
	close(jobs)
	<-done
}
