package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 채널
	// 예제 1 (동기 : 버퍼 미사용)
	// 버퍼 : 발신 -> 가득차면 대기, 비어있음 작동 | 수신 -> 비어있으면 대기, 가득차 있으면 작동
	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 3)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
