package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외 처리 구조체
type PowError struct {
	time    time.Time   // 에러 발생 시간
	value   interface{} // 파라미터
	message string      // 에러 메세지
}

func (e *PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value : %g) - %s", e.time, e.value, e.message)
}

// 아래 처럼 인터페이스로 값을 받을 시 !!
// func Power2(f, i float64) (float64, error) {
func Power2(f, i interface{}) (float64, error) {
	if f == 0 {
		return 0, &PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	//if math.IsNaN(f) || math.IsInf(f, 0) {
	if exp, ok := f.(float64); !ok || exp == 0 {
		return 0, &PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	return math.Pow(f.(float64), i.(float64)), nil

}

func main() {
	// 에러 처리 고급
	// error 타입이 아닌 경우 에러 처리 방법
	// Error 메소드를 구현해서 사용자 정의 에러 처리 예제 심화
	// 구조체를 사용해서 세부적인 정보 출력

	// 예제 1
	v, err := Power2(10.0, 3.0) // 정상 처리
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ex 1 : ", v)
	}

	t, err := Power2(0.0, 3.0) // 정상 처리
	if err != nil {
		log.Fatal(err.Error())
		//fmt.Println(err.(*PowError).value)
		//fmt.Println(err.(*PowError).message)
		//fmt.Println(err.(*PowError).time)

	} else {
		fmt.Println("ex 2 : ", t)
	}

}
