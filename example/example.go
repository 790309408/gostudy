package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[6:] // [6, 7, 8, 9]
	fmt.Println(s1)
	s2 := data[:3] // [0, 1, 2]
	fmt.Println(s2)
	copy(s1, s2)
	fmt.Println("copy slice s1 : ", s1)
	fmt.Println("copy slice s2 : ", s2)
	fmt.Println("slice data : ", data)
	// s1   [0, 1, 2, 9]
	// s2   [0, 1, 2]
	// data [0, 1, 2, 3, 4, 5, 0, 1, 2, 9]
}
