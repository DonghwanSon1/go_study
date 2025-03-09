package main

import "fmt"

func main() {
	// map - 맵 조회 및 순회
	// 예제 1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex 1 : ", map1["google"])
	fmt.Println("ex 1 : ", map1["daum"])

	fmt.Println()

	// 예제 2 - 순서가 없으니 랜덤
	for key, value := range map1 {
		fmt.Println("ex 2 : ", key, value)
	}

	fmt.Println()

	for _, value := range map1 {
		fmt.Println("ex 3 : ", value)
	}

	fmt.Println()
}
