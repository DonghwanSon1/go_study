package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}() // 맨 마지막에 호출한다~~

	func() {
		fmt.Println("Hi!!")
	}()
}

func main() {
	// 예제 1
	sayHello("Golang!!")

}
