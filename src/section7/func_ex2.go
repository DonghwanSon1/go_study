package main

import "fmt"

func multiply4(x, y int) (r int) {
	r = x * y
	return r
}

func sum3(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// 함수 고급
	// 함수를 변수에 할당
	// 예제 1 - 슬라이스에 할당
	f := []func(int, int) int{multiply4, sum3}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println("ex 1 : ", a, f[0](10, 10))
	fmt.Println("ex 1 : ", b, f[1](10, 10))

	// 예제 2 - 변수에 할당
	var f1 func(int, int) int = multiply4
	f2 := sum3
	fmt.Println("ex 2 : ", f1(10, 10))
	fmt.Println("ex 2 : ", f2(10, 10))

	// 예제 3 - map에 할당
	m1 := map[string]func(int, int) int{
		"mul_func": multiply4,
		"sum_func": sum3,
	}
	fmt.Println("ex 3 : ", m1["mul_func"](10, 10))
	fmt.Println("ex 3 : ", m1["sum_func"](10, 10))

}
