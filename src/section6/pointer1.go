package main

import "fmt"

func main() {
	// 포인터
	// golang : 포인터 지원한다.
	// 변수의 지역성, 연속된 메모리 참조 ... , 힙/스택 ....
	// 주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
	// *(에스터리스크)로 사용
	// nil 로 초기화 (nil != 0)

	// 예제 1
	var a *int            // 방법 1
	var b *int = new(int) // 정석 선언

	i := 7
	a = &i // & 주소값 전달
	b = &i

	*a = 77

	fmt.Println("ex 1 : ", a, &i)
	fmt.Println("ex 1 : ", &a)
	fmt.Println("ex 1 : ", *a) // 역참조

	fmt.Println()

	fmt.Println("ex 1 : ", b, &i)
	fmt.Println("ex 1 : ", &b)
	fmt.Println("ex 1 : ", *b) // 역참조

	var c = &i
	d := &i

	fmt.Println()

	fmt.Println("ex 2 : ", c, &i)
	fmt.Println("ex 2 : ", &c)
	fmt.Println("ex 2 : ", *c) // 역참조

	fmt.Println()

	fmt.Println("ex 2 : ", d, &i)
	fmt.Println("ex 2 : ", &d)
	fmt.Println("ex 2 : ", *d) // 역참조

}
