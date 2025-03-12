package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Open File : ", f.Name())
	}

	defer f.Close()
}

func main() {
	// Panic / Recover 최종 실습

	// 예제
	fileOpen("test.txt")
	fmt.Println("End Main!!")

}
