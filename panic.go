package main 

import "os"
func main() {
	panic("a probem")

	_, err := os.Create("/data/abc")
	if err != nil {
		panic(err)
	}
}