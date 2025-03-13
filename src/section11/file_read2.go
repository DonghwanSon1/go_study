package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일 읽기

	// 예제 1 - CSV 읽기
	// 파일생성
	file, err := os.Open("sample.csv")
	errCheck7(err)

	// CSV Reader 생성
	// rr := csv.NewReader(file)  // 용량이 작을땐 해당으로 가져와도 상관 X
	rr := csv.NewReader(bufio.NewReader(file))

	// 내용 읽기
	row, err := rr.Read() // 한줄 씩 (가로) 가져온다.(첫번째 줄)
	errCheck8(err)

	row2, err2 := rr.Read() // 한줄 씩 (가로) 가져온다. (두번째 줄)
	errCheck8(err2)
	fmt.Println("CSV Read Example")
	//fmt.Println(row)
	fmt.Println(row[0])
	fmt.Println(row2[0])
	fmt.Println("--------------------------------------")

	// 예제 2
	row3, err3 := rr.ReadAll() // 위에서 첫번째, 두번째 줄을 가져왔으니 세번째 줄부터 시작한다.
	errCheck8(err3)
	fmt.Println("CSV Read All Example")
	fmt.Println(row3[5][1])
	fmt.Println("--------------------------------------")

	// Row 단위로 CSV 파일 읽기
	for i, row := range row3 {
		for j := range row {
			fmt.Printf("%s, ", row3[i][j])
		}
		fmt.Println()
	}

	// 리스소 해제
	defer file.Close()
}

// 에러 체크 방식 1
func errCheck7(err error) {
	if err != nil {
		panic(err)
	}
}

// 에러 체크 방식 2
func errCheck8(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
