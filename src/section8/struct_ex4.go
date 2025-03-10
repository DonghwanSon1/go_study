package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

type Executives struct {
	Employee     // is a 관계 라고 한다.
	specialBonus float64
}

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재사용 하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	// 예제 1
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	ex := Executives{Employee{"lee", 5000000, 1000000}, 10000000}

	fmt.Println("ex 1 : ", int(ep1.Calculate()))
	fmt.Println("ex 1 : ", int(ep2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("ex 2 : ", int(ex.Calculate()+ex.specialBonus))

}
