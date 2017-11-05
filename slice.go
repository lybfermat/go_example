
package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apped:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)
	l := s[2:5]
	fmt.Println("slice1:", l)
	l = s[:5]
	fmt.Println("slice2:", l)
	l = s[2:]
	fmt.Println("slice3:", l)
	t := []string{"i", "j", "k"}
	fmt.Println("decl:", t)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
