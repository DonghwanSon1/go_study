package main

import "fmt"

func main() {
	// 제어문(조건문)
	// switch
	// switch 뒤 표현식 생략 가능, case 뒤 표현식 사용 가능, 자동 break 때문에 fallthrough 존재한다.
	// => fallthrough -> 해당 명령어를 붙이면 아래 케이스까지 같이 진행한다.
	//					(즉, 해당 case에 break 가 아닌 아래 case 까지 진행 후 Break 한다.)
	// Type 분기 가능 -> 값이 아닌 변수 Type 으로 분기 가능하다.

	// 예제 1
	a := 7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 예제 2
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// 예제 3
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치 값 없음.")
	}

	// 예제 4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("GoLang!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("NONE")
	}

	// 예제 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i와 j는 같다.")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}
}
