package main

import "fmt"

func main() {
	// 상수 - const 사용으로 초기화하며, 한번 선언 후 값 변경 금지 및 고정 된 값 (관리용)
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	const e = 35.6
	const f = false

	/**
	에러발생 경우
	const g string = "Test1"
	g = "Test"  => 이럴경우 에러 발생함 - 한번 선언 후 수정 불가

	또한 함수의 리턴값으로 const로 선언 후 담는건 불가능 - 항상 동일한 값으로 리턴 한다는 보장이 없기 때문에
	*/
	fmt.Println(a, b, c, e, f)
}
