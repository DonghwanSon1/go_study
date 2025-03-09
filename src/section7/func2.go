package main

import "fmt"

func sum1(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println("ex 1 : ", a+b)
}

func multiValue(i int) {
	i = i * 10
}

func multiReference(i *int) {
	*i *= 10
}

func main() {
	// 함수(콜백) 전달, 참조 전달, 값 전달

	// 예제 1
	sum1(100, add)

	// 예제 2 - 값 전달
	a := 100
	multiValue(a)
	fmt.Println("ex 2 : ", a)

	// 예제 3 - 참조 전달
	b := 100
	multiReference(&b)
	fmt.Println("ex 3 : ", b)
}
