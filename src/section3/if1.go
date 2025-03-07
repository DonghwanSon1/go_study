package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF문 : 반드시 Boolean 으로 검사해야함(1,0 으로는 안됨 true, false 로 해야함)
	// 소괄호 사용 X

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15 이상입니다.")
	}

	if b >= 25 {
		fmt.Println("25 이상입니다.")
	}

	if c := 40; c >= 35 {
		fmt.Println("35 이상입니다.")
	}

	/**
	에러발생
	if b >= 25
	{
	}
	=> 이렇게 {}를 내려버리면 에러 발생한다.

	if b >= 25
		fmt.Println("~~")
	=> 이렇게 {}를 생략하면 에러 발생한다.

	if c :=1; c {
	}
	=> 이렇게 0,1 으로 조건문을 걸면 에러 발생한다 -> True, False 로 해야한다.
	*/
}
