package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex 2 : (", n, ")", "hi")
	prtHello(n - 1)
}

func main() {
	// 함수 고급
	// 재귀 함수(Recursion)
	// 프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이해진다. -> 장점
	// 코드 이해하기 어렵고, 기억공간을 많이 사용, 무한 루프 가능성 존재 -> 단점

	// 예제 1
	x := fact(4)
	fmt.Println("ex 1 : ", x)

	// 예제 2
	prtHello(10)

}
