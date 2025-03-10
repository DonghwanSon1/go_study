package main

import "fmt"

type Employee1 struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee1) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives1) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives1 struct {
	Employee1    // is a 관계 라고 한다.
	specialBonus float64
}

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴
	// 다른 관점으로 메소드를 재사용 하는 장점 제공
	// 상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	// 예제 1
	ep1 := Employee1{"kim", 2000000, 150000}
	ep2 := Employee1{"park", 1500000, 200000}

	ex := Executives1{Employee1{"lee", 5000000, 1000000}, 1000000}

	fmt.Println("ex 1 : ", int(ep1.Calculate()))
	fmt.Println("ex 1 : ", int(ep2.Calculate()))

	// Employee1 구조체 통해서 메소드 호출
	fmt.Println("ex 2 : ", int(ex.Calculate()))
	fmt.Println("ex 2 : ", int(ex.Employee1.Calculate()+ex.specialBonus))

}
