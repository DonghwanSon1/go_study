package main

import "fmt"

func a() {
	defer end(start("b"))
	fmt.Println("in Function a")
}

func start(msg string) string {
	fmt.Println("Start : ", msg)
	return msg
}

func end(msg string) {
	fmt.Println("End : ", msg)
}

func main() {
	// 예제 1 - defer 문 안에 중첩문이라면 defer 문의 첫번째에만 걸리고, 그 첫번째안에 함수는 상관없다.
	// 따라서 start 먼저, a , end 이렇게 작동한다.
	a()
}
