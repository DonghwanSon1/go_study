package main

import "fmt"

func ex_f1() {
	fmt.Println("f1 : Start!")
	defer ex_f2() // 마지막에 호출된다.
	fmt.Println("f1 : End!")
}

func ex_f2() {
	fmt.Println("f2 : Called!")

}

func main() {
	// Defer 함수 실행 (지연 실행이다)
	// Defer 를 호출한 함수가 종료되기 직전에 호출 된다.
	// 타 언어의 finally 문 과 비슷하다.
	// 주로 리소스 반환, 열린 파일 닫기, Mutex 잠금 해제

	// 예제 1
	ex_f1()

}
