package main

import (
	"fmt"
)

func main() {
	// 예제 1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex 1 = ", n1+n2)
	fmt.Println("ex 1 = ", n1-n2)
	fmt.Println("ex 1 = ", n1*n2)
	fmt.Println("ex 1 = ", n1/n2)
	fmt.Println("ex 1 = ", n1%n2)
	fmt.Println("ex 1 = ", n1<<2)
	fmt.Println("ex 1 = ", n1>>2)
	fmt.Println("ex 1 = ", ^n1)

	fmt.Println()

	// 예제 2
	var n3 int = 12
	var n4 float32 = 8.9
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	fmt.Println("ex 2 : ", float32(n3)+n4)
	fmt.Println("ex 2 : ", n3+int(n4))    // 주의 - 보통은 실수형에 맞춰서 계산하는게 보편적이다. 이유 : 나머지값이 소실되기 때문
	fmt.Println("ex 2 : ", n5+uint16(n6)) // 주의 - Uint16은 최대값이 65535 가 최대인데 현제 n6 는 120000이기에 최대 형변환이 되는 값을 확인 후 형변환 해야한다. - 값이 이상하기 때문

}
