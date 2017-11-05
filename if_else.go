package main 
import "fmt"

func  main() {
	if 7%2 == 0 {
		fmt.Println("7 is a even")
	} else {
		fmt.Println("7 is odd")
	}
	
	if 8 % 4 == 0 {
		fmt.Println("8 is div by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(" is nagetive")
	} else if num  < 10 {
		fmt.Println("on digit")
	} else {
		fmt.Println("is multi digits")
	}
}