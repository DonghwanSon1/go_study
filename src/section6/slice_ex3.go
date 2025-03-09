package main

import "fmt"

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// make로 공간을 할당 후 복사 해야한다.
	// 복사 된 슬라이스 값 변경해도 원본에는 영향 없음

	// 예제 1 - 복사
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println("ex 1 : ", slice2)
	fmt.Println("ex 1 : ", slice3)

	fmt.Println()

	// 예제 2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex 2 : ", a)
	fmt.Println("ex 2 : ", b)

	fmt.Println()

	// 예제 3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적으로 슬라이스 추출은 참조이다.(복사가 아닌) -> 원본 값 변경 된다.

	d[1] = 7
	fmt.Println("ex 3 : ", c)
	fmt.Println("ex 3 : ", d)

	fmt.Println()

	// 예제 4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 마지막 값은 용량 지정이다!

	fmt.Println("ex 4 : ", len(f), cap(f))
	fmt.Println("ex 4 : ", f)

}
