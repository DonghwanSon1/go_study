package main

import "fmt"

func main() {
	// 생략을 원할 시 _ 를 사용하여 생략한다.
	const (
		_ = iota
		A
		B
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println(A, B, C)
	fmt.Println(DEFAULT, SILVER, PLATINUM)
}
