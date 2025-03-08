package main

import "fmt"

func main() {
	// 예제 1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex 1 : ", i, j)
		}
	}

	// 예제 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex 2 : ", i)
	}

	// 예제 3
Loop2:
	/**
	Loop 레이블 밑에 관련이 없는 소스코드가 있다면, 에러발생한다.
	*/
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue Loop2
			}
			fmt.Println("ex 3 : ", i, j)
		}
	}
}
