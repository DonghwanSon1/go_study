package main

import "fmt"

type Dog4 struct {
	name   string
	weight int
}

type Cat2 struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex 1 : ", s)
}

func main() {
	// 인터페이스 활용 (빈 인터페이스)
	// 함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. (만능)
	// -> 즉, 모든 타입 지정가능하다.

	dog := Dog4{"poll", 10}
	cat := Cat2{"bob", 5}

	printValue(dog)
	printValue(cat)
	printValue(15)
	printValue(true)
	printValue("Animal")
	printValue(25.5)
	printValue([]Dog4{})
	printValue([5]Cat2{})

}
