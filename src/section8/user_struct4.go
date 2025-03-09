package main

import "fmt"

type shoppingBasket struct {
	cnt, price int
}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본 수정(참조 형식이여서)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본 수정 X (값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 리시버(구조체 메소드) 전달(값, 참조)형식
	// 함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달이다.
	// 리시버(구조체)도 마찬가지로 포인터를 활용해서 메소드 내에서 원본 수정 가능

	// 예제 1
	bs1 := shoppingBasket{3, 5000}
	fmt.Println("ex(totPrice) 1 : ", bs1.purchase())

	bs1.rePurchaseP(7, 5000) // 참조 전달 (원본 값 수정)
	fmt.Println("ex(totPrice) 2 : ", bs1.purchase())
	bs1.rePurchaseD(10, 0) // 값 전달 (원본 값 수정 X)
	fmt.Println("ex(totPrice) 3 : ", bs1.purchase())
}
