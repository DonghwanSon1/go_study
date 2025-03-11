package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, notify, notifyAll => 기타 언어에서 이렇게 사용
	// Wait, Signal, Broadcast => golang 에서 사용 (wait : 대기, signal : 단일 깨움, Broadcast 전체깨움)

	runtime.GOMAXPROCS(runtime.NumCPU())
	mutex := new(sync.Mutex)
	condition := sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널 (버퍼가 있으면 비동기 없으면 동기)

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting : ", n)
			condition.Wait() // 고루틴 대기(멈춤)
			fmt.Println("Waiting End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("received : ", <-c)
	}

	/**
	// 하나씩 깨울때
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake Goroutine(Signal) : ", i)
		condition.Signal() // 한개씩 깨운다. (모든 고루틴 생성 후)
		mutex.Unlock()
	}
	*/

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}
