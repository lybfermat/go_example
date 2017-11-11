package main 

import "fmt"
import "time"

func worker(id int, jobs <- chan int, results chan <- int) {
	for i := range jobs {
		fmt.Println("worker id:", id, "started job:", i)
		time.Sleep(time.Second *1)	
		fmt.Println("worker:",id, "finish job:",i)
		results <- i *22
	}
}

func main() {
	jobs := make (chan int, 100)
	results := make(chan int, 100)
	 for w := 1; w <= 3; w++ {
	 	go worker(w, jobs, results)
	 }

	 for j := 1; j <= 5; j++ {
	 	jobs <- j
	 }

	 close(jobs)
	 for a := 1; a <= 5; a++ {
	 	<-results
	 }
}