package main

import (
	"fmt"
)

func runFunc1() {
	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println("ex 1 : ", a[i]) // 에러 발생 (3개의 배열에서 4,5번의 인덱스를 호출하기 때문에) -> 인덱스 범위 초과 Panic 발생
	}
}

func main() {
	// Golang panic / recover 함수
	// 에러 복구 가능
	// Panic 으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료 되지 않는다.)
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// Panic 에서 설정한 메시지를 받아 올 수 있다.

	// 예제
	runFunc1()

	fmt.Println("Hello Golang !!")

}
