package main

import "fmt"

func main() {
	// 예제 1
	i := 7
	p := &i
	fmt.Println("ex 1 : ", i, *p)
	fmt.Println("ex 1 : ", &i, p)
	fmt.Println()

	*p++
	fmt.Println("ex 2 : ", i, *p)
	fmt.Println("ex 2 : ", &i, p)
	fmt.Println()

	*p = 77777
	fmt.Println("ex 3 : ", i, *p)
	fmt.Println("ex 3 : ", &i, p)
	fmt.Println()

	i = 123
	fmt.Println("ex 4 : ", i, *p)
	fmt.Println("ex 4 : ", &i, p)
	fmt.Println()

}
