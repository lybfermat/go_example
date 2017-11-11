package main 

import "fmt"

func main () {
	messages := make(chan string)
	signals := make(chan string)
	select {
	case msg := <-messages:
		fmt.Println("recv message from messages:", msg)
	default :
		fmt.Println("not recv msg")
	}
	msg := "hi"
	select {
	case messages <-msg:
		fmt.Println("send msg:", msg)
	default:
		fmt.Println("not send msg")
	}
	select {
	case msg := <- messages:
		fmt.Println("recv msg from messages", msg)
	case msg := <- signals:
		fmt.Println("recv msg from messages", msg)
	default:
		fmt.Println("not ativity")

	}
}