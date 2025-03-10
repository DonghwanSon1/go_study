package main

import "fmt"

type Car2 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	// 중첩 구조체
	// 예제 1
	car1 := Car2{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	// 내부 spec 구조체 값 출력
	fmt.Println("ex 1 : ", car1.name)
	fmt.Println("ex 1 : ", car1.color)
	fmt.Println("ex 1 : ", car1.company)
	fmt.Println("ex 1 : ", car1.detail.width)
	fmt.Println("ex 1 : ", car1.detail.height)
	fmt.Println("ex 1 : ", car1.detail.length)
	fmt.Printf("ex 1 : %#v", car1.detail)

	fmt.Println()

	// 예제 2
	fmt.Println("ex 2 : ", car1.detail.width)
	fmt.Println("ex 2 : ", car1.detail.height)
	fmt.Println("ex 2 : ", car1.detail.length)
}
