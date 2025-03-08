package main

import "fmt"

func main() {
	// 문자열 표현 - UTF-8 인코딩으로 golang은 지원한다. (유니코드 문자 집합)
	// 그러나 바이트 수 주의!!

	// 예제 1
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	// 이렇게 출력을 하면, 유니코드 값들이 출력된다.
	fmt.Println("ex 1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("ex 1 : ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("ex 1 : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	fmt.Println()

	// 예제 2
	fmt.Printf("ex 2 : %c %c %c %c %c %c \n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("ex 2 : %c %c %c %c %c \n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("ex 2 : %c %c %c %c %c %c \n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) // 한글 깨짐

	conStr := []rune(str3)
	fmt.Printf("ex 2 : %c %c %c %c %c %c \n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5]) // 한글 정상 출력

	fmt.Println()

	// 예제 3
	for i, char := range str1 {
		fmt.Printf("ex 3 : %c(%d)\t", char, i)
	}

	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("ex 4 : %c(%d)\t", char, i)
	}
}
