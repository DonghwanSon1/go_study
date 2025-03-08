package main

import "fmt"

func init() {
	fmt.Println("Init1 Method Start!")
}

func init() {
	fmt.Println("Init2 Method Start!")
}

func init() {
	fmt.Println("Init3 Method Start!")
}

func init() {
	fmt.Println("Init4 Method Start!")
}

func init() {
	fmt.Println("Init5 Method Start!")
}

func main() {

	// 여러 init 메서드가 있어도 상관은 없다.
	// 또한, 만약 import 했던 패키지 안에 init 메서드가 있다면, 그 init 메서드가 먼저 호출되고, 그다음 순차적으로 init 메서드가 실행되는 구조이다.
	// 그 후 main 메서드가 실행된다.
	fmt.Println("Main Method Start!")
}
