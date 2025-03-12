package main

import (
	"fmt"
)

func runFunc() {
	// defer 를 먼저 선언 후 아래 패닉을 선언해야한다.
	defer func() { // 순서 2 -> defer 함수는 지연함수 즉 finally 같은 느낌의 함수이므로 패닉이 발생 후 여기로 들어온다.
		s := recover()                     // 그 후 recover 로 해당 에러를 가져오고,
		fmt.Println("Error Message : ", s) // 가져온 에러를 보여준다.
	}()

	panic("Error occurred!") // 순서 1
	fmt.Println("TEST")      // 위에서 패닉이 발생했기에 호출 안됨.
}

func main() {
	// Golang panic / recover 함수
	// 에러 복구 가능
	// Panic 으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료 되지 않는다.)
	// 즉, 코드 흐름을 정상상태로 복구하는 기능
	// Panic 에서 설정한 메시지를 받아 올 수 있다.

	// 예제
	runFunc()

	fmt.Println("Hello Golang !!")

}
