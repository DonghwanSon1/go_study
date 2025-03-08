package main

import "fmt"

func main() {
	// 반목문 - for
	// Go 에서 유일하게 제공되는 반복문
	// 다양한 사용법 숙지해야함.

	// 예제 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex 1 : ", i)
	}

	/**
	에러 발생 1
	for i := 0; i < 5; i++
	{
	}
	이렇게 {}를 아래에다가 찍히면 에러발생한다.

	에러 발생 2
	for i := 0; i < 5; i++
	 fmt.Println("~~")
	이렇게 {} 없으면 에러가 발생한다.
	*/

	// 예제 2 - 무한 루프
	for {
		fmt.Println("ex 2 : Hello golang!")
		break // 무한 루프 하고 싶을 시 break 없애기
	}

	// 예제 3 - Range 용법 (index 생략하고 싶을 시 _ 로 하면된다.)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex 3 : ", index, name)
	}

}
