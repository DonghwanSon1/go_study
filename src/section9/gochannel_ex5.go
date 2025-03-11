package main

import (
	"fmt"
	"time"
)

func main() {
	// 채널 - 셀렉트

	// 예제 1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1 // 값 수신하겠다.
			fmt.Println("Channel 1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Golang Hello World!" // 값 발신한다.
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case ch1 <- 777: // 발신 용도
			case str := <-ch2:
				fmt.Println("Channel 2 : ", str)
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
