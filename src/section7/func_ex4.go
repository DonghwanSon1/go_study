package main

import "fmt"

func main() {
	// 함수 고급
	// 익명 함수(Anonymous Function)
	// 즉시 실행(선언과 동시에)

	// 예제 1
	func() {
		fmt.Println("ex 1 : Anonymous function")
	}()

	// 예제 2
	msg := "Hello Golang"
	func(m string) {
		fmt.Println("ex 2 : ", m)
	}(msg)

	// 예제 3
	func(x, y int) {
		fmt.Println("ex 3 : ", x+y)
	}(10, 20)

	// 예제 4
	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("ex 4 : ", r)

}
