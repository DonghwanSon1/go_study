package main

import "fmt"

// ... 으로 할 시 몇개가 들어와도 상관 없다 라는 표현임
func multiply3(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum2(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) string {
	var totalMsg string
	for _, value := range msg {
		totalMsg += value + " "
	}
	return totalMsg
}

func main() {
	// 함수 고급편
	// 가변 인자 실습 (매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	// 예제 1
	x := multiply3(1, 2, 3)
	y := sum2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("ex 1 : ", x)
	fmt.Println("ex 1 : ", y)

	// 예제 2
	msg := prtWord("Hello", "World", "Son")
	fmt.Println("ex 2 : ", msg)

	// 예제 3
	a := []int{1, 2, 3, 4, 5}
	m := multiply3(a...) // 슬라이스의 각각의 int 로 매개변수로 주는것이다.
	n := sum2(a...)

	fmt.Println("ex 3 : ", m)
	fmt.Println("ex 3 : ", n)
}
