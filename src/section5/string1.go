package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열 - 큰 따음표 "", 백스쿼트 ``로 감싼다.
	// golang : 문자 char '' 제공 하지 않는다. -> rune(int32) 으로 문자 코드 값으로 표현 해야한다.
	// 문자 : '' 로 작성은 가능하다.
	// 자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄 바꿈), \r(복귀), \t(탭) ....

	// 예제 1
	var str1 string = "c:\\go_study\\src\\" // -> C:\go_study\src\
	str2 := `c:\go_study\src\`              // -> Escape 를 사용하지 않아도됨

	fmt.Println("ex 1 = ", str1)
	fmt.Println("ex 1 = ", str2)

	fmt.Println()

	// 예제 2
	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요,"
	var str5 string = "\ud55c\uae00"

	fmt.Println("ex 2 = ", str3)
	fmt.Println("ex 2 = ", str4)
	fmt.Println("ex 2 = ", str5)

	fmt.Println()

	// 예제 3 - 길이(바이트 수) = 한글은 3바이트, 영문 및 나머지 띄어쓰기 등등은 1바이트
	fmt.Println("ex 3 = ", len(str3))
	fmt.Println("ex 3 = ", len(str4))

	fmt.Println()

	// 예제 4 - 실제 길이 => 이게 우리가 알고 있는 len 함수다......!!!
	fmt.Println("ex 4 = ", utf8.RuneCountInString(str3))
	fmt.Println("ex 4 = ", utf8.RuneCountInString(str4))
	fmt.Println("ex 4 = ", len([]rune(str3))) // 이 방식도 가능하다.
	fmt.Println("ex 4 = ", len([]rune(str4)))

}
