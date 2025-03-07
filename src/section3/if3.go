package main

import "fmt"

func main() {
	i := 130

	// if - else if - else 예제 1
	if i >= 120 {
		fmt.Println("120 이상입니다.")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만입니다.")
	} else if i >= 50 && i < 100 {
		fmt.Println("50 이상 100 미만입니다.")
	} else {
		fmt.Println("50 미만입니다.")
	}
}
