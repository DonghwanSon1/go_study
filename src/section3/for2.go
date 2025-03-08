package main

import "fmt"

func main() {
	// 예제 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("sum1: ", sum1)

	// 예제 2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
		// 후치연산은 반환 값 X - EX) j := i++ 이런건 X
	}
	fmt.Println("sum2: ", sum2)

	// 예제 3 - while 형태랑 비슷한 방식
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("sum3: ", sum3)

	// 예제 4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex 4 : ", i, j)
	}

	/**
	에러 발생 1
	for i, j := 0, 0; i <= 10; i++, j+= 10 {
	}
	이렇게 후치연산 (i++)은 반환값이 없어서 에러발생한다.
	*/

}
