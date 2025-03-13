package main

import (
	"fmt"
	"os"
)

func main() {
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의 (오류 메시지 확인)
	// 예외 처리 정말 중요!
	// 탐색 Seek 중요
	// https://golang.org/pkg/os

	// 예제 - 파일 읽기
	// 파일 열기
	file, err := os.Open("sample.txt")
	errCheck5(err)

	// 파일 읽기
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	errCheck6(err)

	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력 (1) : ", fileInfo, "\n")
	fmt.Println("파일 이름 (2) : ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기 (3) : ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간 (4) : ", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 (1) 완료 (%d bytes) \n\n", ct1)
	fmt.Println(string(fd1))
	fmt.Println("-------------------------------------------")

	// 예제 2 - 탐색 (Seek(Offset))
	o1, err := file.Seek(20, 0) // offset -> 시작할 위치(커서) EX) 20번째부터 시작 | whence -> 0: 처음 위치, 1 : 현재 위치, 2 : 마지막 위치 Ex) 20번째부터
	errCheck5(err)

	fd2 := make([]byte, o1, 20)
	ct2, err := file.Read(fd2)
	errCheck6(err)

	fmt.Printf("읽기 작업 (2) 완료 (%d bytes) (%d ret) \n\n", ct2, o1)
	fmt.Println(string(fd2)) // 시작한 위치는 20번째부터 시작하고, whence 는 첫위치니깐 20번째부터 시작이며, 읽은 슬라이스 크기가 20이니깐 20개까지 가져온다.
	fmt.Println("-------------------------------------------")

	// 읽기 예제 3 - readAt
	o2, err := file.Seek(0, 0) // 0,0 이니 전체 가져온다.
	errCheck5(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8)
	errCheck6(err)

	fmt.Printf("읽기 작업 (3) 완료 (%d bytes) (%d ret) \n\n", ct3, o2)
	fmt.Println(string(fd3))
	fmt.Println("-------------------------------------------")

	defer file.Close()
}

// 에러 체크 방식 1
func errCheck5(err error) {
	if err != nil {
		panic(err)
	}
}

// 에러 체크 방식 2
func errCheck6(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
