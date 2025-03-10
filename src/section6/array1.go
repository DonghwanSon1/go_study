package main

import "fmt"

func main() {
	// 배열
	// 배열은 용량, 길이 항상 같다.
	//    배열 : 길이 고정,  값타입,    복사 전달, 전체 비교연산 사용 가능
	// 슬라이스 : 길이 가변, 참조타입, 참조 값 전달,   비교 연산자 사용 불가
	// 따라서 대부분 슬라이스 사용한다.

	// cap() : 배열, 슬라이스 용량 구하는 함수
	// len() : 배열, 슬라이스 개수 구하는 함수

	// 예제 1
	var arr1 [5]int // 초기화를 선언하지 않은 빈 배열 선언
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} // ... 은 배열크기 자동 맞춤
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	arr1[2] = 5

	fmt.Printf("%-5T %d, %v \n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d, %v \n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d, %v \n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d, %v \n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d, %v \n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d, %v \n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d, %v \n", arr7, len(arr7), arr7)

	fmt.Println()

	// 예제 2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}

	fmt.Printf("%-5T %d, %v \n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d, %v \n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d, %v \n", arr10, len(arr10), arr10)

}
