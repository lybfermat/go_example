package main 

import "fmt"
import "os"
 func main( ){
 	f := createFile("./help")
 	defer closeFile(f)
 	writeFile(f)
 }

 func createFile(fname string) *os.File {
 	fmt.Println("crearte file ")
 	f, err := os.Create(fname)
 	if err != nil {
 		panic(err)
 	}
 	return f
 }

 func writeFile(f *os.File) {
 	fmt.Println("writing")
 	fmt.Fprintf(f, "hello")
 }

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}