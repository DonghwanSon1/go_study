package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 연산(산술, 비교)
	// 타입이 같아야 가능하다. -> 타입이 다르면 에러 발생하며, 무조건 한쪽을 타입 맞춰줘야 한다.
	// +, -, *, /, %, <<, >>, &, ^

	// 예제 1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex 1 =", n1)
	fmt.Println("ex 1 =", n2)
	fmt.Println("ex 1 =", n3)
	fmt.Println("ex 1 =", n4)
	fmt.Println("ex 1 =", math.MaxInt8)
	fmt.Println("ex 1 =", math.MaxInt16)
	fmt.Println("ex 1 =", math.MaxInt32)
	fmt.Println("ex 1 =", math.MaxInt64)
	fmt.Println("ex 1 =", math.MaxFloat32)
	fmt.Println("ex 1 =", math.MaxFloat64)

	n5 := 100000 //int
	n6 := int16(10000)
	n7 := uint8(100)

	/**
	예외 발생
	fmt.Println("ex 2 =", n5 + n6)
	해당 부분은 n5는 Int 이고, n6는 Int16 이기에 형변환을 하지 않아 에러 발생한다.
	*/
	fmt.Println("ex 2 = ", n5+int(n6))
	fmt.Println("ex 2 = ", n6+int16(n7))
	fmt.Println("ex 2 = ", n6 > int16(n7)) // 비교 연산도 형변환 후 가능하다.
	fmt.Println("ex 2 = ", n6-int16(n7) > 5000)

}
