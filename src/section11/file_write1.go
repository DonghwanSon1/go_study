package main

import (
	"fmt"
	"os"
)

func main() {
	// 파일 쓰기
	// Create : 새 파일 작성 및 파일 열기
	// Close : 리소스 닫기
	// Write, WriteString, WriteAt : 파일 쓰기
	// 각 운영체제 권한 주의(오류 메시지 확인)
	// 예외 처리 정말 중요!!
	// https://golang.org/pkg/os

	// 예제1 - 파일 쓰기
	file, err := os.Create("test_write.txt")
	errCheck1(err)

	s1 := []byte{48, 49, 50, 51, 52}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 반환 후 쓰기
	errCheck2(err)
	fmt.Printf("쓰기 작업(1) 완료 (%d byte) -1\n", n1)

	file.Sync() // Write Commit -> 파일 동기화 작업 함수

	// 예제 2 - WriteString
	s2 := "\nHello Golang! \n File Write Test !! -2\n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("쓰기 작업(2) 완료 (%d byte) -2\n", n2)

	file.Sync()

	// 예제 3 - WriteAt
	s3 := "Test WriteAt!! -3"
	n3, err := file.WriteAt([]byte(s3), 100) // - 70번째에서 부터 쓰기 작업 한다 (s3 를)
	errCheck1(err)
	fmt.Printf("쓰기 작업(3) 완료 (%d byte)- 3 \n", n3)

	file.Sync()

	// 예제 4
	n4, err := file.WriteString("\nHello Golang!! \n File Write Test!! -4 \n")
	errCheck1(err)
	fmt.Printf("쓰기 작업(4) 완료 (%d byte) \n", n4)

	// 리소스 해제
	defer file.Close()
}

// 에러 체크 방식 1
func errCheck1(err error) {
	if err != nil {
		panic(err)
	}
}

// 에러 체크 방식 2
func errCheck2(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
