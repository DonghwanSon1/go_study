package main

import (
	"fmt"
	checkUp "section4/lib"
	_ "section4/lib2"
)

func main() {
	// 패키지 접근제어
	// 별칭 사용
	// 빈 식별자 사용 => _ "section4/lib2" 이렇게 _ 로 넣어주면 일단 임시로 import 할 수 있다.

	fmt.Println("10 보다 큰 수 ? : ", checkUp.CheckNum(100))
}
