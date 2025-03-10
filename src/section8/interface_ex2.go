package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 타입 변환
	// 실행 (런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	// 인터페이스.(타입) -> 형변환 을 이렇게 한다. EX) interfaceVal.(type)  => 이렇게 사용하는 방식이 Type Assertion 이다.

	// 예제 1
	var a interface{} = 15
	b := a
	c := a.(int)
	/**
	d := a.(float64)
	이렇게는 처음 인터페이스에 할당한 값이 int 이기 때문에 형변환은 int 로만 타입을 변경 할 수 있다.
	*/

	fmt.Println("ex 1 : ", a)
	fmt.Println("ex 1 : ", reflect.TypeOf(a))
	fmt.Println("ex 1 : ", b)
	fmt.Println("ex 1 : ", reflect.TypeOf(b))
	fmt.Println("ex 1 : ", c)
	fmt.Println("ex 1 : ", reflect.TypeOf(c))

	fmt.Println()

	// 예제 2 - 저장된 실제 타입 검사
	if v, ok := a.(int); ok {
		fmt.Println("ex 2 : ", v, ok)
	}

}
