package main

import "fmt"

func main() {
	// 문장 끝 세미클론(;) 주의 -> 자동으로 끝에 세미클론 삽입 하기 때문에 붙이지 않는다.
	// -> 단, 두 문장을 한 문장에 표현할 경우 명시적으로 세미클론 사용해야한다. (권장 X) -> EX) fmt.Println("ex 1 : ", i); fmt.Println("ex 2 : ", i)
	// 반복문 및 제어문에서는 세미클론 필요하다. (선언 시)

	// 예제 1
	for i := 0; i <= 10; i++ {
		fmt.Print("ex 1 : ")
		fmt.Println(i)
	}

	fmt.Println()

	// 예제 2
	for j := 10; j >= 0; j-- {
		fmt.Println("ex 2 : ", j)
	}
}
