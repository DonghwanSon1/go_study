package main

import "fmt"

func main() {
	// 예제 1
	cnt := increaseCnt()
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
}

func increaseCnt() func() int {
	n := 0 // 지역변수(캡처됨)
	return func() int {
		n += 1
		return n
	}
}
