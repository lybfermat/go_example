package main 
import ( "fmt"
         "math"
        )

const s string = "hello const"
func main() {

	fmt.Println(s)
	const n = 10000
	fmt.Println(n)
 
    //const d := 2e5/n error
	const d = 2e5/n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))

	
}