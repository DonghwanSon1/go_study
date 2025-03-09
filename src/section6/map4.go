package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 조회 할 경우 주의 할 점

	// 예제 1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, exist1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, exist2 := map1["kiwi"] // 두번째 리턴 값으로 키 존재 유무 확인 할 수 있다.

	fmt.Println("ex 1 : ", value1, exist1)
	fmt.Println("ex 1 : ", value2)
	fmt.Println("ex 1 : ", value3, exist2)
	fmt.Println()

	// 예제 2
	if value, ok := map1["banana"]; ok {
		fmt.Println("ex 2 : ", value)
	} else {
		fmt.Println("ex 2 : ", "Banana is not exist!!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex 3 : ", "Kiwi is not exist!!")
	}
}
