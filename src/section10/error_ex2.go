package main

import (
	"fmt"
	"math"
)

// f의 i 제곱 구하는 함수
func Power1(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(%g)는/은 사용 불가능 합니다", f)
	}
	return math.Pow(f, i), nil
}

func main() {
	// 에러 처리 고급
	// Golang error 패키지 New 메소드 사용 -> 사용자 에러 처리 생성

	// 예제 1
	if f, err := Power1(7, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex 1 : ", f)
	}

	// 예제 2
	if f, err := Power1(0, 3); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("ex 2 : ", f)
	}
}
