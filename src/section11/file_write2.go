package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일 쓰기
	// csv 파일 쓰기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/teaLeg/xlsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Create("test_write.csv")
	errCheck3(err)

	// csv Write  생성
	wr := csv.NewWriter(file)
	//wr := csv.NewWriter(bufio.NewWriter(file)) // bufio 사용 (큰 파일 사용시 권장)

	// csv 내용 쓰기
	wr.Write([]string{"Kim", "4.8"})
	wr.Write([]string{"Lee", "4.2"})
	wr.Write([]string{"Park", "4.4"})
	wr.Write([]string{"Cho", "4.45"})
	wr.Write([]string{"Hong", "4.9"})

	wr.Flush() // 버퍼 -> 파일로 쓰기

	fi, err := file.Stat()
	errCheck3(err)
	fmt.Printf("CSV 쓰기 작업 후 파일 크기 (%d byte)\n", fi.Size())
	fmt.Println("CSV 파일명 : ", fi.Name())
	fmt.Println("운영체제 파일 권한 : ", fi.Mode())

	defer file.Close()
}

// 에러 체크 방식 1
func errCheck3(err error) {
	if err != nil {
		panic(err)
	}
}

// 에러 체크 방식 2
func errCheck4(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
