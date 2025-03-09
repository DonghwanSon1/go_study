package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

// 일반 메소드
func price(c Car) int64 {
	return c.price + c.tax
}

// 구조체는 <-> 메소드 바인딩
func (c Car) price2() int64 {
	return c.price + c.tax
}

func main() {
	// Golang -> 객체지향 타입을 구조체로 정의한다. (클래스, 상속 개념이 없기 때문에)
	// 객체지향 -> 클래스(속성, 기능(상태)) => 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그래밍
	// 객체지향 언어가 될까?? -> O
	// Golang은 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 클래스형태의 코딩 가능하다.
	// 객체지향의 기본 개념은 가지고 있다.
	// 상태, 메소드 분리해서 정의(결합성 없음)
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본타입(int, float, string 등), 함수
	// 구조체와 메소드 연결을 통해서 타 언어의 클래스 형식 처럼 사용 가능 (객체 지향)

	// 예제 1
	bmw := Car{
		name:  "520d",
		color: "white",
		price: 50000000,
		tax:   5000000,
	}

	benz := Car{
		name:  "220d",
		color: "white",
		price: 60000000,
		tax:   6000000,
	}

	fmt.Println("ex 1 : ", bmw, &bmw)
	fmt.Println("ex 1 : ", benz, &benz)

	// 일반 메서드
	fmt.Println("ex 2 : ", price(bmw))
	fmt.Println("ex 2 : ", price(benz))

	// 객체 지향적 메서드
	fmt.Println("ex 3 : ", bmw.price2())
	fmt.Println("ex 3 : ", benz.price2())

	fmt.Println("ex 4 : ", benz == bmw)

}
