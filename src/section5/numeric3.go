package main

import "fmt"

func main() {
	// 실수(부동소수점) - float32(7자리), float64(15자리)

	// 예제 1
	var num1 float32 = 0.14
	var num2 float32 = .647
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	// 지수 표기법
	var num5 float32 = 14e6
	var num6 float32 = .156875e+3
	var num7 float32 = 5.32521e-10

	fmt.Println("ex 1 : ", num1)
	fmt.Println("ex 2 : ", num2)
	fmt.Println("ex 3 : ", num3)
	fmt.Println("ex 4 : ", num4)
	fmt.Println("ex 5 : ", num4-0.1)
	fmt.Println("ex 6 : ", float32(num4-0.1))
	fmt.Println("ex 7 : ", float64(num4-0.1)) // 주의해야한다. 미세하게 부동소수점 오차 때문에 값이 달라진다. math.Round(float64(num4-0.1)*100)/100) => 이렇게 해결할 수 있다.
	fmt.Println("ex 8 : ", num5)
	fmt.Println("ex 9 : ", num6)
	fmt.Println("ex 10 : ", num7)

}
