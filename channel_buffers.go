package main 

import "fmt"

func main() {
	message := make(chan string, 2)
	go func() { 
		message <- "ping"
		message <- "pong"
	}()

	msg := <-message
	fmt.Println(msg)
	msg = <-message
	fmt.Println(msg)		
}