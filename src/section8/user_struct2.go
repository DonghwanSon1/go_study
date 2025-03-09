package main

import "fmt"

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입
	// 예제 1
	a := cnt(15)
	fmt.Println("ex 1 : ", a)

	// 예제 2
	var b cnt = 15
	fmt.Println("ex 1 : ", b)

	// 예제 3
	/**
	에러 발생
	testConverT(b)
	타입이 Int 여도 type안에 생성된 타입으로 b를 줬기때문에 타입에러로 에러발생한다.
	따라서 사용자 정의 타입과 기본타입은 매개변수 전달 시에 변환해야 사용 가능하다. (int(변수))
	*/
	testConverT(int(b)) // 형변환 후 요청해야 한다.
	testConverD(b)

}

func testConverT(i int) {
	fmt.Println("ex 3 : (Default Type) : ", i)
}

func testConverD(i cnt) {
	fmt.Println("ex 3 : (Custom Type) : ", i)
}
