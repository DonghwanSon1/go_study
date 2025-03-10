package main

import "fmt"

type Dog3 struct {
	name   string
	weight int
}

type Cat1 struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현
func (d Dog3) run() {
	fmt.Println(d.name, " : Dog running!!")
}

// 구조체 Cat1 메소드 구현
func (c Cat1) run() {
	fmt.Println(c.name, " : Cat running!!")
}

func act1(animal interface{ run() }) {
	animal.run()
}

func main() {
	// 익명 인터페이스 사용 예제 (즉시 선언 후 사용)

	// 예제 1
	dog := Dog3{"poll", 10}
	cat := Cat1{"bob", 5}

	// 개 행동 실행
	act1(dog)

	// 고양이 행동 실행
	act1(cat)
}
