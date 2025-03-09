package main

import "fmt"

func main() {
	// 슬라이스 추가 및 병합
	// 예제 1
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7, 8, 9, 10, 11)
	s2 = append(s1, s2...)      // 슬라이스를 삽입 할 시 ... 사용해야함.
	s3 = append(s2, s3[0:3]...) // 슬라이스를 추출 후 해당 슬라이스를 삽입 하니 동일하게 ... 사용해야함.

	fmt.Println("ex 1 : ", s1)
	fmt.Println("ex 1 : ", s2)
	fmt.Println("ex 1 : ", s3)

	fmt.Println()

	// 예제 2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex 2 -> len : %d, cap : %d, value : %v \n", len(s4), cap(s4), s4)
	}
}
