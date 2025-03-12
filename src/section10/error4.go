package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// 에러 처리
	// Errorf 를 이용한 에러 처리

	a, err := notZero1(1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ex 1 : ", a)

	b, err := notZero1(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ex 2 : ", b)

	fmt.Println("End Error Handling!!")
}

func notZero1(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang -> ", n)
		return s, nil
	}
	return "", errors.New("-> 0을 입력했습니다 에러 발생") // 이렇게 보통 많이 사용한다.
}
