package main

import (
	"fmt"
	"strconv"
)

// 함수 선언 위치는 어느 곳이든 상관 없음
func helloGolang() {
	fmt.Println("ex 1 : Hello Golang!!")
}

func sayOne(m string) {
	fmt.Println("ex 2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	// 함수
	// 선언 시 : func 키워드로 선언
	// func 함수명 (매개변수) (반환타입 Or 반환 값 변수명) -> 반환 값 존재
	// func 함수명 () (반환타입 Or 반환 값 변수명) -> 매개변수 X, 반환 값 존재
	// func 함수명 (매개변수) -> 매개변수 존재, 반환 값 X
	// func 함수명() -> 매개변수 X, 반환 값 X
	// 타 언어와 달리 리턴 값이 다수 가능하다.

	// 예제 1
	helloGolang()

	// 예제 2
	sayOne("function 2~~")

	// 예제 3
	result := sum(5, 5)
	fmt.Println("ex 3 : ", result)
	fmt.Println("ex 3 : ", sum(5, 5))
	fmt.Println("ex 3 : ", strconv.Itoa(sum(5, 5)))

}
