package main

import "fmt"

func main() {
	// 클로저(closure)
	// 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이때, 외부 함수에 선언 된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에 함수가 실행 될때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 함수로 호출 할때, 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트, 무분별한 전역변수를 남발하는걸 막기 위함 ...
	// 클로저 남발 -> 객체들이 메모리에 존재하므로, -> 메모리 부족, 오버플로우 현상이 생길 수 도 있다
	// 따라서 정확하게 이해하고 사용해야한다.

	// 예제 1
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)
	fmt.Println("ex 1 : ", r1)

	// 예제 2
	m, n := 5, 10 // 이 변수들이 익명함수 안에서 사용하기 때문에 스냅샷(객체)가 된것이다.
	sum := func(c int) int {
		return m + n + c // 지역 변수 소멸되지 않는다. (함수 호출 시 마다 지역변수 호출 가능) == 클로저
	}

	r2 := sum(10)
	fmt.Println("ex 2 : ", r2)
}
