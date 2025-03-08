package main

import (
	"fmt"
	"section4/lib"
)

var num int32

// 변수 초기화
func init() {
	num = 30
	fmt.Println("Main init Start!")
}

func main() {

	// 설명 : import 안에 있는 Lib 패키지에 있는 checkNum 파일안에 Init 메서드 실행 후
	// 해당 main 안에 있는 Init 메서드 호출 한 후
	// main 함수 안에 있는 checkNum 함수 호출 하는것이다.
	fmt.Println("10 보다 큰 수 ? : ", lib.CheckNum(num))

}
