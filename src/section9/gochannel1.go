package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work 1 : S ----> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work 1 : E ----> ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work 2 : S ----> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work 2 : E ----> ", time.Now())
	v <- 2
}

func main() {
	// 채널 (channel) -> 동기식이다.
	// 고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용
	// 실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	// 데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고받을 수 있음)
	// interface{} 전달을 통해서 자료형 상관없이 전송 및 수신가능
	// 값을 전달(복사 후 : bool, int 등), 포인터(슬라이스 , 맵)등을 전달시에는 주의 해야한다. > 동기화 사용 (Mutex)
	// 멀티프로세싱 처리에서 교착상태(경쟁) 주의!
	// 선언 방식 : <- , -> 으로 하며 EX) =>  (채널 <- 데이터) : 송신이며, 데이터가 채널로 가라~ , (  <- 채널) : 수신

	// 예제 1
	fmt.Println("Main : S ----> ", time.Now())

	v := make(chan int) // Int 형 채널 선언
	go work1(v)
	go work2(v)

	<-v
	<-v
	fmt.Println("Main : E ----> ", time.Now())

}
