package main

import "fmt"

func multiply1(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = x * 20
	return r1, r2
}

func multiply2(x, y int) (int, int) {
	return x * 10, x * 20
}

func main() {
	// 리턴 값 변수 사용

	// 예제 1
	a, b := multiply1(10, 5)
	fmt.Println("ex 1 : ", a, b)

	// 예제 2
	c, d := multiply2(10, 5)
	fmt.Println("ex 1 : ", c, d)

}
