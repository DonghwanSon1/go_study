package main

import "fmt"

func main() {
	// 슬라이스 참조 타입 증명

	// 예제 1 - 배열
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 // 복사한거임
	arr2[0] = 7

	fmt.Println("ex 1 : ", arr1)
	fmt.Println("ex 1 : ", arr2)

	fmt.Println()

	// 예제 2 - 슬라이스
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1 // 참조하는거임
	slice2[0] = 7

	fmt.Println("ex 2 : ", slice1)
	fmt.Println("ex 2 : ", slice2)

	fmt.Println()

	/**
	예외 상항
	slice3 := make([]int, 50, 100)
	fmt.Println("ex 3 : ", slice3[50])
	이렇게 되면 길이를 50 즉, 인덱스로 하면 49까지 초기화가 되어 있는데 인덱스[50]을 가져오라고 하면 에러가 발생된다.
	*/

	// 예제 3 - 슬라이스 순회
	for _, v := range slice1 {
		fmt.Println("ex 3 : ", v)
	}
}
