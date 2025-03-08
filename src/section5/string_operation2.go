package main

import "fmt"

func main() {

	// 예제 2 - 비교
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex 1 : ", str1 == str2)
	fmt.Println("ex 2 : ", str1 != str2)

	// Golang 은 문자열 비교할 시 아스키 코드에 대한 사전식 비교를 한다. -> 즉, 소문자와 대문자 비교를 할 시 대문자가 아스키 코드가 뒤에 있기 때문에 대문자가 더 크다고 표현한다.
	fmt.Println("ex 3 : ", str1 > str2)
	fmt.Println("ex 4 : ", str1 < str2)
}
